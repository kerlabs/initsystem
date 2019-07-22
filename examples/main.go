package main

import (
	"log"

	"github.com/pytimer/initsystem"
)

func main() {
	initSystem, err := initsystem.GetInitSystem()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(initSystem.IsActive("sshd.service"))
}
