package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type IP struct {
	Status   string
	Country  string
	Timezone string
	Isp      string
	Query    string
}

var ip IP

func main() {
	a := app.New()
	w := a.NewWindow("What is my ip")
	w.Resize(fyne.NewSize(600, 500))

	menuItem1 := fyne.NewMenuItem("Light", func() {
		Light := theme.LightTheme()
		a.Settings().SetTheme(Light)
	})
	menuItem2 := fyne.NewMenuItem("Dark", func() {
		Dark := theme.DarkTheme()
		a.Settings().SetTheme(Dark)
	})

	newMenu1 := fyne.NewMenu("File")
	newMenu2 := fyne.NewMenu("Edit")
	newMenu3 := fyne.NewMenu("View")
	newMenu := fyne.NewMenu("Theme", menuItem1, menuItem2)

	menu := fyne.NewMainMenu(newMenu1, newMenu2, newMenu3, newMenu)
	w.SetMainMenu(menu)

	img := canvas.NewImageFromFile("images/logo.jpg")
	img.FillMode = canvas.ImageFillOriginal

	status := widget.NewLabel("Status")
	statusValue := widget.NewLabel("..........")
	cuntry := widget.NewLabel("Cuntry")
	cuntryValue := widget.NewLabel("..........")
	timeZone := widget.NewLabel("Time Zone")
	timeZoneValue := widget.NewLabel("..........")
	isp := widget.NewLabel("Isp")
	ispValue := widget.NewLabel("..........")
	ip := widget.NewLabel("Internet ip")
	ipValue := widget.NewLabel("..........")
// 	run button
	btn := widget.NewButton("Run", func() {
		status, cuntry, time, isp, ip := myIp()
		statusValue.Text = status
		cuntryValue.Text = cuntry
		timeZoneValue.Text = time
		ispValue.Text = isp
		ipValue.Text = ip
		statusValue.Refresh()
		cuntryValue.Refresh()
		timeZoneValue.Refresh()
		ispValue.Refresh()
		ipValue.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2,
					status,
					statusValue,
				),
				container.NewGridWithColumns(2,
					cuntry,
					cuntryValue,
				),
				container.NewGridWithColumns(2,
					timeZone,
					timeZoneValue,
				),
				container.NewGridWithColumns(2,
					isp,
					ispValue,
				),
				container.NewGridWithColumns(2,
					ip,
					ipValue,
				),
				btn,
			),
			container.NewCenter(
				img,
			),
		),
	)
	w.ShowAndRun()
}

// Error Handling Function
func ErrorHandling(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Get isp full information
func myIp() (string, string, string, string, string) {
	req, err := http.Get("http://ip-api.com/json")
	ErrorHandling(err)

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	ErrorHandling(err)

	json.Unmarshal(body, &ip)
	return ip.Status, ip.Country, ip.Timezone, ip.Isp, ip.Query
}
