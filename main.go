package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Raihan")
	w.SetContent(widget.NewLabel("Hello world"))
	w.Resize(fyne.NewSize(600, 300))
	w.ShowAndRun()
}
