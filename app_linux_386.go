//+build !windows,386
package main

import (
	"github.com/visualfc/goqt/ui"
	"os"
	//"log"
)

func AddLibraryPath(app *ui.QApplication) {
		app.AddLibraryPath("plugins/platforms/linux/386") //for linux 386 platforms
}