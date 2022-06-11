package utils

import (
	"fmt"
	"os/exec"
)

func RunCMD(name string, arg string) bool {
	fmt.Println(name)
	fmt.Println(arg)
	cmd := exec.Command(name, arg)
	//cmd.Dir = path
	err := cmd.Run()
	if err != nil {
		return true
	}
	return false
}
