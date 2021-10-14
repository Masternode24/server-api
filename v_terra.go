package main

import (
	"fmt"
	"log"
	"os/exec"
)

func VerTerra() string {
	command := []string{"version --long | grep version | awk -F ':' '{print $2}' | head -1 | tr -d ' '"}
	out, err := exec.Command("terrad", command...).Output()
	if err != nil {
		fmt.Println("an error has occurred while checking the cpu")
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
