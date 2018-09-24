package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var s []byte

func ram(usageString string) {
	fmt.Println(usageString)
	var size uint64
	if strings.HasSuffix(usageString, "%") {
		usageString = strings.TrimSuffix(usageString, "%")
		percent, _ := strconv.ParseUint(usageString, 10, 64)
		si := &syscall.Sysinfo_t{}

		err := syscall.Sysinfo(si)
		if err != nil {
			panic("Commander, we have a problem. syscall.Sysinfo:" + err.Error())
		}
		fmt.Println(si.Freeram)
		fmt.Println(si.Totalram)
		fmt.Println(si.Bufferram)
		fmt.Println(si.Sharedram)
		size = (si.Freeram + si.Bufferram + si.Sharedram) * percent / 100
	} else {
		size, _ = strconv.ParseUint(usageString, 10, 64)
	}
	fmt.Println(size)
	s = make([]byte, size, size) // do not work well...

	time.Sleep(time.Hour * 100000)
	fmt.Println(s)
}
