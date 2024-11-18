package guide

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func UpdateTime(clock *widget.Label) {
	formatted := time.Now().Format("2006-01-02 15:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock") // 新建一个窗口

	clock := widget.NewLabel("Clock_Test")
	UpdateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			UpdateTime(clock)
		}
	}()
	w.ShowAndRun()
}
