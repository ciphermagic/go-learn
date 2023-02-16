package gui

import (
	"ciphermagic.cn/go-micro-service/chatserver/client"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartUi(c client.Client) {
	a := app.New()
	loginWindow := a.NewWindow("登录")
	input := widget.NewEntry()
	input.Resize(fyne.NewSize(24, 5))
	label := widget.NewLabel("Please input your name: ")
	button := widget.NewButton("login", func() {
		if len(input.Text) > 0 {
			c.SetName(input.Text)
			label.Hidden = true

			input.SetText("")
			input.Hidden = true
			changeWindow(loginWindow, c)
		}
	})
	loginWindow.SetContent(container.NewVBox(label, input, button))
	loginWindow.Resize(fyne.NewSize(24, 24))
	loginWindow.ShowAndRun()
}

func changeWindow(window fyne.Window, c client.Client) {
	history := widget.NewMultiLineEntry()
	history.Disable()
	input := widget.NewEntry()
	send := widget.NewButton("send", func() {
		if len(input.Text) > 0 {
			fmt.Println("Send start")
			c.SendMess(input.Text)
			input.SetText("")
		}
	})
	content := container.New(layout.NewVBoxLayout(), history, input, send)
	content.Resize(fyne.NewSize(480, 320))
	window.SetContent(content)
	window.Resize(fyne.NewSize(480, 320))

	go func() {
		for msg := range c.InComing() {
			AddMessage(history, msg.Name, msg.Message)
		}
	}()
}

func AddMessage(history *widget.Entry, name string, message string) {
	history.SetText(history.Text + "\n" + name + ": " + message)
}
