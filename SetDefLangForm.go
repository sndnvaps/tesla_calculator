package main

import (
	"github.com/sndnvaps/tesla_calculator/setting"
	"github.com/visualfc/goqt/ui"
	"os"
	"path"
)

//func SetDefLangForm(w *MainWindowForm)
func SetDefLangForm(w *MainWindowForm) {

	label := ui.NewLabel()
	label.SetText(Lang.Tr("setting.changedeflang"))

	comboBox := ui.NewComboBox()
	comboBox.AddItems(setting.Names)

	comboBox.SetCurrentIndex(int32(setting.LangsMapNum[setting.DefLang]))

	SetDefLangBtn := ui.NewPushButton()
	SetDefLangBtn.SetText(Lang.Tr("setting.saveBtn"))

	SetDefLangBtn.OnClicked(func() {
		ChooseLangText := comboBox.CurrentText()
		if ChooseLangText != setting.DefLang {
			setting.Cfg.Section("i18n").Key("defaultLang").SetValue(setting.LangsMap[ChooseLangText])
			os.MkdirAll(path.Dir(setting.CFG_CUSTOM_PATH), os.ModePerm)
			setting.Cfg.SaveTo(setting.CFG_CUSTOM_PATH)

			msgBox := ui.NewMessageBox()
			msgBox.SetText(Lang.Tr("setting.msgBox"))
			msgBox.Show()
		}
	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(label)
	hbox.AddWidget(comboBox)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(SetDefLangBtn)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddLayout(hbox2)

	widget := ui.NewWidgetWithParentFlags(w, ui.Qt_Window)
	widget.SetLayout(vbox)
	widget.SetWindowTitle(Lang.Tr("setting.WinTitle"))
	widget.Show()

}
