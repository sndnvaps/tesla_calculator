package main

import (
	"errors"
	"fmt"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
)

type SphereForm struct {
	*ui.QWidget
	btn1    *ui.QPushButton // 计算按钮
	le1     *ui.QLineEdit   //输入框
	le2     *ui.QLineEdit   //输出框
	label_1 *ui.QLabel      //
	label_2 *ui.QLabel
	label_3 *ui.QLabel
}

/*
//常量
var k=1.01;
//当输入的单位为mm的时候
var u=25.4;

//球形的半径假设为 381mm(即为15 Inch)
var r=381;
//电容的单位使用pf
//9e+9 = 9*10^9
var c=0.0;
c=(k*r/u/39.3701)/(9e+9)
c=c*(1e+15)
c=c/(1e+3)
*/
/*
 *@parame
 * radius -> 球的半径, 单位 mm
 *@return
 * 返回的是计算得到的电容 ,单位pf
 */
func calccapacitance(radius string) string {
	const k = 1.01
	const ki = 39.3701
	const u = 25.4 // 此用于将  mm -> inch
	var c float64
	r, _ := strconv.ParseFloat(radius, 32)
	c = (k * r / u / ki) / (9.0 * math.Pow(10, 9))
	c = c * (math.Pow(10, 15))
	c = c / (math.Pow(10, 3))

	return fmt.Sprintf("%0.6f", c)
}

func NewSphereForm() (*SphereForm, error) {
	w := &SphereForm{}
	w.QWidget = ui.NewWidget()

	file := ui.NewFileWithName(":/forms/sphere.ui")
	if !file.Open(ui.QIODevice_ReadOnly) {
		return nil, errors.New("error load ui")
	}

	loader := ui.NewUiLoader()
	formWidget := loader.Load(file)
	if formWidget == nil {
		return nil, errors.New("error load form widget")
	}

	w.btn1 = ui.NewPushButtonFromDriver(formWidget.FindChild("pushButton_1"))
	w.le1 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_1"))
	w.le2 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_2"))
	w.label_1 = ui.NewLabelFromDriver(formWidget.FindChild("label_1"))
	w.label_2 = ui.NewLabelFromDriver(formWidget.FindChild("label_2"))
	w.label_3 = ui.NewLabelFromDriver(formWidget.FindChild("label_3"))

	//设置 le2 为只读模式
	w.le2.SetReadOnly(true)

	w.btn1.OnClicked(func() {
		if IsValidDriver(w.le1) {
			output := calccapacitance(w.le1.Text())
			w.le2.SetText(output)
		} else {
			messagebox := ui.NewMessageBox()
			messagebox.SetText("输入的数有问题！")
			messagebox.Show()
			w.le1.Clear()
			w.le2.Clear()
		}
	})

	layout := ui.NewVBoxLayout()
	layout.AddWidget(formWidget)
	w.SetLayout(layout)

	w.SetWindowTitle("球形电容容量计算")
	return w, nil
}
