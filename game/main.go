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
	// Set the scene
	clearScreen()

	// Wargames intro
	slowStringPrint("Would you like to play a game?\n", time.Millisecond*50)
	waitOnKey()
	clearScreen()

	//
	slowStringPrint("We're playing one regardless...\n", time.Millisecond*50)
	waitOnKey()
	clearScreen()

	// Display game choices
	challenge, err := challenges.SelectChallenge()
	if err != nil {
		panic(err)
	}
	logrus.Infof("Challenge [%s] selected", challenge.Name)
	logrus.Infof("You will have [%s] to complete the challenge before the cluster will self destruct", challenge.AllowedTime.String())

	// Create the cluster

	//err = kind.CreateKind("falken")
	err = k3d.CreateCluster("falken")
	if err != nil {
		panic(err)
	}
	//defer kind.DeleteKind("falken")

	// Get kubernetes client
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	homeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	clientset, err := k8s.NewClientset(homeConfigPath, false, "")
	if err != nil {
		log.Fatalf("could not create k8s clientset from external file: %q: %v", homeConfigPath, err)
	}
	logrus.Debugf("Using external Kubernetes configuration from file [%s]", homeConfigPath)

	// Give the challenge the Kubernetes GO client
	challenge.SetK8sClient(clientset)
	err = challenge.Deploy(ctx)
	if err != nil {
		panic(err)
	}

	// Start the shell (blocking until timeout)
	err = startShell(ctx, challenge.AllowedTime)
	if err != nil {
		panic(err)
	}

	clearScreen()
	slowStringPrint("That's all folks !\n", time.Millisecond*50)
	waitOnKey()
}

func clearScreen() {
	fmt.Print("\033[2J")
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
	c := exec.CommandContext(ctx, "/bin/bash")

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
