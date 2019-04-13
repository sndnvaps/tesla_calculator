package main

import (
	"github.com/visualfc/goqt/ui"
	"os"
)

func main() {
	ui.RunEx(os.Args, func() {

		InitI18n() //Init I18n first

		app := ui.Application()
		app.SetOrganizationName("sndnvaps.com")
		app.SetApplicationName(Lang.Tr("app.tc_desc"))

		Icon := ui.NewPixmap()
		IconData, _ := Asset("images/tesla.ico")
		Icon.LoadFromData(IconData)
		TC_Icon := ui.NewIconWithPixmap(Icon)
		app.SetWindowIcon(TC_Icon)

		w := NewMainWindow()
		w.Show()
	})
}
