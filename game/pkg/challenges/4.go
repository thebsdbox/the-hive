package challenges

import (
	"context"
	"time"

	apiv1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var challenge4 = Challenge{
	Name:        "Fear and Loathing with network policies 🕸️",
	Description: "What's in a policy ¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas

		// Create Network Policy

		net := &networkv1.NetworkPolicy{
			ObjectMeta: v1.ObjectMeta{
				Name: "hive-policy",
			},
			Spec: networkv1.NetworkPolicySpec{
				PodSelector: *&v1.LabelSelector{
					MatchLabels: map[string]string{
						"app": "middle", // RUHROH
					},
				},
				Ingress: []networkv1.NetworkPolicyIngressRule{},
			},
		}

		clientSet.NetworkingV1().NetworkPolicies(apiv1.NamespaceDefault).Create(ctx, net, v1.CreateOptions{})

		return deployObjects(ctx, clientSet, false)

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
