package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("App title")
	w.Resize(fyne.NewSize(400, 200))

	img := canvas.NewImageFromFile("images/fyne.png")

	w.SetContent(img)

	w.ShowAndRun()
}
