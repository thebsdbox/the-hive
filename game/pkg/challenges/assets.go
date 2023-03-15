package challenges

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This configmap contains the static HTML that makes up the webpage

var configMap = &apiv1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name: "static-html-files",
	},
	Data: map[string]string{
		"index.html": `
<!DOCTYPE html>
<html lang="en">
<head>
<title>Welcome</title>
</head>
<body>

<h1>Looks like you finally made it!</h1>
<img src="https://uploads-ssl.webflow.com/5d2b950d9ea87fc61f0c1f3e/5daa2f818bfff61e250fd7da_ezgif.com-resize.gif" alt="place holder">

</body>
</html>
`,
	},
}

var deployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name: "demo-deployment",
	},
	Spec: appsv1.DeploymentSpec{
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
						VolumeMounts: []apiv1.VolumeMount{
							{
								Name:      "html-files",
								ReadOnly:  true,
								MountPath: "/usr/share/nginx/html/",
							},
						},
					},
				},
				Volumes: []apiv1.Volume{
					{
						Name: "html-files",
						VolumeSource: apiv1.VolumeSource{
							ConfigMap: &apiv1.ConfigMapVolumeSource{
								LocalObjectReference: apiv1.LocalObjectReference{
									Name: "static-html-files",
								},
							},
						},
					},
				},
			},
		},
	},
}
