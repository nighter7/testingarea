package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"	
)

func main() {
	// Getting os name
	fmt.Println(runtime.GOOS)
	
	// Executing commands
	cmd := exec.Command("sleep", "55")	
	
	stdout, erro := cmd.CombinedOutput()
	
	if erro != nil {
		fmt.Printf("stderr: %s\n", erro)
	}	
	
	fmt.Printf("stdout: %s\n", stdout)	
}
