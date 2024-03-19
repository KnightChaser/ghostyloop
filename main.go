package main

import (
	"fmt"

	"github.com/terra-farm/go-virtualbox"
)

func createVirtualMachine(machineName string, machineVMFileLocation string) *virtualbox.Machine {
	// Create machine shell
	machine, err := virtualbox.CreateMachine(machineName, machineVMFileLocation)
	if err != nil {
		fmt.Printf("Error while creating machine: %s\n", err)
	} else {
		fmt.Printf("Machine created: %s(UUID: %v)\n", machine.Name, machine.UUID)
	}

	// Set machine settings like disks
	machine.CPUs = 2
	machine.Memory = 2048

	return machine
}

func main() {
	machineName := "ubuntu_test_2204_x64"
	machineLocation := "./ghostyloop/virtualbox/ubuntu_test_2204_x64/"
	machine := createVirtualMachine(machineName, machineLocation)
	fmt.Printf("Machine: %s\n", machine.Name)
}
