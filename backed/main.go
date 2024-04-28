package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/gin-gonic/gin"
	"io"
	. "main/binary"
	. "main/request"
)

var myapp fyne.App
var w fyne.Window

func windows_init() {
	myapp = app.New()
	w = myapp.NewWindow("Load View")
	w.Resize(fyne.NewSize(720, 540))
	ginLog := widget.NewMultiLineEntry()
	w.SetContent(container.NewScroll(ginLog))
	out := io.MultiWriter(&FyneLogWriter{LogText: ginLog})
	R.Use(gin.LoggerWithWriter(out))

}

type FyneLogWriter struct {
	LogText *widget.Entry
}

func (w *FyneLogWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	w.LogText.SetText(w.LogText.Text + message)
	return len(p), nil
}

func main() {
	LogInit()
	DBInit()
	ServeInit()
	windows_init()
	defer F.Close()
	go func() {
		R.Run(":7234")
	}()
	w.ShowAndRun()
}
