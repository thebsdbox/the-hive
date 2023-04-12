package challenges

import (
	"context"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var challenge6 = Challenge{
	Name:        "A problem with the 🧠",
	Description: "¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas
		backEndDeployment.Spec.Template.Spec.Containers[0].Resources = apiv1.ResourceRequirements{
			Limits: apiv1.ResourceList{
				"memory": resource.MustParse("1Mi"),
			},
		}

		return deployObjects(ctx, clientSet)
	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge6)
}
