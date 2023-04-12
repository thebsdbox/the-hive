package challenges

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var challenge3 = Challenge{
	Name:        "Not on target (port)🎯",
	Description: "Something doesn't match up  ¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas
		backendEndService.Spec.Ports[0].TargetPort = intstr.FromInt(81) // ruhroh

		return deployObjects(ctx, clientSet)

	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge3)
}
