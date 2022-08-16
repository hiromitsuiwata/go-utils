package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-ps"
)

func main() {
	fmt.Println("Hello")
	processes, err := ps.Processes()
	if err != nil {
		log.Println("cannot get processes: " + err.Error())
	}

	for _, proc := range processes {
		fmt.Printf("%6d %6d %s\n", proc.Pid(), proc.PPid(), proc.Executable())
	}

}
