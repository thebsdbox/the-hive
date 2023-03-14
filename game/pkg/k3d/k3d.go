package k3d

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/docker/go-connections/nat"
	cliutil "github.com/k3d-io/k3d/v5/cmd/util"
	k3drt "github.com/k3d-io/k3d/v5/pkg/runtimes"
	k3d "github.com/k3d-io/k3d/v5/pkg/types"
	"github.com/sirupsen/logrus"

	k3dCluster "github.com/k3d-io/k3d/v5/pkg/client"
	"github.com/k3d-io/k3d/v5/pkg/config"
	"github.com/k3d-io/k3d/v5/pkg/config/types"
	conf "github.com/k3d-io/k3d/v5/pkg/config/v1alpha4"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
)

func CreateCluster(name string) error {
	ctx := context.TODO()
	//c, err := conf.GetConfigByKind("simple")
	//c.GetAPIVersion()
	c := &conf.SimpleConfig{
		TypeMeta: types.TypeMeta{Kind: "Simple", APIVersion: "APIVersion:k3d.io/v1alpha4"},
		ObjectMeta: types.ObjectMeta{
			Name: name,
		},
		Servers: 1,
		Agents:  3,
		Image:   "docker.io/rancher/k3s:v1.25.6-k3s1",
		Ports: []conf.PortWithNodeFilters{
			{
				Port: "30000-30010:30000-30010",
				NodeFilters: []string{
					"server:0",
				},
			},
		},
		Options: conf.SimpleConfigOptions{
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
		return fmt.Errorf("error processing/sanitizing simple config: %v", err)
	}

	clusterConfig, err := config.TransformSimpleToClusterConfig(ctx, runtimes.SelectedRuntime, *c)
	if err != nil {
		return fmt.Errorf("error processing/sanitizing simple config: %v", err)
	}
	logrus.Infof("===== Merged Cluster Config =====\n%+v\n===== ===== =====\n", clusterConfig)

	clusterConfig, err = config.ProcessClusterConfig(*clusterConfig)
	if err != nil {
		return fmt.Errorf("error processing cluster configuration: %v", err)
	}

	if err := config.ValidateClusterConfig(ctx, runtimes.SelectedRuntime, *clusterConfig); err != nil {
		return fmt.Errorf("Failed Cluster Configuration Validation: ", err)
	}

	/**************************************
	 * Create cluster if it doesn't exist *
	 **************************************/

	// check if a cluster with that name exists already
	if _, err := k3dCluster.ClusterGet(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster); err == nil {
		k3dCluster.ClusterDelete(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, k3d.ClusterDeleteOpts{SkipRegistryCheck: true})

		return fmt.Errorf("Failed to create cluster '%s' because a cluster with that name already exists", clusterConfig.Cluster.Name)
	}

	// create cluster
	if clusterConfig.KubeconfigOpts.UpdateDefaultKubeconfig {
		clusterConfig.ClusterCreateOpts.WaitForServer = true
	}
	//if err := k3dCluster.ClusterCreate(cmd.Context(), runtimes.SelectedRuntime, &clusterConfig.Cluster, &clusterConfig.ClusterCreateOpts); err != nil {
	if err := k3dCluster.ClusterRun(ctx, runtimes.SelectedRuntime, clusterConfig); err != nil {
		// rollback if creation failed
		//l.Log().Errorln(err)
		if c.Options.K3dOptions.NoRollback { // TODO: move rollback mechanics to pkg/
			return fmt.Errorf("cluster creation FAILED, rollback deactivated.")
		}
		// rollback if creation failed
		logrus.Errorln("Failed to create cluster >>> Rolling Back")
		if err := k3dCluster.ClusterDelete(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, k3d.ClusterDeleteOpts{SkipRegistryCheck: true}); err != nil {
			logrus.Errorln(err)
			return fmt.Errorf("Cluster creation FAILED, also FAILED to rollback changes!")
		}
		return fmt.Errorf("Cluster creation FAILED, all changes have been rolled back!")
	}
	logrus.Infof("Cluster '%s' created successfully!", clusterConfig.Cluster.Name)

	clusterConfig.KubeconfigOpts.SwitchCurrentContext = true

	logrus.Infof("Updating default kubeconfig with a new context for cluster %s", clusterConfig.Cluster.Name)
	if _, err := k3dCluster.KubeconfigGetWrite(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, "", &k3dCluster.WriteKubeConfigOptions{UpdateExisting: true, OverwriteExisting: true, UpdateCurrentContext: true}); err != nil {
		logrus.Warningln(err)

	}

	// Post cluster fixing of eBPF and cgroupsv2 (otherwise cilium will hang)

	nodes, err := k3dCluster.NodeList(ctx, runtimes.SelectedRuntime)
	if err != nil {
		return err
	}
	for _, node := range nodes {
		if strings.HasSuffix(node.Name, "lb") {
			continue
		}
		err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "bpffs", "-t", "bpf", "/sys/fs/bpf"})
		if err != nil {
			return err
		}
		err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "--make-shared", "/sys/fs/bpf"})
		if err != nil {
			return err
		}
		err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mkdir", "-p", "/run/cilium/cgroupv2"})
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
		err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "-t", "cgroup2", "none", "/run/cilium/cgroupv2"})
		if err != nil {
			logrus.Error(err)
		}
		err = k3drt.SelectedRuntime.ExecInNode(ctx, node, []string{"mount", "--make-shared", "/run/cilium/cgroupv2"})
		if err != nil {
			return err
		}
	}

	return nil
}

