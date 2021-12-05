package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Liene: canvas line")
	// window size
	w.Resize(fyne.NewSize(500, 400))

	line := canvas.NewLine(color.White)
	line.StrokeColor = color.NRGBA{R: 255, G: 0, B: 255, A: 255}
	line.StrokeWidth = 3
	w.SetContent(line)

	w.ShowAndRun()
}
