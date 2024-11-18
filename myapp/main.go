package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w1 := a.NewWindow("Hello1")
	w1.SetContent(widget.NewLabel("Hello1"))
	w1.Show()
	w1.SetMaster()

	w2 := a.NewWindow("Hello2")
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(300, 300))
	w2.Show()

	a.Run()
}
