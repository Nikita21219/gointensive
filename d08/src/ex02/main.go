package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	schoolApp := app.New()
	w := schoolApp.NewWindow("School 21")
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
	fmt.Println("Hello world")
}
