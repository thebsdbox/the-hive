package challenges

import (
	"context"
	"time"

	"k8s.io/client-go/kubernetes"
)

var challenge5 = Challenge{
	Name:        "Image 💥",
	Description: "¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset) error {

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas
		backEndDeployment.Spec.Template.Spec.Containers[0].Image = "nginx:1.12-nope"

		return deployObjects(ctx, clientSet)
	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge5)
}
