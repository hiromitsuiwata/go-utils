package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	processes, err := Processes()
	if err != nil {
		log.Println("cannot get processes: " + err.Error())
	}

	fmt.Printf("%6s %6s %s\n", "PID", "PPID", "Executable")
	for _, proc := range processes {
		fmt.Printf("%6d %6d %21s %s\n", proc.Pid(), proc.PPid(), proc.CreationTime().Format(time.RFC3339), proc.Executable())
	}

}
