package challenges

import (
	"context"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var challenge8 = Challenge{
	Name:        "Too affinity and Beyond! 🚀",
	Description: "¯\\_(ツ)_/¯",
	AllowedTime: 2 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas

		backEndDeployment.Spec.Template.Spec.Affinity = &apiv1.Affinity{
			NodeAffinity: &apiv1.NodeAffinity{
				RequiredDuringSchedulingIgnoredDuringExecution: &apiv1.NodeSelector{
					NodeSelectorTerms: []apiv1.NodeSelectorTerm{{
						MatchExpressions: []apiv1.NodeSelectorRequirement{
							{Key: "Buzz", Operator: apiv1.NodeSelectorOpExists},
						},
					}},
				},
			},
		}

		return deployObjects(ctx, clientSet, true)
	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge8)
}
