package guide

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()            //creat app instance
	w := a.NewWindow("Hello") // open a windows
	w.SetContent(widget.NewLabel("Hello World"))
	w.ShowAndRun()
	tidyUP()
}

func tidyUP() {
	fmt.Println("Exited!")
}
