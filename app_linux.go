//+build !windows
//+build arm
//+build 386
package main

import (
     "github.com/visualfc/goqt/ui"
)

func AddLibraryPath(app *ui.QApplication) {
    app.AddLibraryPath("../plugins")
}
