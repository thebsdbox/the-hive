package kind

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	kindconfigv1alpha4 "sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cmd"
)

const localNodeImage = "the-hive.cloud:5000/kindest/node:v1.25.3"

var provider *cluster.Provider

func CreateKind(name string) error {
	// nodePorts := []kindconfigv1alpha4.PortMapping{
	// 	{
	// 		ContainerPort: 30000,
	// 		HostPort:      30000,
	// 	},
	// }
	clusterConfig := kindconfigv1alpha4.Cluster{
		Networking: kindconfigv1alpha4.Networking{
			IPFamily:          kindconfigv1alpha4.IPv4Family,
			DisableDefaultCNI: true, // remove the default networks
		},
		Nodes: []kindconfigv1alpha4.Node{
			{
				Role:  kindconfigv1alpha4.ControlPlaneRole,
				Image: localNodeImage, // User our local image
				//	ExtraPortMappings: nodePorts,
			},
			{
				Role:  kindconfigv1alpha4.WorkerRole,
				Image: localNodeImage,
			},
			{
				Role:  kindconfigv1alpha4.WorkerRole,
				Image: localNodeImage,
			},
			{
				Role:  kindconfigv1alpha4.WorkerRole,
				Image: localNodeImage,
			},
		},
	}

	// imagePath := os.Getenv("E2E_IMAGE_PATH")

	provider = cluster.NewProvider(cluster.ProviderWithLogger(cmd.NewLogger()), cluster.ProviderWithDocker())
	clusters, err := provider.List()
	found := false
	for x := range clusters {
		if clusters[x] == name {
			found = true
		}
	}
	if found {
		log.Error("Cluster already exists ")
	} else {
		err := provider.Create(name, cluster.CreateWithV1Alpha4Config(&clusterConfig))
		if err != nil {
			log.Error(err)
			return DeleteKind(name)
		}
		cmd := exec.Command("kubectl", "create", "configmap", "--namespace", "kube-system", "kubevip", "--from-literal", "range-global=172.18.103.10-172.18.103.30")
		if _, err := cmd.CombinedOutput(); err != nil {
			DeleteKind(name)
			return err
		}
		cmd = exec.Command("kubectl", "create", "-f", "https://raw.githubusercontent.com/kube-vip/kube-vip-cloud-provider/main/manifest/kube-vip-cloud-controller.yaml")
		if _, err := cmd.CombinedOutput(); err != nil {
			DeleteKind(name)
			return err
		}
		cmd = exec.Command("kubectl", "create", "-f", "https://kube-vip.io/manifests/rbac.yaml")
		if _, err := cmd.CombinedOutput(); err != nil {
			DeleteKind(name)
			return err
		}
		cmd = exec.Command("docker", "run", "--network", "host", "--rm", "ghcr.io/kube-vip/kube-vip:v0.5.11", "manifest", "daemonset", "--services", "--inCluster", "--arp", "--interface", "eth0", "|", "kubectl", "apply", "-f", "-")
		if _, err := cmd.CombinedOutput(); err != nil {
			// DeleteKind(name)
			// return err
			log.Error(err)
		}

	}
	// loadImageCmd := load.NewCommand(cmd.NewLogger(), cmd.StandardIOStreams())
	// loadImageCmd.SetArgs([]string{"--name", "services", imagePath})
	// err = loadImageCmd.Execute()
	// if err != nil {
	// 	log.Error(err)
	// 	return deleteKind()
	// }

	log.Infof("💤 sleeping for a few seconds to let controllers start")
	time.Sleep(time.Second * 5)

	err = installCNI(name)
	if err != nil {
		DeleteKind(name)
		return err
	}
	return nil
}

func DeleteKind(name string) error {
	log.Info("🧽 deleting Kind cluster")
	return provider.Delete(name, "")
}

func updateHelm() error {
	log.Infof("🧑‍💻 configuring helm for Cilium")
	cmd := exec.Command("helm", "repo", "add", "cilium", "https://helm.cilium.io/")
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func installCNI(clusterName string) error {
	address, err := getLeaderIP(fmt.Sprintf("%s-control-plane", clusterName))
	if err != nil {
		log.Error(address)
		return err
	}

	// err = updateHelm()
	// if err != nil {
	// 	return err
	// }
	log.Infof("🧑‍💻 installing Cilium with Kubernets host [%s]", address)
	cmd := exec.Command("cilium", "install",
		"--helm-set", "kubeProxyReplacement=strict",
		"--helm-set", fmt.Sprintf("k8sServiceHost=%s", trimQuotes(address)),
		"--helm-set", "k8sServicePort=6443")
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func getLeaderIP(leaderName string) (string, error) {
	log.Infof("🕵️ Finding address of Control Plane node [%s]", leaderName)

	// find the control plane instance and retrieve its IP address
	cmd := exec.Command(
		"docker", "inspect", leaderName,
		"--format", "'{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'")
	cmdOut := new(bytes.Buffer)
	cmd.Stdout = cmdOut
	err := cmd.Run()
	return strings.TrimSuffix(cmdOut.String(), "\n"), err
}
