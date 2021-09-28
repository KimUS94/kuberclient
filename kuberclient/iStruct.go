package kuberclient

type NodeInfo struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Hostname string `json:"hostname"`

	Usage struct {
		Cpu_cores string `json:"cpu_cores"`
		Cpu_usage string `json:"cpu_usage"`
		Mem_cores string `json:"mem_cores"`
		Mem_usage string `json:"mem_usage"`
	}

	Allocatable struct {
		Cpu               string `json:"cpu"`
		Ephemeral_storage string `json:"ephemeral_storage"`
		Hugepages_1Gi     string `json:"hugepages_1Gi"`
		Hugepages_2Mi     string `json:"hugepages_2Mi"`
		Memory            string `json:"memory"`
		Pods              string `json:"pods"`
	}

	Capacity struct {
		Cpu               string `json:"cpu"`
		Ephemeral_storage string `json:"ephemeral_storage"`
		Hugepages_1Gi     string `json:"hugepages_1Gi"`
		Hugepages_2Mi     string `json:"hugepages_2Mi"`
		Memory            string `json:"memory"`
		Pods              string `json:"pods"`
	}

	Condition struct {
		NetworkUnavailable string `json:"NetworkUnavailable"`
		MemoryPressure     string `json:"MemoryPressure"`
		DiskPressure       string `json:"DiskPressure"`
		PIDPressure        string `json:"PIDPressure"`
		Ready              string `json:"Ready"`
	}

	SystemInfo struct {
		Architecture            string `json:"architecture"`
		BootID                  string `json:"bootID"`
		ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
		KernelVersion           string `json:"kernelVersion"`
		KubeProxyVersion        string `json:"kubeProxyVersion"`
		KubeletVersion          string `json:"kubeletVersion"`
		MachineID               string `json:"machineID"`
		OperatingSystem         string `json:"operatingSystem"`
		OsImage                 string `json:"osImage"`
		SystemUUID              string `json:"systemUID"`
	}
}

type nodeResource struct {
	cpu               string
	ephemeral_storage string
	hugepages_1Gi     string
	hugepages_2Mi     string
	memory            string
	pods              string
}

type nodeCondition struct {
	NetworkUnavailable string
	MemoryPressure     string
	DiskPressure       string
	PIDPressure        string
	Ready              string
}

type nodeSysteminfo struct {
	architecture            string
	bootID                  string
	containerRuntimeVersion string
	kernelVersion           string
	kubeProxyVersion        string
	kubeletVersion          string
	machineID               string
	operatingSystem         string
	osImage                 string
	systemUUID              string
}

type nodeUsage struct {
	cpu_cores string
	cpu_usage string
	mem_cores string
	mem_usage string
}
