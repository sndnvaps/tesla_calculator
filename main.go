package main

import (
	"github.com/visualfc/goqt/ui"
	"os"
)

func main() {
	ui.RunEx(os.Args, func() {
		w := NewMainWindow()
		w.Show()
	})
}
