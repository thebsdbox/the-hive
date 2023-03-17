package challenges

import (
	"context"
	"time"

	apiv1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

var challenge4 = Challenge{
	Name:        "Fear and Loathing with network policies 🕸️",
	Description: "What's in a policy ¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
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
					"app": "demo", // RUHROH
				},
			},
		}

		_, err = clientSet.CoreV1().Services(apiv1.NamespaceDefault).Create(ctx, service, v1.CreateOptions{})
		if err != nil {
			return err
		}

		net := &networkv1.NetworkPolicy{
			ObjectMeta: v1.ObjectMeta{
				Name: "hive-policy",
			},
			Spec: networkv1.NetworkPolicySpec{
				PodSelector: *&v1.LabelSelector{
					MatchLabels: map[string]string{
						"app": "demo", // RUHROH
					},
				},
				Ingress: []networkv1.NetworkPolicyIngressRule{},
			},
		}

		clientSet.NetworkingV1().NetworkPolicies(apiv1.NamespaceDefault).Create(ctx, net, v1.CreateOptions{})

		return nil

	},
	Readme: `
Welcome to "The Hive"
--------------------------------

Enable hubble (optional)
-------------
kubectl expose -n kube-system deploy hubble-ui --type=NodePort --name hubble-node
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge4)
}
