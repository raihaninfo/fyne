package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Main Menu")
	w.Resize(fyne.NewSize(500, 400))
	menuItem := &fyne.Menu{
		Label: "File",
		Items: nil,
	}
	menu := fyne.NewMainMenu(menuItem)

	w.SetMainMenu(menu)

	w.ShowAndRun()
}
