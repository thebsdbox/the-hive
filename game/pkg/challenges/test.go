package challenges

import (
	"context"
	"time"

	"k8s.io/client-go/kubernetes"
)

var test = Challenge{
	Name:        "Testy 🔬",
	Description: "As in a LITERAL test",
	AllowedTime: 1 * time.Hour,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset) error {
		return nil
	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {

	Challenges = append(Challenges, test)

}
