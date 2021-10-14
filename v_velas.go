package main

import (
	"fmt"
	"log"
	"os/exec"
)

func VerService() string {
	command := []string{"2>&1 | head -n 1"}
	out, err := exec.Command("velas", command...).Output()
	if err != nil {
		fmt.Println("an error has occurred while checking the cpu")
		log.Fatal(err)
	}
	verString := string(out[:])
	return verString
}
