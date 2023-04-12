package challenges

import (
	"context"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Details

var challenge1 = Challenge{
	Name:        "The labels one 🏷️ ",
	Description: "For some reason the NodePort isn't working  ¯\\_(ツ)_/¯",
	AllowedTime: 3 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset, r *rest.Config) error {

		replicas := int32(2)
		backEndDeployment.Spec.Replicas = &replicas

		backendEndService.Spec.Selector["app"] = "wrong" // ruhroh

		return deployObjects(ctx, clientSet)
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
	Challenges = append(Challenges, challenge1)
}
