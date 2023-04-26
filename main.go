package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	dependenciesFlag = ".dependencies_installed"
)

func runScript(scriptPath string, out chan<- string, err chan<- error) {
    cmd := exec.Command("bash", scriptPath)
    output, e := cmd.Output()
    if e != nil {
        err <- fmt.Errorf("failed to execute %s: %v", scriptPath, e)
    } else {
        out <- string(output)
    }
}

func main() {
	// check if dependencies have been installed before
	if _, err := os.Stat(dependenciesFlag); os.IsNotExist(err) {
		// install dependencies
		cmd := exec.Command("bash", "bashScripts/install_dependencies.sh")
		output, err := cmd.Output()
		if err != nil { 
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(output))

		// save flag indicating dependencies have been installed
		f, err := os.Create(dependenciesFlag)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		f.Close()
	}
	
	// Run ComfyUI and Port forward with ngrok concurrently
	out1 := make(chan string, 1)
	err1 := make(chan error, 1)
	out2 := make(chan string, 1)
	err2 := make(chan error, 1)


	go runScript("bashScripts/run_comfy.sh", out1, err1)
	go runScript("bashScripts/run_ngrok.sh", out2, err2)

	if err := <-err1; err != nil {
		fmt.Println("Error:", err)
	}
	if err := <-err2; err != nil {
		fmt.Println("Error:", err)
	}
	
	for i := 0; i < 2; i++ {
		select {
		case output1 := <-out1:
			fmt.Println(output1)
		case output2 := <-out2:
			fmt.Println(output2)
		case err1 := <-err1:
			fmt.Println("Error:", err1)
		case err2 := <-err2:
			fmt.Println("Error:", err2)
		}
	}
}
