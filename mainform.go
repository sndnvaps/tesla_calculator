package main

import (
	//"errors"
	"github.com/Unknwon/i18n"
	"github.com/sndnvaps/tesla_calculator/setting"
	"github.com/visualfc/goqt/ui"
	"log"
	"reflect"
)

type MainWindowForm struct {
	*ui.QMainWindow
	widget *MainForm //我们要加载 MainForm 到这里
}

type MainForm struct { //此为widget,
	*ui.QWidget
	btn1 *ui.QPushButton //
	btn2 *ui.QPushButton
	btn3 *ui.QPushButton
	btn4 *ui.QPushButton
	btn5 *ui.QPushButton
}

func IsValidDriver(v ui.Driver) bool {
	return !reflect.ValueOf(v).IsNil()
}

func NewMainWindow() *MainWindowForm {

	w := &MainWindowForm{}
	w.QMainWindow = ui.NewMainWindow()
	w.InstallEventFilter(w)

	w.widget, _ = NewMainForm()

	w.SetCentralWidget(w.widget)
	w.createActions()

	return w
}

func (w *MainWindowForm) createActions() {

	aboutAct := ui.NewActionWithTextParent(Lang.Tr("help.about"), w)
	aboutAct.OnTriggered(func() {
		ui.QMessageBoxAbout(w, Lang.Tr("help.aboutTC"), Lang.Tr("help.aboutTC_desc"))
	})
	aboutQtAct := ui.NewActionWithTextParent(Lang.Tr("help.aboutQt"), w)
	aboutQtAct.OnTriggered(func() { ui.QApplicationAboutQt() })

	helpMenu := w.MenuBar().AddMenuWithTitle(Lang.Tr("help.help"))
	helpMenu.AddAction(aboutAct)
	helpMenu.AddSeparator()
	helpMenu.AddAction(aboutQtAct)

	w.SetWindowTitle(Lang.Tr("app.tc_desc"))
}

func NewMainForm() (*MainForm, error) {

	w := &MainForm{}
	w.QWidget = ui.NewWidget()

	w.SetFixedWidth(331)
	w.SetFixedHeight(233)

	w.btn1 = ui.NewPushButton()
	w.btn1.SetText(Lang.Tr("main.calPriInfo"))

	w.btn2 = ui.NewPushButton()
	w.btn2.SetText(Lang.Tr("main.calTopCap"))

	w.btn3 = ui.NewPushButton()
	w.btn3.SetText(Lang.Tr("main.calSecInfo"))

	w.btn4 = ui.NewPushButton()
	w.btn4.SetText(Lang.Tr("main.CalSparkLength"))

	w.btn5 = ui.NewPushButton()
	w.btn5.SetText(Lang.Tr("main.CalCoilCoupling"))

	w.btn1.OnClicked(func() {

		primaryCoil, err := NewPrimaryCoilForm()
		if err != nil {
			log.Fatalln(err)
		}
		primaryCoil.Show()
	})

	w.btn2.OnClicked(func() {

		topload, err := NewToploadForm()
		if err != nil {
			log.Fatalln(err)
		}
		topload.Show()
	})

	w.btn3.OnClicked(func() {
		SecCoilInfoForm()
	})

	w.btn4.OnClicked(func() {
		SparkLengthForm()
	})

	w.btn5.OnClicked(func() {
		cc, err := NewCoefficientForm()
		if err != nil {
			log.Fatalln(err)
		}
		cc.Show()
	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(w.btn1)
	hbox.AddWidget(w.btn2)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(w.btn3)
	hbox2.AddWidget(w.btn4)

	hbox3 := ui.NewHBoxLayout()
	hbox3.AddWidget(w.btn5)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddLayout(hbox2)
	vbox.AddLayout(hbox3)

	w.SetLayout(vbox)
	w.SetWindowTitle(Lang.Tr("app.tc_desc"))
	return w, nil

}
