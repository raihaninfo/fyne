package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// light
	// a.Settings().SetTheme(theme.LightTheme())
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Dark and Light theme")

	w.Resize(fyne.NewSize(500, 400))

	label1 := widget.NewLabel("Fyne Themes")
	btn1 := widget.NewButton("Light", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})
	btn2 := widget.NewButton("Dark", func() {
		a.Settings().SetTheme(theme.DarkTheme())
	})
	btn3 := widget.NewButton("Exit", func() {
		a.Quit()
	})

	w.SetContent(
		container.NewVBox(
			label1,
			btn1,
			btn2,
			btn3,
		),
	)

	w.ShowAndRun()
}
