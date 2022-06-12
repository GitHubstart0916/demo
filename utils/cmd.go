package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

const ShellToUse = "bash"
const Path = "/home/start_0916/SE/ROS/team08-proj/robot-ros"
const Shell = "./mk.sh"

func RunCMD(arg string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	ret := "'" + arg + "'"

	command := Shell + " " + Path + " " + ret
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Run()

	fmt.Println(stdout.String())
}
