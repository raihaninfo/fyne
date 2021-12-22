package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Custom app icon in title bar")
	w.Resize(fyne.NewSize(500, 400))

	icon, err := fyne.LoadResourceFromPath("icon/app_icon.png")
	if err != nil {
		fmt.Println(err.Error())
	}
	// app icon
	a.SetIcon(icon)

	w.ShowAndRun()
}
