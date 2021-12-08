package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Random Number Generator")

	// window size
	w.Resize(fyne.NewSize(400, 400))

	label1 := canvas.NewText("Rand Num Gen", color.White)
	label1.TextSize = 40

	btn1 := widget.NewButton("Generate", func() {
		// logic
		rand := rand.Intn(100)
		label1.Text = fmt.Sprint(rand)
		label1.Refresh()
	})

	// show
	w.SetContent(
		container.NewVBox(
			label1,
			btn1,
		),
	)

	w.ShowAndRun()
}
