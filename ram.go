package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/mem"
)

var overall [][]int

func ram(usageString string) {
	fmt.Println(usageString)
	var size uint64
	if strings.HasSuffix(usageString, "%") {
		usageString = strings.TrimSuffix(usageString, "%")
		percent, _ := strconv.ParseUint(usageString, 10, 64)
		vm, err := mem.VirtualMemory()
		if err != nil {
			panic("Commander, we have a problem. mem.VirtualMemory:" + err.Error())
		}
		// fmt.Println(vm)
		size = vm.Available
		size = size * percent / 100
	} else {
		size, _ = strconv.ParseUint(usageString, 10, 64)
		size *= 1000000 //change from MB to B
	}
	fmt.Println(size)
	var sum uint64 = 0
	for sum < size {
		sum += 1000000
		a := make([]int, 1000000)
		for i := 0; i < len(a); i += 4096 {
			a[i] = 'x'
		}
		overall = append(overall, a)
	}
	time.Sleep(time.Hour * 100000)
	fmt.Println(len(overall))
}
