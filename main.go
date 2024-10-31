package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container" // Import container package
	"fyne.io/fyne/v2/widget"    // Import widget package
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(container.NewVBox(widget.NewLabel("Hello World!")))
	w.ShowAndRun()
}
