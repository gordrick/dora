package main

import (
	"bytes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gordrick/dora/pkg/ui"
	"github.com/gordrick/dora/pkg/utils"
	"os/exec"
	"strings"
)

func isServiceRunning() bool {
	cmd := exec.Command("launchctl", "list", "com.gordrick.dora")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}

	return strings.Contains(out.String(), "com.gordrick.dora")
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Dora Service Manager")

	startButton := widget.NewButton("Start Service", func() {
		ui.StartService()
	})

	stopButton := widget.NewButton("Stop Service", func() {
		ui.StopService()
	})

	logsArea := widget.NewMultiLineEntry()
	logsArea.SetText("Logs will be displayed here...")
	logsArea.Disable()

	scrollableLogsArea := container.NewScroll(logsArea)
	scrollableLogsArea.SetMinSize(fyne.NewSize(400, 200))

	refreshLogsButton := widget.NewButton("Refresh Logs", func() {
		logs, err := utils.ParseThreadLogFile("/var/log/dora_thread.log")
		if err != nil {
			logs = err.Error()
		}
		logsArea.SetText(logs)
	})

	buttonContainer := container.NewVBox(
		startButton,
		stopButton,
		refreshLogsButton,
	)

	content := container.NewBorder(buttonContainer, nil, nil, nil, scrollableLogsArea)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.ShowAndRun()
}
