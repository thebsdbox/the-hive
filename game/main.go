package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/creack/pty"
	"github.com/kube-vip/kube-vip/pkg/k8s"
	"github.com/mattn/go-tty"
	"github.com/sirupsen/logrus"
	"github.com/thebsdbox/the-hive/game/pkg/challenges"
	"github.com/thebsdbox/the-hive/game/pkg/k3d"
	"golang.org/x/term"
)

// Intro
// Deploy the cluster
// Drop to a shell
// Timeout

func main() {

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	// Set the scene
	clearScreen()
	MoveTo(5, 10)
	// This is to allow localised testing
	_, testing := os.LookupEnv("TEST")
	if !testing {
		os.Setenv("DOCKER_HOST", "tcp://localhost:2375")
	}
	// Wargames intro
	slowStringPrint("Would you like to play a game?\n", time.Millisecond*50)
	waitOnKey()
	clearScreen()
	MoveTo(5, 10)

	//
	slowStringPrint("We're playing one regardless...\n", time.Millisecond*50)
	waitOnKey()
	clearScreen()

	// Display game choices
	challenge, err := challenges.SelectChallenge()
	if err != nil {
		panic(err)
	}

	// Create the cluster
	cluster, err := k3d.CreateCluster("falken")
	if err != nil {
		panic(err)
	}
	defer k3d.DeleteCluser(ctx, cluster)

	// Get kubernetes client
	homeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	clientset, err := k8s.NewClientset(homeConfigPath, false, "")
	if err != nil {
		log.Fatalf("could not create k8s clientset from external file: %q: %v", homeConfigPath, err)
	}

	// Give the challenge the Kubernetes GO client
	err = challenge.DeployFunc(ctx, clientset)
	if err != nil {
		panic(err)
	}

	logrus.Infof("Challenge [%s] selected", challenge.Name)
	logrus.Infof("You will have [%s] to complete the challenge before the cluster will self destruct", challenge.AllowedTime.String())

	// Deploy the readme file
	err = challenge.CreateReadme()
	if err != nil {
		panic(err)
	}

	go func() {

		// We ignore the errors (they're unlikely, but ultimate wont impact playing the game)
		f, _ := os.Create("/tmp/counter")
		defer f.Close()

		t := time.Now().Add(challenge.AllowedTime)
		for range time.Tick(1 * time.Second) {
			f.Truncate(0) // Reduce the file to 0 bytes before re-writing the time remaining
			y := t.Sub(time.Now())
			f.WriteString(baseName(y.String()))
			time.Sleep(time.Second)

		}
	}()
	// Start the shell (blocking until timeout)
	err = startShell(ctx, challenge.AllowedTime)
	if err != nil {
		panic(err)
	}

	clearScreen()
	MoveTo(5, 10)

	slowStringPrint("That's all folks !\n", time.Millisecond*50)
	time.Sleep(time.Second * 3)
}

func waitOnKey() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()
	tty.ReadRune()
}

func startShell(ctx context.Context, t time.Duration) error {
	deadline := time.Now().Add(t)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	defer cancelCtx()
	c := exec.CommandContext(ctx, "/bin/bash", "-l")

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	// NOTE: The goroutine will keep reading until the next keystroke before returning.
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	return nil
}

// This is a quick function to take a string and print a single character between each delay
func slowStringPrint(s string, t time.Duration) {

	for _, rune := range s {
		fmt.Printf("%c", rune)
		time.Sleep(t)
	}
}

func baseName(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n == -1 {
		return s
	}
	return s[:n]
}

// ClearScreen clears the screen.
func clearScreen() {
	fmt.Printf("\033[2J")
	MoveTo(1, 1)
}

// MoveTo moves the cursor to (x, y).
func MoveTo(x, y int) {
	fmt.Printf("\033[%d;%df", y, x)
}
