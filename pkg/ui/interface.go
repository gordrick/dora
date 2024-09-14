package ui

import (
	"bytes"
	"fmt"
	"os/exec"
)

func StartService() {
	output, err := runWithSudo("launchctl load /Library/LaunchDaemons/com.gordrick.dora.plist")
	if err != nil {
		fmt.Println("Error starting service:", err)
	} else {
		fmt.Println("Service started successfully:", output)
	}
}

func StopService() {
	output, err := runWithSudo("launchctl unload /Library/LaunchDaemons/com.gordrick.dora.plist")

	if err != nil {
		fmt.Println("Error stopping service:", err)
	} else {
		fmt.Println("Service stopped successfully:", output)
	}
}
func runWithSudo(command string) (string, error) {
	osascript := fmt.Sprintf(`
		do shell script "%s" with administrator privileges
	`, command)

	cmd := exec.Command("osascript", "-e", osascript)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
