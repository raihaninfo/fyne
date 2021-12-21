package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("New H Split")
	w.Resize(fyne.NewSize(500, 400))

	label1 := canvas.NewText("Text 1", color.White)
	label2 := canvas.NewText("Text 2", color.White)
	w1 := widget.NewIcon(theme.CancelIcon())
	btn1 := widget.NewButton("Play", func() {

	})

	w.SetContent(
		container.NewHSplit(
			container.NewVSplit(
				label1,
				w1,
			),
			container.NewVBox(
				label2,
				btn1,
			),
		),
	)

	w.ShowAndRun()
}
