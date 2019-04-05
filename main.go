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
		w := NewMainWindow()
		w.Show()
	})
}
