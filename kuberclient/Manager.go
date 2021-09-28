package kuberclient

func getNodeInfo() *map[string]NodeInfo {
	nodes := map[string]NodeInfo{}

	cnt := getNodeCount()

	for idx := 0; idx < cnt; idx++ {
		name := getNodeName(idx)
		addr := getNodeAddress(idx)
		host := getNodeHostname(idx)
		usag := getNodeUsage(name)
		allc := getNodeAllocatable(idx)
		capa := getNodeCapacity(idx)
		cndi := getNodeCondition(idx)
		sysi := getNodeSysteminfo(idx)

		var node NodeInfo
		node.Name = name
		node.Address = addr
		node.Hostname = host

		node.Usage.Cpu_cores = usag.cpu_cores
		node.Usage.Cpu_usage = usag.cpu_usage
		node.Usage.Mem_cores = usag.mem_cores
		node.Usage.Mem_usage = usag.mem_usage

		node.Allocatable.Cpu = allc.cpu
		node.Allocatable.Ephemeral_storage = allc.ephemeral_storage
		node.Allocatable.Hugepages_1Gi = allc.hugepages_1Gi
		node.Allocatable.Hugepages_2Mi = allc.hugepages_2Mi
		node.Allocatable.Memory = allc.memory
		node.Allocatable.Pods = allc.pods

		node.Capacity.Cpu = capa.cpu
		node.Capacity.Ephemeral_storage = capa.ephemeral_storage
		node.Capacity.Hugepages_1Gi = capa.hugepages_1Gi
		node.Capacity.Hugepages_2Mi = capa.hugepages_2Mi
		node.Capacity.Memory = capa.memory
		node.Capacity.Pods = capa.pods

		node.Condition.NetworkUnavailable = cndi.NetworkUnavailable
		node.Condition.MemoryPressure = cndi.MemoryPressure
		node.Condition.DiskPressure = cndi.DiskPressure
		node.Condition.PIDPressure = cndi.PIDPressure
		node.Condition.Ready = cndi.Ready

		node.SystemInfo.Architecture = sysi.architecture
		node.SystemInfo.BootID = sysi.bootID
		node.SystemInfo.ContainerRuntimeVersion = sysi.containerRuntimeVersion
		node.SystemInfo.KernelVersion = sysi.kernelVersion
		node.SystemInfo.KubeProxyVersion = sysi.kubeProxyVersion
		node.SystemInfo.KubeletVersion = sysi.kubeletVersion
		node.SystemInfo.MachineID = sysi.machineID
		node.SystemInfo.OperatingSystem = sysi.operatingSystem
		node.SystemInfo.OsImage = sysi.osImage
		node.SystemInfo.SystemUUID = sysi.systemUUID

		nodes[name] = node
	}

	return &nodes

}
