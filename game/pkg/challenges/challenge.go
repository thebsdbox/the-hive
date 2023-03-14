package challenges

import (
	"time"

	"k8s.io/client-go/kubernetes"
)

var Challenges []Challenge

type Challenge struct {
	Name        string
	Description string
	AllowedTime time.Duration
	ClusterName string
	clientSet   *kubernetes.Clientset
}

func (c *Challenge) SetK8sClient(clientSet *kubernetes.Clientset) {
	c.clientSet = clientSet
}