// func applyOverrides(cfg conf.SimpleConfig) (conf.SimpleConfig, error) {

// 	/****************************
// 	 * Parse and validate flags *
// 	 ****************************/

// 	// -> API-PORT
// 	// parse the port mapping
// 	var (
// 		err       error
// 		exposeAPI *k3d.ExposureOpts
// 	)

// 	// Apply config file values as defaults
// 	exposeAPI = &k3d.ExposureOpts{
// 		PortMapping: nat.PortMapping{
// 			Binding: nat.PortBinding{
// 				HostIP:   cfg.ExposeAPI.HostIP,
// 				HostPort: cfg.ExposeAPI.HostPort,
// 			},
// 		},
// 		Host: cfg.ExposeAPI.Host,
// 	}

// 	// Overwrite if cli arg is set
// 	if ppViper.IsSet("cli.api-port") {
// 		if cfg.ExposeAPI.HostPort != "" {
// 			l.Log().Debugf("Overriding pre-defined kubeAPI Exposure Spec %+v with CLI argument %s", cfg.ExposeAPI, ppViper.GetString("cli.api-port"))
// 		}
// 		exposeAPI, err = cliutil.ParsePortExposureSpec(ppViper.GetString("cli.api-port"), k3d.DefaultAPIPort)
// 		if err != nil {
// 			return cfg, fmt.Errorf("failed to parse API Port spec: %w", err)
// 		}
// 	}

// 	// Set to random port if port is empty string
// 	if len(exposeAPI.Binding.HostPort) == 0 {
// 		var freePort string
// 		port, err := cliutil.GetFreePort()
// 		freePort = strconv.Itoa(port)
// 		if err != nil || port == 0 {
// 			logrus.Warnf("Failed to get random free port: %+v", err)
// 			logrus.Warnf("Falling back to internal port %s (may be blocked though)...", k3d.DefaultAPIPort)
// 			freePort = k3d.DefaultAPIPort
// 		}
// 		exposeAPI.Binding.HostPort = freePort
// 	}

// 	cfg.ExposeAPI = conf.SimpleExposureOpts{
// 		Host:     exposeAPI.Host,
// 		HostIP:   exposeAPI.Binding.HostIP,
// 		HostPort: exposeAPI.Binding.HostPort,
// 	}

// 	// -> VOLUMES
// 	// volumeFilterMap will map volume mounts to applied node filters
// 	volumeFilterMap := make(map[string][]string, 1)
// 	for _, volumeFlag := range ppViper.GetStringSlice("cli.volumes") {

// 		// split node filter from the specified volume
// 		volume, filters, err := cliutil.SplitFiltersFromFlag(volumeFlag)
// 		if err != nil {
// 			l.Log().Fatalln(err)
// 		}

// 		if strings.Contains(volume, k3d.DefaultRegistriesFilePath) && (cfg.Registries.Create != nil || cfg.Registries.Config != "" || len(cfg.Registries.Use) != 0) {
// 			l.Log().Warnf("Seems like you're mounting a file at '%s' while also using a referenced registries config or k3d-managed registries: Your mounted file will probably be overwritten!", k3d.DefaultRegistriesFilePath)
// 		}

