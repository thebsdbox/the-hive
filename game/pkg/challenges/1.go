package challenges

import (
	"context"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var challenge1 = Challenge{
	Name:        "The spicy one",
	Description: "blah blah",
	AllowedTime: 5 * time.Minute,
}

func init() {

	Challenges = append(Challenges, challenge1)

}

func (c *Challenge) Deploy(ctx context.Context) error {
	replicas := int32(2)
	deploymentsClient := c.clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := deploymentsClient.Create(ctx, deployment, v1.CreateOptions{})
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

	_, err = c.clientSet.CoreV1().Services(apiv1.NamespaceDefault).Create(ctx, service, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil

}
