package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("Custom app icon in title bar")
	w.Resize(fyne.NewSize(500, 400))

	// app icon
	w.SetIcon(theme.FyneLogo())

	w.ShowAndRun()
}
