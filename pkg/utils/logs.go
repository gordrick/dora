package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ThreadLogEntry struct {
	Path  string `json:"path"`
	Mtime int64  `json:"mtime"`
}

func ParseThreadLogFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open log file: %v", err)
	}
	defer file.Close()

	var logEntries []ThreadLogEntry

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ", ")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid log format")
		}

		path := strings.TrimPrefix(parts[0], "Path: ")

		// Get the mtime part (remove "Mtime: " prefix)
		mtimeStr := strings.TrimPrefix(parts[1], "Mtime: ")

		mtime, err := strconv.ParseInt(mtimeStr, 10, 64)
		if err != nil {
			return "", fmt.Errorf("failed to parse mtime: %v", err)
		}

		// Add the parsed log entry to the slice
		logEntries = append(logEntries, ThreadLogEntry{
			Path:  path,
			Mtime: mtime,
		})

	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading log file: %v", err)
	}

	jsonData, err := json.MarshalIndent(logEntries, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal log entries to JSON: %v", err)
	}

	return string(jsonData), nil
}