// 		// create new entry or append filter to existing entry
// 		if _, exists := volumeFilterMap[volume]; exists {
// 			volumeFilterMap[volume] = append(volumeFilterMap[volume], filters...)
// 		} else {
// 			volumeFilterMap[volume] = filters
// 		}
// 	}

// 	for volume, nodeFilters := range volumeFilterMap {
// 		cfg.Volumes = append(cfg.Volumes, conf.VolumeWithNodeFilters{
// 			Volume:      volume,
// 			NodeFilters: nodeFilters,
// 		})
// 	}

// 	l.Log().Tracef("VolumeFilterMap: %+v", volumeFilterMap)

// 	// -> PORTS
// 	portFilterMap := make(map[string][]string, 1)
// 	for _, portFlag := range ppViper.GetStringSlice("cli.ports") {
// 		// split node filter from the specified volume
// 		portmap, filters, err := cliutil.SplitFiltersFromFlag(portFlag)
// 		if err != nil {
// 			l.Log().Fatalln(err)
// 		}

// 		// create new entry or append filter to existing entry
// 		if _, exists := portFilterMap[portmap]; exists {
// 			l.Log().Fatalln("Same Portmapping can not be used for multiple nodes")
// 		} else {
// 			portFilterMap[portmap] = filters
// 		}
// 	}

// 	for port, nodeFilters := range portFilterMap {
// 		cfg.Ports = append(cfg.Ports, conf.PortWithNodeFilters{
// 			Port:        port,
// 			NodeFilters: nodeFilters,
// 		})
// 	}

// 	l.Log().Tracef("PortFilterMap: %+v", portFilterMap)

// 	// --k3s-node-label
// 	// k3sNodeLabelFilterMap will add k3s node label to applied node filters
// 	k3sNodeLabelFilterMap := make(map[string][]string, 1)
// 	for _, labelFlag := range ppViper.GetStringSlice("cli.k3s-node-labels") {

// 		// split node filter from the specified label
// 		label, nodeFilters, err := cliutil.SplitFiltersFromFlag(labelFlag)
// 		if err != nil {
// 			l.Log().Fatalln(err)
// 		}

// 		// create new entry or append filter to existing entry
// 		if _, exists := k3sNodeLabelFilterMap[label]; exists {
// 			k3sNodeLabelFilterMap[label] = append(k3sNodeLabelFilterMap[label], nodeFilters...)
// 		} else {
// 			k3sNodeLabelFilterMap[label] = nodeFilters
// 		}
// 	}

// 	for label, nodeFilters := range k3sNodeLabelFilterMap {
// 		cfg.Options.K3sOptions.NodeLabels = append(cfg.Options.K3sOptions.NodeLabels, conf.LabelWithNodeFilters{
// 			Label:       label,
// 			NodeFilters: nodeFilters,
// 		})
// 	}

// 	l.Log().Tracef("K3sNodeLabelFilterMap: %+v", k3sNodeLabelFilterMap)

// 	// --runtime-label
// 	// runtimeLabelFilterMap will add container runtime label to applied node filters
// 	runtimeLabelFilterMap := make(map[string][]string, 1)
// 	for _, labelFlag := range ppViper.GetStringSlice("cli.runtime-labels") {

// 		// split node filter from the specified label
// 		label, nodeFilters, err := cliutil.SplitFiltersFromFlag(labelFlag)
// 		if err != nil {
// 			l.Log().Fatalln(err)
// 		}

// 		cliutil.ValidateRuntimeLabelKey(strings.Split(label, "=")[0])

// 		// create new entry or append filter to existing entry
// 		if _, exists := runtimeLabelFilterMap[label]; exists {
// 			runtimeLabelFilterMap[label] = append(runtimeLabelFilterMap[label], nodeFilters...)
// 		} else {
// 			runtimeLabelFilterMap[label] = nodeFilters
// 		}
// 	}

// 	for label, nodeFilters := range runtimeLabelFilterMap {
// 		cfg.Options.Runtime.Labels = append(cfg.Options.Runtime.Labels, conf.LabelWithNodeFilters{
// 			Label:       label,
// 			NodeFilters: nodeFilters,
// 		})
// 	}

