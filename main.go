package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("Raihan")
	w.SetContent(widget.NewLabel("First Desktop app"))

	w.ShowAndRun()
}
