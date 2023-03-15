package challenges

import (
	"context"
	"os"
	"time"

	"k8s.io/client-go/kubernetes"
)

// Challenges contains all of the Kubernetes challenges
var Challenges []Challenge

// Challenge is an individual challenge, with a specified time limit
type Challenge struct {
	Name        string
	Description string
	AllowedTime time.Duration
	ClusterName string
	DeployFunc  func(ctx context.Context, clientSet *kubernetes.Clientset) error
	Readme      string
}

// CreateReadme will write a readme to the home directory, it can contain clues and other useful information
func (c *Challenge) CreateReadme() error {
	f, err := os.Create("Readme.txt")

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(c.Readme)
	if err != nil {
		return err
	}
	return nil
}
