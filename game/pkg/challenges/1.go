package challenges

import (
	"context"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

var challenge1 = Challenge{
	Name:        "The labels one 🏷️ ",
	Description: "For some reason the NodePort isn't working  ¯\\_(ツ)_/¯",
	AllowedTime: 3 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset) error {

		_, err := clientSet.CoreV1().ConfigMaps(apiv1.NamespaceDefault).Create(ctx, configMap, v1.CreateOptions{})
		if err != nil {
			return err
		}

		replicas := int32(2)
		deployment.Spec.Replicas = &replicas

		deploymentsClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

		_, err = deploymentsClient.Create(ctx, deployment, v1.CreateOptions{})
		if err != nil {
			return err
		}

		service := &apiv1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "exposewebsite",
				Namespace: apiv1.NamespaceDefault,
				// Labels:    GetLabels(),
			},
			Spec: apiv1.ServiceSpec{
				Type: apiv1.ServiceTypeNodePort,
				Ports: []apiv1.ServicePort{
					{
						Name:       "web",
						TargetPort: intstr.FromInt(80),
						Port:       80,
						Protocol:   "TCP",
						NodePort:   30000,
					},
				},
				Selector: map[string]string{
					"app": "demos", // RUHROH
				},
			},
		}

		_, err = clientSet.CoreV1().Services(apiv1.NamespaceDefault).Create(ctx, service, v1.CreateOptions{})
		if err != nil {
			return err
		}
		return nil

	},
	Readme: `
	Welcome to "The Hive"
	--------------------------------


	`,
}

func init() {
	Challenges = append(Challenges, challenge1)
}
