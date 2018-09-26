package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func io(usageString string, path string) { // usage in MB/s
	size, error := strconv.ParseInt(usageString, 10, 64)
	byteArray := make([]byte, size*1000000)
	if error != nil {
		log.Fatal(error)
	}
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	second := time.Duration(time.Second)
	for {
		start := time.Now()
		f.Write(byteArray)
		f.Sync()
		elapsed := time.Since(start)

		timeTosleep := second - elapsed
		if timeTosleep > 0 {
			time.Sleep(timeTosleep)
		}
	}

}
