package main

import (
	"errors"
	"reflect"
	"github.com/visualfc/goqt/ui"
)

type MainForm struct {
	*ui.QWidget
	btn1    *ui.QPushButton //
	btn2    *ui.QPushButton
	btn3	*ui.QPushButton
	btn4	*ui.QPushButton
	btn5	*ui.QPushButton
}

func IsValidDriver(v ui.Driver) bool {
	return !reflect.ValueOf(v).IsNil()
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
		version()
	})
	
	w.btn2.OnClicked(func() {
		version()
	})
	
	w.btn3.OnClicked(func() {
		SecCoilInfoForm()
	})

	w.btn4.OnClicked(func() {
		SparkLengthForm()
	})

	w.btn5.OnClicked(func() {
		CoefficientOfCouplinForm()
	})


	layout := ui.NewVBoxLayout()
	layout.AddWidget(formWidget)
	w.SetLayout(layout)

	w.SetWindowTitle("特斯拉线圈计算器")
	return w, nil
}
	
