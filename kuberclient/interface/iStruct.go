package kuberclient

type NodeInfo struct {
	name     string `json:"ame`
	address  string `json:address`
	hostname string `json:hostname`

	pods string `json:pods`

	usage struct {
		cpu_cores string `json:cpu_cores`
		cpu_usage string `json:cpu_usage`
		mem_cores string `json:mem_cores`
		mem_usage string `json:mem_usage`
	}

	allocatable struct {
		cpu               string `json:cpu`
		ephemeral_storage string `json:ephemeral_storage`
		hugepages_1Gi     string `json:hugepages_1Gi`
		hugepages_2Mi     string `json:hugepages_2Mi`
		memory            string `json:memory`
		pods              string `json:pods`
	}

	capacity struct {
		cpu               string `json:cpu`
		ephemeral_storage string `json:ephemeral_storage`
		hugepages_1Gi     string `json:hugepages_1Gi`
		hugepages_2Mi     string `json:hugepages_2Mi`
		memory            string `json:memory`
		pods              string `json:pods`
	}

	condition struct {
		NetworkUnavailable string `json:NetworkUnavailable`
		MemoryPressure     string `json:MemoryPressure`
		DiskPressure       string `json:DiskPressure`
		PIDPressure        string `json:PIDPressure`
		Ready              string `json:Ready`
	}

	systemInfo struct {
		architecture            string `json:architecture`
		bootID                  string `json:bootID`
		containerRuntimeVersion string `json:containerRuntimeVersion`
		kernelVersion           string `json:kernelVersion`
		kubeProxyVersion        string `json:kubeProxyVersion`
		kubeletVersion          string `json:kubeletVersion`
		machineID               string `json:machineID`
		operatingSystem         string `json:operatingSystem`
		osImage                 string `json:osImage`
		systemUID               string `json:systemUID`
	}
}
