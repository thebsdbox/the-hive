package k3d

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/docker/go-connections/nat"
	cliutil "github.com/k3d-io/k3d/v5/cmd/util"
	l "github.com/k3d-io/k3d/v5/pkg/logger"
	k3drt "github.com/k3d-io/k3d/v5/pkg/runtimes"
	k3d "github.com/k3d-io/k3d/v5/pkg/types"
	"github.com/sirupsen/logrus"

	k3dCluster "github.com/k3d-io/k3d/v5/pkg/client"
	"github.com/k3d-io/k3d/v5/pkg/config"
	"github.com/k3d-io/k3d/v5/pkg/config/types"
	conf "github.com/k3d-io/k3d/v5/pkg/config/v1alpha4"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
)

func CreateCluster(name string) (*k3d.Cluster, error) {
	l.Logger.SetLevel(logrus.ErrorLevel) // Disabling the additional output from cluster creation
	ctx := context.TODO()
	rand.Seed(time.Now().UnixNano())

	serviceRange := fmt.Sprintf("10.%d.%d.0/17", rand.Intn(254), rand.Intn(127))
	logrus.Info("🔧 creating your cluster, this should take around 60 seconds")
	logrus.Infof("👨‍🔧 using unique service range [%s]", serviceRange)
	c := &conf.SimpleConfig{
		TypeMeta: types.TypeMeta{Kind: "Simple", APIVersion: "APIVersion:k3d.io/v1alpha4"},
		ObjectMeta: types.ObjectMeta{
			Name: name,
		},
		Servers: 1,
		Agents:  2,
		Image:   "the-hive.cloud:5000/rancher/k3s:v1.25.6-k3s1",
		Ports: []conf.PortWithNodeFilters{
			{
				Port: "30000-30010:30000-30010",
				NodeFilters: []string{
					"server:0:direct",
				},
			},
		},
		Options: conf.SimpleConfigOptions{
			K3dOptions: conf.SimpleConfigOptionsK3d{
				Wait:                true,
				DisableLoadbalancer: true,
			},
			K3sOptions: conf.SimpleConfigOptionsK3s{
				ExtraArgs: []conf.K3sArgWithNodeFilters{
					{
						Arg: "--flannel-backend=none",
						NodeFilters: []string{
							"server:0",
						},
					},
					{
						Arg: "--disable-network-policy",
						NodeFilters: []string{
							"server:0",
						},
					},
					{
						Arg: "--disable=traefik",
						NodeFilters: []string{
							"server:0",
						},
					},
					{
						Arg: "--disable=servicelb",
						NodeFilters: []string{
							"server:0",
						},
					},
					{
						Arg: "--disable=metrics-server",
						NodeFilters: []string{
							"server:0",
						},
					},
					{
						Arg: fmt.Sprintf("--service-cidr=%s", serviceRange),
						NodeFilters: []string{
							"server:0",
						},
					},
					{
						Arg: "--kube-apiserver-arg=service-node-port-range=30000-30010",
						NodeFilters: []string{
							"server:0",
						},
					},
				},
			},
		},
	}

	var exposeAPI *k3d.ExposureOpts

	// Apply config file values as defaults
	exposeAPI = &k3d.ExposureOpts{
		PortMapping: nat.PortMapping{
			Binding: nat.PortBinding{
				HostIP:   c.ExposeAPI.HostIP,
				HostPort: c.ExposeAPI.HostPort,
			},
		},
		Host: c.ExposeAPI.Host,
	}

	var freePort string
	port, err := cliutil.GetFreePort()
	freePort = strconv.Itoa(port)
	if err != nil || port == 0 {
		logrus.Warnf("Failed to get random free port: %+v", err)
		logrus.Warnf("Falling back to internal port %s (may be blocked though)...", k3d.DefaultAPIPort)
		freePort = k3d.DefaultAPIPort
	}
	exposeAPI.Binding.HostPort = freePort

	c.ExposeAPI = conf.SimpleExposureOpts{
		Host:     exposeAPI.Host,
		HostIP:   exposeAPI.Binding.HostIP,
		HostPort: exposeAPI.Binding.HostPort,
	}

	if err := config.ProcessSimpleConfig(c); err != nil {
		return nil, fmt.Errorf("error processing/sanitizing simple config: %v", err)
	}

	clusterConfig, err := config.TransformSimpleToClusterConfig(ctx, runtimes.SelectedRuntime, *c)
	if err != nil {
		return nil, fmt.Errorf("error processing/sanitizing simple config: %v", err)
	}
	logrus.Debugf("===== Merged Cluster Config =====\n%+v\n===== ===== =====\n", clusterConfig)

	clusterConfig, err = config.ProcessClusterConfig(*clusterConfig)
	if err != nil {
		return nil, fmt.Errorf("error processing cluster configuration: %v", err)
	}

	if err := config.ValidateClusterConfig(ctx, runtimes.SelectedRuntime, *clusterConfig); err != nil {
		return nil, fmt.Errorf("failed Cluster Configuration Validation: %v", err)
	}

	/**************************************
	 * Create cluster if it doesn't exist *
	 **************************************/

	// check if a cluster with that name exists already
	if _, err := k3dCluster.ClusterGet(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster); err == nil {
		k3dCluster.ClusterDelete(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, k3d.ClusterDeleteOpts{SkipRegistryCheck: true})

		return nil, fmt.Errorf("Failed to create cluster '%s' because a cluster with that name already exists", clusterConfig.Cluster.Name)
	}

	// create cluster
	if clusterConfig.KubeconfigOpts.UpdateDefaultKubeconfig {
		clusterConfig.ClusterCreateOpts.WaitForServer = true
	}
	//if err := k3dCluster.ClusterCreate(cmd.Context(), runtimes.SelectedRuntime, &clusterConfig.Cluster, &clusterConfig.ClusterCreateOpts); err != nil {
	if err := k3dCluster.ClusterRun(ctx, runtimes.SelectedRuntime, clusterConfig); err != nil {
		// rollback if creation failed
		if c.Options.K3dOptions.NoRollback { // TODO: move rollback mechanics to pkg/
			return nil, fmt.Errorf("cluster creation FAILED, rollback deactivated [%v]", err)
		}
		// rollback if creation failed
		logrus.Errorln("Failed to create cluster >>> Rolling Back")
		if err := k3dCluster.ClusterDelete(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, k3d.ClusterDeleteOpts{SkipRegistryCheck: true}); err != nil {
			return nil, fmt.Errorf("Cluster creation FAILED, also FAILED to rollback changes [%v]", err)
		}
		return nil, fmt.Errorf("Cluster creation FAILED, all changes have been rolled back![%v]", err)
	}
	logrus.Infof("Cluster '%s' created successfully!", clusterConfig.Cluster.Name)

	clusterConfig.KubeconfigOpts.SwitchCurrentContext = true

	logrus.Infof("Updating default kubeconfig 📄 with a new context for cluster %s", clusterConfig.Cluster.Name)
	if _, err := k3dCluster.KubeconfigGetWrite(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, "", &k3dCluster.WriteKubeConfigOptions{UpdateExisting: true, OverwriteExisting: true, UpdateCurrentContext: true}); err != nil {
		logrus.Warningln(err)

	}

	// Post cluster fixing of eBPF and cgroupsv2 (otherwise cilium will hang)
	nodes, err := k3dCluster.NodeList(ctx, runtimes.SelectedRuntime)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var execErr error

	for _, node := range nodes {
		if strings.HasSuffix(node.Name, "lb") {
			continue
		}
		logrus.Infof("adding eBPF 🐝  and 📂  cgroupv2 to %s", node.Name)
		wg.Add(1)

		go func(node *k3d.Node) {
			defer wg.Done()

			err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "bpffs", "-t", "bpf", "/sys/fs/bpf"})
			if err != nil {
				execErr = err
			}
			err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "--make-shared", "/sys/fs/bpf"})
			if err != nil {
				execErr = err
			}
			err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mkdir", "-p", "/run/cilium/cgroupv2"})
			if err != nil {
				execErr = err
			}
			err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "-t", "cgroup2", "none", "/run/cilium/cgroupv2"})
			if err != nil {
				execErr = err
			}
			err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "--make-shared", "/run/cilium/cgroupv2"})
		}(node)

	}
	wg.Wait()
	if execErr != nil {
		return nil, execErr
	}

	logrus.Infof("🧑‍💻  installing Cilium")

	cmd := exec.CommandContext(ctx, "kubectl", append([]string{"create", "-f", "-"})...)

	// Pass through our environment
	cmd.Env = os.Environ()
	// Pass through our std{out,err} and make our resolved buffer stdin.
	cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout
	cmd.Stdin = strings.NewReader(cilium)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// logrus.Infof("🧑‍💻  installing Tetragon")

	// cmd = exec.CommandContext(ctx, "kubectl", append([]string{"create", "-f", "-"})...)

	// // Pass through our environment
	// cmd.Env = os.Environ()
	// // Pass through our std{out,err} and make our resolved buffer stdin.
	// cmd.Stderr = os.Stderr
	// //cmd.Stdout = os.Stdout
	// cmd.Stdin = strings.NewReader(tetragon)
	// err = cmd.Run()
	// if err != nil {
	// 	return nil, err
	// }
	// cmd := exec.Command("cilium", "install", //,
	// 	"--helm-set", "hubble.relay.enabled=true",
	// 	"--helm-set", "hubble.ui.enabled=true",
	// 	"--helm-set", "kubeProxyReplacement=strict",
	// 	"--helm-set", "hubble.enabled=true")
	// if _, err := cmd.CombinedOutput(); err != nil {
	// 	return nil, err
	// }

	return &clusterConfig.Cluster, nil
}

// DeleteCluster will remove the existing cluster, based upon the cluster configuration
func DeleteCluster(ctx context.Context, cluster *k3d.Cluster) error {
	logrus.Info("🧽 deleting K3D cluster")
	if err := k3dCluster.ClusterDelete(ctx, runtimes.SelectedRuntime, cluster, k3d.ClusterDeleteOpts{SkipRegistryCheck: true}); err != nil {
		return fmt.Errorf("Cluster creation FAILED, also FAILED to rollback changes [%v]", err)
	}
	return nil
}
