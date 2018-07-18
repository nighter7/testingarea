package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// Getting os name
	fmt.Println(runtime.GOOS)

	// Executing commands
	cmd := exec.Command(`tasklist`)

	stdout, erro := cmd.CombinedOutput()

	if erro != nil {
		fmt.Printf("stderr: %s\n", erro)
	}

	fmt.Printf("stdout: %s\n", stdout)
}
