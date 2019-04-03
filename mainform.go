package main

import (
	//"errors"
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

	aboutAct := ui.NewActionWithTextParent("&About", w)
	aboutAct.SetStatusTip("Show the application's About box")
	aboutAct.OnTriggered(func() {
		ui.QMessageBoxAbout(w, "关于特斯拉线圈计算器",
			w.Tr("特斯拉线圈计算器是使用 goqt 来编写的工具 "+
				"用于计算制作特斯拉线圈各个过程中使用到的相关参数, "+
				"目的是简化复杂的计算流程."))
	})
	aboutQtAct := ui.NewActionWithTextParent("About &Qt", w)
	aboutQtAct.SetStatusTip("Show the Qt library's About box")
	aboutQtAct.OnTriggered(func() { ui.QApplicationAboutQt() })

	helpMenu := w.MenuBar().AddMenuWithTitle("&Help")
	helpMenu.AddAction(aboutAct)
	helpMenu.AddSeparator()
	helpMenu.AddAction(aboutQtAct)

	w.SetWindowTitle("特斯拉线圈计算器")
	w.StatusBar().ShowMessage("Ready")
}

func NewMainForm() (*MainForm, error) {

	w := &MainForm{}
	w.QWidget = ui.NewWidget()

    w.SetFixedWidth(331)
    w.SetFixedHeight(233)

    w.btn1 = ui.NewPushButton()
    w.btn1.SetText("计算初始参数")
    
    w.btn2 = ui.NewPushButton()
    w.btn2.SetText("计算顶端电容")

    w.btn3 = ui.NewPushButton()
    w.btn3.SetText("计算次级线圈参数")

    w.btn4 = ui.NewPushButton()
    w.btn4.SetText("估算电弧长度")

    w.btn5 = ui.NewPushButton()
    w.btn5.SetText("计算初级与次级的耦合度")


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
	w.SetWindowTitle("特斯拉线圈计算器")
	return w, nil
    
}
