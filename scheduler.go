package main

import (
	"os/exec"
)

func main() {
	// app := "echo"
	// arg1 := "Running My first Commandss in Golang"

	cmd := exec.Command("echo", "HOla")
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	println(string(stdout))
}
