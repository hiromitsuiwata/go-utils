package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	procs, err := process.Processes()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(procs); i++ {
		proc := procs[i]
		name, _ := proc.Name()
		cmdline, _ := proc.Cmdline()
		ctime, _ := proc.CreateTime()
		ctimeStr := time.Unix(ctime/1000, 0).Format("2006-01-02 15:04:05")

		exe, _ := proc.Exe()
		ppid, _ := proc.Ppid()
		numThreads, _ := proc.NumThreads()
		cpuPercent, _ := proc.CPUPercent()
		memoryInfo, _ := proc.MemoryInfo()
		var rss uint64 = 0
		if memoryInfo != nil {
			rss = memoryInfo.RSS / 1024 / 1024
		}

		fmt.Printf("%d %d %s %s %s\n", proc.Pid, ppid, name, ctimeStr, exe)
		fmt.Printf("%f %d %d\n", cpuPercent, rss, numThreads)
		fmt.Printf("%s\n\n", cmdline)
	}

	v, _ := mem.VirtualMemory()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// convert to JSON. String() is also implemented
	fmt.Println(v)
}
