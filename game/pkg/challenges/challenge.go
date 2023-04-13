package challenges

import (
	"context"
	"os"
	"time"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Challenges contains all of the Kubernetes challenges
var Challenges []Challenge

// Challenge is an individual challenge, with a specified time limit
type Challenge struct {
	Name        string
	Description string
	AllowedTime time.Duration
	ClusterName string
	DeployFunc  func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error
	Readme      string
}

// CreateReadme will write a readme to the home directory, it can contain clues and other useful information
func (c *Challenge) CreateReadme() error {
	f, err := os.Create("Readme.txt")

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(readmeHeader + c.Readme)
	if err != nil {
		return err
	}
	return nil
}

func deployObjects(ctx context.Context, clientSet *kubernetes.Clientset, win bool) error {
	// Create Front End

	_, err := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Create(ctx, frontEndDeployment, v1.CreateOptions{})
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().Services(apiv1.NamespaceDefault).Create(ctx, frontEndService, v1.CreateOptions{})
	if err != nil {
		return err
	}

	// Create Middle tier

	_, err = clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Create(ctx, middleEndDeployment, v1.CreateOptions{})
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().Services(apiv1.NamespaceDefault).Create(ctx, middleEndService, v1.CreateOptions{})
	if err != nil {
		return err
	}

	// Create Backend
	if win {
		_, err = clientSet.CoreV1().ConfigMaps(apiv1.NamespaceDefault).Create(ctx, winConfigMap(), v1.CreateOptions{})

	} else {
		_, err = clientSet.CoreV1().ConfigMaps(apiv1.NamespaceDefault).Create(ctx, trayagainConfigMap(), v1.CreateOptions{})

	}
	if err != nil {
		return err
	}

	_, err = clientSet.AppsV1().Deployments(apiv1.NamespaceDefault).Create(ctx, backEndDeployment, v1.CreateOptions{})
	if err != nil {
		return err
	}
	_, err = clientSet.CoreV1().Services(apiv1.NamespaceDefault).Create(ctx, backendEndService, v1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}
