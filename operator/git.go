package operator

import (
	"fmt"
	"os"
	"os/exec"
)

func HasGit() {
	cmd := exec.Command("git", "--version")
	_, err := cmd.Output()

	if err != nil {
		s := "You need to install git!\n"
		s += "Please, take a look on: https://git-scm.com/\n"
		fmt.Println(s)
		os.Exit(1)
	}
}
