package main

import (
	"log"

	"github.com/pytimer/initsystem"
)

var serviceName = "docker.service"

func main() {
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		log.Fatal(err)
	}
	isExists, _ := initSystem.IsExists(serviceName)
	log.Printf("Check the %s exists: %t \n", serviceName, isExists)
	log.Printf("Stop the %s \n", serviceName)
	if err := initSystem.Stop(serviceName); err != nil {
		log.Fatalf("Failed to stop the %s: %s", serviceName, err)
	}
	log.Printf("Start the %s \n", serviceName)
	if err := initSystem.Start(serviceName); err != nil {
		log.Fatalf("Failed to stop the %s: %s", serviceName, err)
	}

	isActive, _ := initSystem.IsActive(serviceName)
	log.Printf("Check the %s active: %t \n", serviceName, isActive)
}
