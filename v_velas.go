package main

import (
	"log"
	"os/exec"
)

func VerService() string {
	out, err := exec.Command("bash", "-c", "velas --version | awk '{print $2}' | tr -d '\n'").Output()
	if err != nil {
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
