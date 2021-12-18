package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("CARD")
	w.Resize(fyne.NewSize(600, 500))

	card := widget.NewCard(
		"Hear is my card",
		"Heare is my subTitle",

		canvas.NewImageFromFile("images/image1.jpg"),
	)

	w.SetContent(card)
	w.ShowAndRun()
}
