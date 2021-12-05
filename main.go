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
	// window size
	w.Resize(fyne.NewSize(500, 400))

	rec := canvas.NewRectangle(color.White)

	rec.FillColor = color.NRGBA{R: 0, G: 255, B: 100, A: 244}
	rec.StrokeColor = color.NRGBA{R: 255, G: 0, B: 100, A: 244}
	rec.StrokeWidth = 5.0
	w.SetContent(rec)

	w.ShowAndRun()
}
