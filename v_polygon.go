package main

import (
	"fmt"
	"log"
	"os/exec"
)

func VerService() string {
	command := []string{"version | grep Version | head -1 | awk -F ':' '{print $2}' | tr -d ' '"}
	out, err := exec.Command("bor", command...).Output()
	if err != nil {
		fmt.Println("an error has occurred while checking the cpu")
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
