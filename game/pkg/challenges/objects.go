package challenges

import (
	"embed"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

///////////////////////////////////////////////////////////////////
//
// Text
//
///////////////////////////////////////////////////////////////////

const readmeHeader = `Welcome to "The Hive"
--------------------------------

Enable hubble (optional)
-------------
kubectl apply -f ./manifests/hubble.yaml

Enable tetragon (optional)
-------------
kubectl apply -f ./manifests/tetragon.yaml

`

///////////////////////////////////////////////////////////////////
//
// Configmaps
//
///////////////////////////////////////////////////////////////////

// This configmap contains the static HTML that makes up the webpage
//go:embed assets/1index.html
//go:embed assets/1image.png
//go:embed assets/1style.css

//go:embed assets/2index.html
//go:embed assets/2script.js
//go:embed assets/2style.css
var f embed.FS

func trayagainConfigMap() *apiv1.ConfigMap {

	index, _ := f.ReadFile("assets/1index.html")
	style, _ := f.ReadFile("assets/1style.css")
	image, _ := f.ReadFile("assets/1image.png")

	return &apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "hive-html",
		},
		Data: map[string]string{
			"index.html": string(index),
			"style.css":  string(style),
		},
		BinaryData: map[string][]byte{
			"image1.png": image,
		},
	}
}

func winConfigMap() *apiv1.ConfigMap {

	index, _ := f.ReadFile("assets/2index.html")
	style, _ := f.ReadFile("assets/2style.css")
	script, _ := f.ReadFile("assets/2script.js")

	return &apiv1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "hive-html",
		},
		Data: map[string]string{
			"index.html": string(index),
			"style.css":  string(style),
			"script.js":  string(script),
		},
	}
}

var backEndConfigMap = &apiv1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name: "hive-html",
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

var middle = &apiv1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name: "middle-html-files",
	},
	Data: map[string]string{
		"index.html": `
		<!DOCTYPE html>
		<html>
		  <head>
			<meta http-equiv="refresh" content="0; url='https://hive.default/'" />
		  </head>
		  <body>
			<p>Loading!</p>
		  </body>
		</html>
`,
	},
}

///////////////////////////////////////////////////////////////////
//
// Deployments
//
///////////////////////////////////////////////////////////////////

var backEndDeployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name: "back-end",
	},
	Spec: appsv1.DeploymentSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": "backend",
			},
		},
		Template: apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "backend",
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
									Name: "hive-html",
								},
							},
						},
					},
				},
			},
		},
	},
}

var middleEndDeployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name: "middle-end",
	},
	Spec: appsv1.DeploymentSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": "middle",
			},
		},
		Template: apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "middle",
				},
			},
			Spec: apiv1.PodSpec{
				Containers: []apiv1.Container{
					{
						Name:  "web",
						Image: "alpine/socat:1.7.4.4",
						Ports: []apiv1.ContainerPort{
							{
								Name:          "http",
								Protocol:      apiv1.ProtocolTCP,
								ContainerPort: 31337,
							},
						},
						Args: []string{"tcp-listen:31337,fork,reuseaddr", "tcp-connect:backend.default:80"},
					},
				},
			},
		},
	},
}

var frontEndDeployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name: "front-end",
	},
	Spec: appsv1.DeploymentSpec{
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": "front",
			},
		},
		Template: apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "front",
				},
			},
			Spec: apiv1.PodSpec{
				Containers: []apiv1.Container{
					{
						Name:  "web",
						Image: "alpine/socat:1.7.4.4",
						Ports: []apiv1.ContainerPort{
							{
								Name:          "http",
								Protocol:      apiv1.ProtocolTCP,
								ContainerPort: 80,
							},
						},
						Args: []string{"tcp-listen:80,fork,reuseaddr", "tcp-connect:middle.default:31337"},
					},
				},
			},
		},
	},
}

//

///////////////////////////////////////////////////////////////////
//
// Services
//
///////////////////////////////////////////////////////////////////

var backendEndService = &apiv1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "backend",
		Namespace: apiv1.NamespaceDefault,
		// Labels:    GetLabels(),
	},
	Spec: apiv1.ServiceSpec{
		Type: apiv1.ServiceTypeClusterIP,
		Ports: []apiv1.ServicePort{
			{
				Name:       "web",
				TargetPort: intstr.FromInt(80),
				Port:       80,
				Protocol:   "TCP",
				//NodePort:   30000,
			},
		},
		Selector: map[string]string{
			"app": "backend",
		},
	},
}

var middleEndService = &apiv1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "middle",
		Namespace: apiv1.NamespaceDefault,
		// Labels:    GetLabels(),
	},
	Spec: apiv1.ServiceSpec{
		Type: apiv1.ServiceTypeClusterIP,
		Ports: []apiv1.ServicePort{
			{
				Name: "tcpforward",
				//TargetPort: intstr.FromInt(31337),
				Port:     31337,
				Protocol: "TCP",
			},
		},
		Selector: map[string]string{
			"app": "middle",
		},
	},
}

var frontEndService = &apiv1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "front",
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
			"app": "front",
		},
	},
}
