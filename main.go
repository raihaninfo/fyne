package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Loading image from internet")
	w.Resize(fyne.NewSize(500, 400))

	img, err := fyne.LoadResourceFromURLString("https://picsum.photos/seed/picsum/200/300")
	if err != nil {
		fmt.Println(err.Error())
	}
	// app icon
	image := canvas.NewImageFromResource(img)
	w.SetContent(image)

	w.ShowAndRun()
}
