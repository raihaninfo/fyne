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
	w := a.NewWindow("Login From")
	a.Settings().SetTheme(theme.LightTheme())
	// resize window
	w.Resize(fyne.NewSize(500, 400))
	label := widget.NewLabel("")
	form := widget.NewForm(
		widget.NewFormItem("Username", widget.NewEntry()),
		widget.NewFormItem("Password", widget.NewPasswordEntry()),
	)
	// working and cancel
	form.OnCancel = func() {
		label.Text = "Canceled"
		label.Refresh()
	}
	form.OnSubmit = func() {
		label.Text = "Submited"
		label.Refresh()
	}

	w.SetContent(
		container.NewVBox(
			form,
			label,
		),
	)

	w.ShowAndRun()
}
