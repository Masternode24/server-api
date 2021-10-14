package main

import (
	"fmt"
	"log"
	"os/exec"
)

func VerService() string {
	command := []string{"version"}
	out, err := exec.Command("kvcli", command...).Output()
	if err != nil {
		fmt.Println("an error has occurred while checking")
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
