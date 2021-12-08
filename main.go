package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gradient")
	// size
	w.Resize(fyne.NewSize(500, 400))

	//
	btn1 := widget.NewButton("click Me", func() {

	})
	label1 := widget.NewLabel("Here is my test")

	box1 := container.NewHBox(
		btn1,
		label1,
	)

	w.SetContent(
		box1,
	)

	w.ShowAndRun()
}
