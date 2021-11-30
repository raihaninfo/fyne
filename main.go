package main

import (
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500, 400))

	command, _ := exec.Command("adb", "devices").Output()

	hello := widget.NewLabel("")
	raihan := widget.NewLabel("")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Show Devices", func() {
			hello.SetText(string(command))
		}),
		raihan,
		widget.NewButton("Reset", func() {
			hello.SetText("")
		}),
	))

	w.ShowAndRun()
}
