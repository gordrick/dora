package main

import (
	"fmt"
	"github.com/gordrick/dora/pkg/config"
	"github.com/gordrick/dora/pkg/daemon"
	"github.com/gordrick/dora/pkg/http"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	var commandQueue = make(chan string, 100)

	go daemon.StartTimerThread(configuration.Directory, configuration.TimeInterval, configuration.CallBackURL)
	go daemon.StartWorkerThread(commandQueue)
	go http.StartServer(commandQueue)
	select {}
}
