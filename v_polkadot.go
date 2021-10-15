package main

import (
	"log"
	"os/exec"
)

func VerService() string {
	out, err := exec.Command("bash", "-c", "/root/polkadot/./target/release/polkadot --version | awk '{print $2}' | head -1 | tr -d '\n'").Output()
	if err != nil {
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
