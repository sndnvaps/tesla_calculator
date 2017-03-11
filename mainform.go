package main

import (
	"errors"
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

	file := ui.NewFileWithName(":/forms/mainform.ui")
	if !file.Open(ui.QIODevice_ReadOnly) {
		return nil, errors.New("error load ui")
	}

	loader := ui.NewUiLoader()
	formWidget := loader.Load(file)
	if formWidget == nil {
		return nil, errors.New("error load form widget")
	}

	w.btn1 = ui.NewPushButtonFromDriver(formWidget.FindChild("pushButton_1"))
	w.btn2 = ui.NewPushButtonFromDriver(formWidget.FindChild("pushButton_2"))
	w.btn3 = ui.NewPushButtonFromDriver(formWidget.FindChild("pushButton_3"))
	w.btn4 = ui.NewPushButtonFromDriver(formWidget.FindChild("pushButton_4"))
	w.btn5 = ui.NewPushButtonFromDriver(formWidget.FindChild("pushButton_5"))

	w.btn1.OnClicked(func() {

		primaryCoil, err := NewPrimaryCoilForm()
		if err != nil {
			log.Fatalln(err)
		}
		primaryCoil.Show()
	})

	w.btn2.OnClicked(func() {
		sphere, err := NewSphereForm()
		if err != nil {
			log.Fatalln(err)
		}
		sphere.Show()
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

	layout := ui.NewVBoxLayout()
	layout.AddWidget(formWidget)
	w.SetLayout(layout)

	w.SetWindowTitle("特斯拉线圈计算器")
	return w, nil
}
