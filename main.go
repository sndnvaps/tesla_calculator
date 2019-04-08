package main

import (
	"github.com/visualfc/goqt/ui"
	"os"
	"runtime"
)

func main() {
	ui.RunEx(os.Args, func() {
		app := ui.Application()
		app.SetOrganizationName("sndnvaps.com")
		app.SetApplicationName("特斯拉线圈计算器")

		Icon := ui.NewPixmap()
		IconData, _ := Asset("images/tesla.ico")
		Icon.LoadFromData(IconData)
		TC_Icon := ui.NewIconWithPixmap(Icon)
		app.SetWindowIcon(TC_Icon)

		if (runtime.GOOS == "linux" && runtime.GOARCH == "arm") {
			AddLibraryPath(app)
		}
		// Test for Qt Plugins path; Begin test
		/*
			LibraryPaths := app.LibraryPaths()
			var LibraryPath string = string("")
			for i := 0; i < len(LibraryPaths); i++ {
			LibraryPath += LibraryPaths[i] + "\n"
			}

			messagebox := ui.NewMessageBox()
			messagebox.SetText(LibraryPath) // == linux
			messagebox.Show()
		*/
		//End test

		w := NewMainWindow()
		w.Show()
	})
}
