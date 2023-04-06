package challenges

import (
	"context"
	"time"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var challengeTemplate = Challenge{
	Name:        "TBD ❓",
	Description: "This could be your opportunity ¯\\_(ツ)_/¯",
	AllowedTime: 4 * time.Minute,
	DeployFunc: func(ctx context.Context, clientSet *kubernetes.Clientset) error {

		_, err := clientSet.CoreV1().ConfigMaps(apiv1.NamespaceDefault).Create(ctx, backEndConfigMap, v1.CreateOptions{})
		if err != nil {
			return err
		}

		replicas := int32(2)
		frontEndDeployment.Spec.Replicas = &replicas
		deploymentsClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

		_, err = deploymentsClient.Create(ctx, frontEndDeployment, v1.CreateOptions{})
		if err != nil {
			return err
		}

		_, err = clientSet.CoreV1().Services(apiv1.NamespaceDefault).Create(ctx, frontEndService, v1.CreateOptions{})
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
	Challenges = append(Challenges, challengeTemplate)
}
