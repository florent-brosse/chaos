package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func makeFile(path string, usageString string) {
	var size int64
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	if strings.HasSuffix(usageString, "%") {
		usageString = strings.TrimSuffix(usageString, "%")
		percent, _ := strconv.ParseInt(usageString, 10, 64)
		var stat syscall.Statfs_t
		syscall.Statfs(path, &stat)

		// Available blocks * size per block = available space in bytes
		byteAvaillable := stat.Bavail * uint64(stat.Bsize)
		size = int64(byteAvaillable) * percent / 100
	} else {
		size, _ = strconv.ParseInt(usageString, 10, 64)
		size *= 1000000 //change from MB to B
	}

	if err := f.Truncate(size); err != nil {
		log.Fatal(err)
	}
}
