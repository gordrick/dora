package daemon

import (
	"fmt"
	"log"
	"os/exec"
)

func StartWorkerThread(commandQueue chan string) {
	fmt.Println("Starting worker thread")
	for {
		select {
		case cmd := <-commandQueue:
			executeShellCommand(cmd)
		}
	}
}

// Function that executes shell commands
func executeShellCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing command '%s': %v", command, err)
	}
	fmt.Printf("Output of '%s':\n%s\n", command, string(output))
}
