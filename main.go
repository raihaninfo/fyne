package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("New H Split")
	w.Resize(fyne.NewSize(500, 400))

	label1 := canvas.NewText("Text 1", color.White)
	label2 := canvas.NewText("Text 2", color.White)

	w.SetContent(
		container.NewHSplit(
			label1,
			label2,
		),
	)

	w.ShowAndRun()
}
