package main

import (
	"fmt"
	"testing"
)

func TestLauncher(t *testing.T) {
	pid := launchCommand("./chaos --cpu --cpuusage 20")
	fmt.Println(pid)
}
