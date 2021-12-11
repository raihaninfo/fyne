package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Dicee Game")
	w.Resize(fyne.NewSize(400, 400))

	// image
	img := canvas.NewImageFromFile("images/dice6.png")
	img.FillMode = canvas.ImageFillOriginal

	// button
	btn1 := widget.NewButton("Play", func() {
		ran := rand.Intn(6) + 1
		img.File = fmt.Sprintf("images/dice%d.png", ran)
		img.Refresh()

	})

	w.SetContent(
		// newVbox
		container.NewVBox(
			img,
			btn1,
		),
	)

	// show and run
	w.ShowAndRun()
}
