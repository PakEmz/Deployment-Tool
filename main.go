package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	os.Setenv("FYNE_THEME", "light")

	a := app.New()
	win := a.NewWindow("Deployment Tool")
	win.Resize(fyne.NewSize(600, 300))
	win.CenterOnScreen()
	username := widget.NewEntry()
	username.SetPlaceHolder("User Name of the Server...")
	ipaddress := widget.NewEntry()
	ipaddress.SetPlaceHolder("IP Address of the Server...")
	textArea := widget.NewMultiLineEntry()
	textArea.Resize(fyne.NewSize(100, 200))
	cancelbutton := widget.NewButton("CLOSE", func() {

		os.Exit(0)
	})

	deploybutton := widget.NewButton("DEPLOY", func() {
		usrname := username.Text
		ipaddr := ipaddress.Text
		if usrname != "" && ipaddr != "" {
			textArea.SetText("The Deployment is in Progress ....")
			result := Executeshell(usrname, ipaddr)
			textArea.SetText(result)
		} else {
			textArea.SetText("Please fill the details!")
		}
	})

	vlayout := container.New(layout.NewVBoxLayout(),
		username,
		ipaddress,
		container.New(layout.NewGridLayout(2), deploybutton, cancelbutton),
	)
	win.SetContent(container.New(layout.NewGridLayout(2), vlayout, textArea))

	win.ShowAndRun()
}

func Executeshell(usrname string, ipaddr string) string {

	script := `
	ls
	`

	c := exec.Command("ssh", usrname+"@"+ipaddr, "-i", "myserver.pem")
	var buf = new(bytes.Buffer)
	buf.WriteString(script)
	c.Stdin = buf

	b, e := c.Output()
	if e != nil {
		fmt.Println(e)
	}
	return (string(b))

}
