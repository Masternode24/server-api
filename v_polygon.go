package main

import (
	"log"
	"os/exec"
)

func VerService() string {
	out, err := exec.Command("bash", "-c", "bor version | grep Version | head -1 | awk -F ':' '{print $2}' | tr -d ' '").Output()
	if err != nil {
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
