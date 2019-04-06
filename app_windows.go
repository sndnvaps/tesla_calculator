//+build windows
package main

import (
	"github.com/visualfc/goqt/ui"
)

func AddLibraryPath(app *ui.QApplication) {
        app.AddLibraryPath("plugins\\platforms\\windows")
}