package main

import (
	"github.com/visualfc/goqt/ui"
	"os"
	//"log"
)

func main() {
	ui.RunEx(os.Args, func() {
		w := NewMainWindow()
		//w,err := NewMainForm()
		//if err != nil {
		//  log.Fatalln(err)
		// }
		w.Show()
	})
}
