package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gradient")
	// size
	w.Resize(fyne.NewSize(500, 400))

	//gradient
	gradient := canvas.NewLinearGradient(color.Black, color.White, 56)
	w.SetContent(gradient)

	w.ShowAndRun()
}
