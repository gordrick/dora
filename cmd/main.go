package main

import (
	"fmt"
	"github.com/gordrick/dora/pkg/config"
	"github.com/gordrick/dora/pkg/daemon"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	var commandQueue = make(chan string, 100)

	commandQueue <- "echo Hello World"
	commandQueue <- "date"
	commandQueue <- "ls -la"

	go daemon.StartTimerThread(configuration.Directory, configuration.TimeInterval)
	go daemon.StartWorkerThread(commandQueue)
	select {}
}
