package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(cmd string, args []string) {
	if err := exec.Command(cmd, args...).Run(); err != nil {
		HandleErr(err)
	}
	fmt.Println("Success!")
}

func HandleErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
