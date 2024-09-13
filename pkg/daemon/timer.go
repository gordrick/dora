package daemon

import (
	"fmt"
	"github.com/gordrick/dora/pkg/http"
	"github.com/osquery/osquery-go"
	"log"
	"os"
	"os/exec"
	"time"
)

// StartTimerThread starts the main timer thread for the daemon
func StartTimerThread(directory string, timeInterval uint, callBackURL string) {
	fmt.Println("Starting timer thread")
	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)

	for {
		select {
		case <-ticker.C:
			//checkFileModificationStats(directory)
			checkFileModificationStatsUsingOsquery(directory, callBackURL)
		}
	}
}

func checkDirectoryExistsAndReadable(directory string) {
	fmt.Println("Checking directory exists and is readable")
	cmd := exec.Command("sh", "-c", "ls "+directory)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func checkFileModificationStats(directory string) {
	fmt.Println("Checking file modification stats")
	//checkDirectoryExistsAndReadable(directory)
	cmd := exec.Command("sh", "-c", "ls -l "+directory)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	println(string(output))
}

// Function that checks file modification stats using osquery
func checkFileModificationStatsUsingOsquery(directory string, callBackURL string) {
	fmt.Println("Checking file modification stats with osquery")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %s", err)

	}
	client, err := osquery.NewClient(homeDir+"/.osquery/shell.em", 5*time.Second)
	if err != nil {
		log.Fatalf("Error connecting to osqueryd: %s", err)
	}
	defer client.Close()

	query := fmt.Sprintf(`SELECT path, mtime FROM file WHERE directory = '%s';`, directory)
	resp, err := client.Query(query)
	if err != nil {
		log.Fatalf("Error running query: %s", err)
	}

	if resp.Status.Code != 0 {
		log.Fatalf("Query failed: %s", resp.Status.Message)
	}

	for _, r := range resp.Response {
		logString := fmt.Sprintf("Path: %s, Mtime: %s\n", r["path"], r["mtime"])
		writeLogsToFile(logString)
	}
	err = http.SendTimerLogsCallback(callBackURL)
	if err != nil {
		log.Fatalf("Error sending logs to callback: %s", err)
	}
}

// Function that writes logs to daemon log file
func writeLogsToFile(logs string) {
	fmt.Println("Writing logs to file")
	f, err := os.OpenFile("/var/log/dora_thread.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(logs); err != nil {
		log.Fatalf("error writing to file: %v", err)
	}
}
