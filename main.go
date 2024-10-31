package main

import (
	"time"

	"fyne.io/fyne/v2/app"    // Import container package
	"fyne.io/fyne/v2/widget" // Import widget package
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("The time is : 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}
