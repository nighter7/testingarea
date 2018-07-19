package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"io/ioutil"
)

func main() {
	// Getting os name
	fmt.Println(runtime.GOOS)

	// Executing commands	
	txt := []byte(`tasklist | find "chrome.exe"`)
	err := ioutil.WriteFile("tmpfile.bat", txt, 0644)
	
	if err != nil {
		fmt.Printf("stderr: %s\n", err)
	}
	
	cmd := exec.Command("tmpfile.bat")

	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("stderr: %s\n", err)
	}

	fmt.Printf("stdout: %s\n", stdout)
}
