package main

import (
    "fmt"
    "os/exec"
)

func main() {

    cmd := exec.Command("bash", "bashScripts/install_dependencies.sh")
    output, err := cmd.Output()
    if err != nil { 
        fmt.Println("Error:", err)
    }
    fmt.Println(string(output))
    
    // Run ComfyUI
    cmd1 := exec.Command("bash", "bashScripts/run_comfy.sh")
    output1, err1 := cmd1.Output()
    if err1 != nil {
        fmt.Println("Error:", err1)
    }
    fmt.Println(string(output1))
}

