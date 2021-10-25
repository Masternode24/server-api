package main

import (
	"log"
	"os/exec"
)

func UpService() string {
	out, err := exec.Command("bash", "-c", "terrad status | jq | grep catching_up | awk '{print $2}' | tr -d '\n'").Output()
	if err != nil {
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
