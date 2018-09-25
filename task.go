package main

import (
	"log"
	"os/exec"
)

func launchCommand(command string) int {
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Waiting for command to finish...")
	//err = cmd.Wait()
	return cmd.Process.Pid
}
