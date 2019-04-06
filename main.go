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
		//func NewIconWithPixmap(pixmap *QPixmap) *QIcon {
		//func NewIconWithFilename(fileName string) *QIcon {
		//func (q *QApplication) SetWindowIcon(icon *QIcon) {
		Icon := ui.NewPixmap()
		IconData, _ := Asset("images/tesla.ico")
		Icon.LoadFromData(IconData)
		TC_Icon := ui.NewIconWithPixmap(Icon)
		app.SetWindowIcon(TC_Icon)
		//func (q *QApplication) AddLibraryPath(value string)
		
        AddLibraryPath(app)

		LibraryPaths := app.LibraryPaths()
		var LibraryPath string = string("")
		for i := 0; i < len(LibraryPaths); i++ {
		LibraryPath += LibraryPaths[i] + "\n"
		}
		messagebox := ui.NewMessageBox()
		messagebox.SetText(LibraryPath)
		messagebox.Show()

		w := NewMainWindow()
		w.Show()
	})
}
