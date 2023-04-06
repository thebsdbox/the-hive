package challenges

import (
	"context"
	"time"

	"k8s.io/client-go/kubernetes"
)

var challenge7 = Challenge{
	Name:        "More copies 👯 required",
	Description: "¯\\_(ツ)_/¯",
	AllowedTime: 2 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset) error {

		replicas := int32(0)
		backEndDeployment.Spec.Replicas = &replicas

		return deployObjects(ctx, clientSet)
	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge7)
}
