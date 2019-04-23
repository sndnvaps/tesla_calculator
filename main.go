package main

import (
	"github.com/sndnvaps/tesla_calculator/setting"
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
		IconData, _ := setting.Asset("images/tesla.ico")
		Icon.LoadFromData(IconData)
		TCIcon := ui.NewIconWithPixmap(Icon)
		app.SetWindowIcon(TCIcon)

		w := NewMainWindow()
		w.Show()
	})
}
