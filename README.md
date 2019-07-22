# initsystem

[![GoDoc](https://godoc.org/github.com/pytimer/initsystem?status.svg)](https://godoc.org/github.com/pytimer/initsystem)

It is the library that detect a supported init system.

This library reference the [kubeadm initsystem package](https://github.com/kubernetes/kubernetes/tree/master/cmd/kubeadm/app/util/initsystem).

## Supported system

- Systemd

## Usage

```go
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
```