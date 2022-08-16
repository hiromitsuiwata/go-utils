package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-ps"
)

func main() {
	processes, err := ps.Processes()
	if err != nil {
		log.Println("cannot get processes: " + err.Error())
	}

	fmt.Printf("%6s %6s %s\n", "PID", "PPID", "Executable")
	for _, proc := range processes {
		fmt.Printf("%6d %6d %s\n", proc.Pid(), proc.PPid(), proc.Executable())
	}

}

