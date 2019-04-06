//+build !windows
//+build arm
package main

import (
	"github.com/visualfc/goqt/ui"
)

func AddLibraryPath(app *ui.QApplication) {
		app.AddLibraryPath("plugins/platforms/linux/arm") //for linux arm platforms
}