// 	l.Log().Tracef("RuntimeLabelFilterMap: %+v", runtimeLabelFilterMap)

// 	// --env
// 	// envFilterMap will add container env vars to applied node filters
// 	envFilterMap := make(map[string][]string, 1)
// 	for _, envFlag := range ppViper.GetStringSlice("cli.env") {

// 		// split node filter from the specified env var
// 		env, filters, err := cliutil.SplitFiltersFromFlag(envFlag)
// 		if err != nil {
// 			l.Log().Fatalln(err)
// 		}

// 		// create new entry or append filter to existing entry
// 		if _, exists := envFilterMap[env]; exists {
// 			envFilterMap[env] = append(envFilterMap[env], filters...)
// 		} else {
// 			envFilterMap[env] = filters
// 		}
// 	}

// 	for envVar, nodeFilters := range envFilterMap {
// 		cfg.Env = append(cfg.Env, conf.EnvVarWithNodeFilters{
// 			EnvVar:      envVar,
// 			NodeFilters: nodeFilters,
// 		})
// 	}

// 	l.Log().Tracef("EnvFilterMap: %+v", envFilterMap)

// 	// --k3s-arg
// 	argFilterMap := make(map[string][]string, 1)
// 	for _, argFlag := range ppViper.GetStringSlice("cli.k3sargs") {

// 		// split node filter from the specified arg
// 		arg, filters, err := cliutil.SplitFiltersFromFlag(argFlag)
// 		if err != nil {
// 			l.Log().Fatalln(err)
// 		}

// 		// create new entry or append filter to existing entry
// 		if _, exists := argFilterMap[arg]; exists {
// 			argFilterMap[arg] = append(argFilterMap[arg], filters...)
// 		} else {
// 			argFilterMap[arg] = filters
// 		}
// 	}

// 	for arg, nodeFilters := range argFilterMap {
// 		cfg.Options.K3sOptions.ExtraArgs = append(cfg.Options.K3sOptions.ExtraArgs, conf.K3sArgWithNodeFilters{
// 			Arg:         arg,
// 			NodeFilters: nodeFilters,
// 		})
// 	}

// 	// --registry-create
// 	if ppViper.IsSet("cli.registries.create") {
// 		flagvalue := ppViper.GetString("cli.registries.create")
// 		fvSplit := strings.SplitN(flagvalue, ":", 2)
// 		if cfg.Registries.Create == nil {
// 			cfg.Registries.Create = &conf.SimpleConfigRegistryCreateConfig{}
// 		}
// 		cfg.Registries.Create.Name = fvSplit[0]
// 		if len(fvSplit) > 1 {
// 			exposeAPI, err = cliutil.ParsePortExposureSpec(fvSplit[1], "1234") // internal port is unused after all
// 			if err != nil {
// 				return cfg, fmt.Errorf("failed to registry port spec: %w", err)
// 			}
// 			cfg.Registries.Create.Host = exposeAPI.Host
// 			cfg.Registries.Create.HostPort = exposeAPI.Binding.HostPort
// 		}
// 	}

// 	// --host-alias
// 	hostAliasFlags := ppViper.GetStringSlice("hostaliases")
// 	if len(hostAliasFlags) > 0 {
// 		for _, ha := range hostAliasFlags {

// 			// split on :
// 			s := strings.Split(ha, ":")
// 			if len(s) != 2 {
// 				return cfg, fmt.Errorf("invalid format of host-alias %s (exactly one ':' allowed)", ha)
// 			}

// 			// validate IP
// 			ip, err := netaddr.ParseIP(s[0])
// 			if err != nil {
// 				return cfg, fmt.Errorf("invalid IP '%s' in host-alias '%s': %w", s[0], ha, err)
// 			}

// 			// hostnames
// 			hostnames := strings.Split(s[1], ",")
// 			for _, hostname := range hostnames {
// 				if err := k3dCluster.ValidateHostname(hostname); err != nil {
// 					return cfg, fmt.Errorf("invalid hostname '%s' in host-alias '%s': %w", hostname, ha, err)
// 				}
// 			}

// 			cfg.HostAliases = append(cfg.HostAliases, k3d.HostAlias{
// 				IP:        ip.String(),
// 				Hostnames: hostnames,
// 			})
// 		}
// 	}

// 	return cfg, nil
// }
