package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Icon")
	w.Resize(fyne.NewSize(500, 300))

	icon := widget.NewIcon(theme.ConfirmIcon())

	w.SetContent(icon)

	w.ShowAndRun()
}
