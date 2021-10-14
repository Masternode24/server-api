package main

import (
	"fmt"
	"log"
	"os/exec"
)

func VerService() string {
	command := []string{"--version | awk '{print $2}' | head -1"}
	out, err := exec.Command("/root/polkadot/./target/release/polkadot'", command...).Output()
	if err != nil {
		fmt.Println("an error has occurred while checking")
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
