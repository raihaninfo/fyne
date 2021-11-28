package main

import (
	"fmt"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Adb Desktop App")
	// w.SetContent(widget.NewLabel("Hello world"))
	w.Resize(fyne.NewSize(600, 300))
	btn := widget.NewButton("Show Devices", func() {
		commant, err := exec.Command("adb", "devices").Output()
		if err != nil {
			panic(err)
		}
		fmt.Print(string(commant))

	})

	// check := widget.NewCheck("Check", func(b bool) {
	// 	fmt.Println(b)
	// })
	// w.SetContent()

	// w.SetContent(btn)
	w.SetContent(btn)
	w.ShowAndRun()
}
