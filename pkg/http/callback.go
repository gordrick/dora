package http

import (
	"bytes"
	"fmt"
	"github.com/gordrick/dora/pkg/utils"
	"log"
	"net/http"
)

func SendTimerLogsCallback(url string) error {
	logs, err := utils.ParseThreadLogFile("/var/log/dora_thread.log")
	if err != nil {
		return err
	}
	reqBody := bytes.NewBuffer([]byte(logs))
	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		return fmt.Errorf("failed to send logs to endpoint: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send logs, received status code: %d", resp.StatusCode)
	}

	log.Println("Successfully sent logs to", url)
	return nil
}
