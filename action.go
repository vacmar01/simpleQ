package main

import (
	"fmt"
	"os/exec"
	"os"
	"time"
)

func runScript(payload map[string]string) {
	cmd := exec.Command(payload["script"], payload["line"])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running script:", err)
		return
	}
}