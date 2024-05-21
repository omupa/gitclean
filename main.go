package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "branch")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(stdout))
}
