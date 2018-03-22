package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func startContainer(containerName string) {
	containInfo, err := getContainerInfoByName(containerName)
	if err != nil {
		log.Errorf("Get container %s info error %v", containerName, err)
		return
	}
	fmt.Println(containInfo)
}
