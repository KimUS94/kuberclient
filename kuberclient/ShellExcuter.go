package kuberclient

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteCommad(cmdType string, args ...string) string {
	var res string

	cmd := exec.Command(cmdType, args...)
	output, err := cmd.Output()
	if err != nil {
		res = string("err")
	} else {
		res = string(output)
	}

	return res
}

func getNodeCount() int {
	//kubectl/get/nodes/-o=jsonpath={.items[*].metadata.name}

	cnt := 0
	output := ExecuteCommad("kubectl", "get", "nodes", "-o=jsonpath={.items[*].metadata.name}")

	if len(output) > 0 {
		arr := strings.Split(output, " ")
		cnt = len(arr)
	}

	return cnt
}

func getNodeName(idx int) string {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].metadata.name}", idx)
	output := ExecuteCommad("kubectl", "get", "nodes", jpath)

	return output
}

func getNodeAddress(idx int) string {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.addresses[0].address}", idx)
	output := ExecuteCommad("kubectl", "get", "nodes", jpath)

	return output
}

func getNodeHostname(idx int) string {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.addresses[1].address}", idx)
	output := ExecuteCommad("kubectl", "get", "nodes", jpath)

	return output
}

func getNodeAllocatable(idx int) *nodeResource {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.cpu}", idx)
	cpu := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.ephemeral-storage}", idx)
	ephemeral_storage := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.hugepages-1Gi}", idx)
	hugepages_1Gi := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.hugepages-2Mi}", idx)
	hugepages_2Mi := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.memory}", idx)
	memory := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.pods}", idx)
	pods := ExecuteCommad("kubectl", "get", "nodes", jpath)

	var allocatable nodeResource
	allocatable.cpu = cpu
	allocatable.ephemeral_storage = ephemeral_storage
	allocatable.hugepages_1Gi = hugepages_1Gi
	allocatable.hugepages_2Mi = hugepages_2Mi
	allocatable.memory = memory
	allocatable.pods = pods

	return &allocatable
}

func getNodeCapacity(idx int) *nodeResource {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.cpu}", idx)
	cpu := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.ephemeral-storage}", idx)
	ephemeral_storage := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.hugepages-1Gi}", idx)
	hugepages_1Gi := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.hugepages-2Mi}", idx)
	hugepages_2Mi := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.memory}", idx)
	memory := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.pods}", idx)
	pods := ExecuteCommad("kubectl", "get", "nodes", jpath)

	var capacity nodeResource
	capacity.cpu = cpu
	capacity.ephemeral_storage = ephemeral_storage
	capacity.hugepages_1Gi = hugepages_1Gi
	capacity.hugepages_2Mi = hugepages_2Mi
	capacity.memory = memory
	capacity.pods = pods

	return &capacity
}

func getNodeCondition(idx int) *nodeCondition {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[0].status}", idx)
	net := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[1].status}", idx)
	mem := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[2].status}", idx)
	disk := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[3].status}", idx)
	pid := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[4].status}", idx)
	ready := ExecuteCommad("kubectl", "get", "nodes", jpath)

	var condition nodeCondition
	condition.NetworkUnavailable = net
	condition.MemoryPressure = mem
	condition.DiskPressure = disk
	condition.PIDPressure = pid
	condition.Ready = ready

	return &condition
}

func getNodeSysteminfo(idx int) *nodeSysteminfo {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.architecture}", idx)
	architecture := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.bootID}", idx)
	bootID := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.containerRuntimeVersion}", idx)
	containerRuntimeVersion := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.kernelVersion}", idx)
	kernelVersion := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.kubeProxyVersion}", idx)
	kubeProxyVersion := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.kubeletVersion}", idx)
	kubeletVersion := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.machineID}", idx)
	machineID := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.operatingSystem}", idx)
	operatingSystem := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.osImage}", idx)
	osImage := ExecuteCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.systemUUID}", idx)
	systemUUID := ExecuteCommad("kubectl", "get", "nodes", jpath)

	var sysInfo nodeSysteminfo
	sysInfo.architecture = architecture
	sysInfo.bootID = bootID
	sysInfo.containerRuntimeVersion = containerRuntimeVersion
	sysInfo.kernelVersion = kernelVersion
	sysInfo.kubeProxyVersion = kubeProxyVersion
	sysInfo.kubeletVersion = kubeletVersion
	sysInfo.machineID = machineID
	sysInfo.operatingSystem = operatingSystem
	sysInfo.osImage = osImage
	sysInfo.systemUUID = systemUUID

	return &sysInfo
}

func getNodeUsage(name string) *nodeUsage {
	//kubectl top nodes --use-protocol-buffers

	var usage nodeUsage
	output := ExecuteCommad("kubectl", "top", "nodes", name, "--use-protocol-buffers")

	if len(output) <= 0 {
		return &usage
	}

	if output == "err" {
		usage.cpu_cores = "Unknown"
		usage.cpu_usage = "Unknown"
		usage.mem_cores = "Unknown"
		usage.mem_usage = "Unknown"
		return &usage
	}

	arr := strings.Split(output, "\n")
	if len(arr) < 2 {
		return &usage
	}

	output = arr[1]
	arr = strings.Fields(output)
	if len(arr) < 5 {
		return nil
	}

	usage.cpu_cores = arr[1]
	usage.cpu_usage = arr[2]
	usage.mem_cores = arr[3]
	usage.mem_usage = arr[4]

	return &usage

}
