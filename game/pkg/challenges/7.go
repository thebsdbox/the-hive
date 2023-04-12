package challenges

import (
	"context"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var challenge7 = Challenge{
	Name:        "More copies 👯 required",
	Description: "¯\\_(ツ)_/¯",
	AllowedTime: 2 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {

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
