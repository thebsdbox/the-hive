package challenges

import (
	"context"
	"time"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var challenge2 = Challenge{
	Name:        "The pod wont start 🙅‍♂️ ",
	Description: "Will it ever be ready?  ¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas

		// RuhRoh
		backEndDeployment.Spec.Template.Spec.Containers[0].ReadinessProbe = &apiv1.Probe{
			ProbeHandler: apiv1.ProbeHandler{
				HTTPGet: &apiv1.HTTPGetAction{
					Scheme: apiv1.URISchemeHTTP,
					Path:   "/index.html",
					Port:   intstr.FromInt(8080),
				},
			},
			InitialDelaySeconds: 5,
			PeriodSeconds:       5,
		}

		return deployObjects(ctx, clientSet, false)

	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge2)
}
