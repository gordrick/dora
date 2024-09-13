package http

import (
	"encoding/json"
	"fmt"
	"github.com/gordrick/dora/pkg/utils"
	"io"
	"log"
	"net/http"
)

var CommandQueue chan string

type CommandsRequest struct {
	Commands []string `json:"commands"`
}

func StartServer(commandQueue chan string) {
	CommandQueue = commandQueue
	http.HandleFunc("/logs", logsHandler)
	http.HandleFunc("/commands", commandHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}

// logsHandler gets logs from the log file formats it to json returns it
func logsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	logs, err := utils.ParseThreadLogFile("/var/log/dora_thread.log")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(logs))
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var commandsRequest CommandsRequest
	if err := json.Unmarshal(body, &commandsRequest); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	if len(commandsRequest.Commands) == 0 {
		http.Error(w, "No commands provided", http.StatusBadRequest)
		return
	}

	for _, command := range commandsRequest.Commands {
		select {
		case CommandQueue <- command:
			fmt.Fprintf(w, "Command '%s' added to queue\n", command)
		default:
			http.Error(w, "Command queue is full", http.StatusServiceUnavailable)
			return
		}
	}
}
