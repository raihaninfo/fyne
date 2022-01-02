package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Contact Form")
	w.Resize(fyne.NewSize(500, 400))

	title := canvas.NewText("Contact US", color.White)
	title.TextSize = 20
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Resize(fyne.NewSize(300, 35))
	title.Move(fyne.NewPos(50, 10))

	// new field name
	entry_name := widget.NewEntry()
	entry_name.SetPlaceHolder("Enter Your name")
	entry_name.Resize(fyne.NewSize(300, 35))
	entry_name.Move(fyne.NewPos(50, 50))

	// new field Email
	entry_email := widget.NewEntry()
	entry_email.SetPlaceHolder("Enter Your Email")
	entry_email.Resize(fyne.NewSize(300, 35))
	entry_email.Move(fyne.NewPos(50, 100))

	// new field Email
	entry_address := widget.NewEntry()
	entry_address.SetPlaceHolder("Enter Your Address")
	entry_address.Resize(fyne.NewSize(300, 35))
	entry_address.Move(fyne.NewPos(50, 150))

	// new field Email


	w.SetContent(
		container.NewWithoutLayout(
			title,
			entry_name,
			entry_email,
			entry_address,
			entry_message,
			submit_btn,
		),
	)
	w.ShowAndRun()
}
