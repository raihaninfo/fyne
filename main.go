package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("App title")
	w.Resize(fyne.NewSize(400, 200))

	// rgba color
	colorx := color.NRGBA{
		R: 0, G: 255, B: 0, A: 255,
	}

	textx := canvas.NewText("My text", colorx) //1st color, 2nd color

	w.SetContent(textx)

	w.ShowAndRun()
}
