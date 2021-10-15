package main

import (
	"log"
	"os/exec"
)

func VerService() string {
	out, err := exec.Command("bash", "-c", "gaiad version").Output()
	if err != nil {
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
