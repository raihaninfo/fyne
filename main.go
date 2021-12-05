package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Titel app")
	w.Resize(fyne.NewSize(400, 300))

	circle1 := canvas.NewCircle(color.NRGBA{R: 0, G: 0, B: 255, A: 100})
	circle1.StrokeColor = color.White
	circle1.StrokeWidth = 3
	

	w.SetContent(circle1)

	w.ShowAndRun()
}
