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

	go daemon.StartTimerThread(configuration.Directory, configuration.TimeInterval)
	select {}
}
