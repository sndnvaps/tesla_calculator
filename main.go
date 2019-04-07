package main

import (
	"github.com/visualfc/goqt/ui"
	"os"
	//"log"
)

func main() {
	ui.RunEx(os.Args, func() {
		app := ui.Application()
		app.SetOrganizationName("GoQt")
		app.SetApplicationName("Application Example")

		Icon := ui.NewPixmap()
		IconData, _ := Asset("images/tesla.ico")
		Icon.LoadFromData(IconData)
		TC_Icon := ui.NewIconWithPixmap(Icon)
		app.SetWindowIcon(TC_Icon)
		
        // Test for Qt Plugins path; Begin test
		/*
		LibraryPaths := app.LibraryPaths()
		var LibraryPath string = string("")
		for i := 0; i < len(LibraryPaths); i++ {
		LibraryPath += LibraryPaths[i] + "\n"
		}
		messagebox := ui.NewMessageBox()
		messagebox.SetText(LibraryPath)
		messagebox.Show()
		*/
        //End test

		w := NewMainWindow()
		w.Show()
	})
}
