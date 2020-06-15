package main

import (
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
	"github.com/helmutkemper/iotmaker.docker/factoryDocker"
)

func main() {
	var err error
	var id string
	var ipAddress string
	var dockerSys *iotmakerDocker.DockerSystem
	var netData iotmakerDocker.ContainerNetworkDataList
	err, dockerSys = factoryDocker.NewClient()
	if err != nil {
		panic(err)
	}

	err, id = dockerSys.ContainerFindIdByNameContains("mongodb_pygocentrus_")
	if err != nil {
		panic(err)
	}

	err, netData = dockerSys.ContainerNetworkInspect(id)
	if err != nil {
		panic(err)
	}

	ipAddress = netData.GetIpAddressByNetworkName("pygocentrus_network")
	fmt.Printf("container ip: %v\n", ipAddress)
}
