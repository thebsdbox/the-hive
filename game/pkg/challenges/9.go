package challenges

import (
	"context"
	"fmt"
	"time"

	ciliumv2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	ciliumslimmetav1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	ciliumpolicyapi "github.com/cilium/cilium/pkg/policy/api"

	ciliumClientset "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var challenge9 = Challenge{
	Name:        "Policies Policies Policies",
	Description: "¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {
		//rest.ConfigToExecCluster()

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas
		ciliumClientset, err := ciliumClientset.NewForConfig(r)
		//ciliumClientset.
		ciliumPolicy := &ciliumv2.CiliumNetworkPolicy{
			ObjectMeta: metav1.ObjectMeta{
				Name: "hive-policy",
			},
			Specs: ciliumpolicyapi.Rules{
				&ciliumpolicyapi.Rule{
					EndpointSelector: ciliumpolicyapi.EndpointSelector{
						LabelSelector: &ciliumslimmetav1.LabelSelector{
							MatchLabels: map[string]string{},
						},
					},
					// allow ingress from same namesapce
					Ingress: []ciliumpolicyapi.IngressRule{
						{
							IngressCommonRule: ciliumpolicyapi.IngressCommonRule{

								FromEndpoints: []ciliumpolicyapi.EndpointSelector{
									ciliumpolicyapi.NewESFromK8sLabelSelector("k8s.", &ciliumslimmetav1.LabelSelector{
										MatchLabels: map[string]string{
											"io.kubernetes.pod.namespace": metav1.NamespaceDefault,
										},
									}),
								},
							},
						},
					},
				},
				&ciliumpolicyapi.Rule{
					EndpointSelector: ciliumpolicyapi.EndpointSelector{
						LabelSelector: &ciliumslimmetav1.LabelSelector{
							MatchLabels: map[string]string{},
						},
					},
					// allow ingress from ingress namespace
					Ingress: []ciliumpolicyapi.IngressRule{
						{
							IngressCommonRule: ciliumpolicyapi.IngressCommonRule{
								FromEndpoints: []ciliumpolicyapi.EndpointSelector{
									ciliumpolicyapi.NewESFromK8sLabelSelector("k8s.", &ciliumslimmetav1.LabelSelector{
										MatchLabels: map[string]string{
											"io.kubernetes.pod.namespace": "ingressControllerNamespaceName",
										},
									}),
								},
							},
						},
					},
				},
			},
		}

		x, err := ciliumClientset.CiliumV2().CiliumNetworkPolicies(v1.NamespaceDefault).Create(ctx, ciliumPolicy, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		fmt.Print(x.APIVersion)
		return deployObjects(ctx, clientSet)
	},
	Readme: `
Welcome to "The Hive"
--------------------------------
	
	
		`,
}

func init() {
	Challenges = append(Challenges, challenge9)
}
