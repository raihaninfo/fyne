package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("App title")

	// window size
	w.Resize(fyne.NewSize(400, 200))

	url, err := url.Parse("https://facebook.com")
	if err != nil {
		panic(err)
	}
	// hyperlink
	hyperLink := widget.NewHyperlink("visit Me ", url)

	w.SetContent(hyperLink)
	w.ShowAndRun()

}
