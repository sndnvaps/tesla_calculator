package main

import (
	"github.com/visualfc/goqt/ui"
	"os"
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

		w := NewMainWindow()
		w.Show()
	})
}
