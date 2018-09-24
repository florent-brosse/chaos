package main

import (
	"runtime"
	"strconv"
	"sync"
	"time"
)

func cpu(percentString string) {
	var wg sync.WaitGroup
	numCPU := runtime.NumCPU()
	percent, _ := strconv.ParseFloat(percentString, 64)
	wg.Add(numCPU)
	for i := 0; i < numCPU; i++ {
		go func() {
			defer wg.Done()
			var x float64 = (100 - percent) / percent
			for {
				sum := 1
				start := time.Now()
				for i := 0; i < 10000000; i++ {
					sum = sum + i
				}
				if sum > 10000000 {
					sum = 1
				}
				elapsed := time.Since(start)
				t := time.Duration(int64(float64(elapsed.Nanoseconds()) * x))
				time.Sleep(t)
			}
		}()
	}
	wg.Wait()
}
