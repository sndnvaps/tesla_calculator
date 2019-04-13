package main

import (
	//"errors"
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

	w.btn1 = ui.NewPushButton()
	w.btn1.SetText(Lang.Tr("sphere.calBtn"))

	w.label_1 = ui.NewLabel()
	w.label_1.SetText(Lang.Tr("sphere.radius"))
	w.le1 = ui.NewLineEdit()

	w.label_2 = ui.NewLabel()
	w.label_2.SetText(Lang.Tr("sphere.outputCap"))
	w.le2 = ui.NewLineEdit()

	//设置 le2 为只读模式
	w.le2.SetReadOnly(true)

	w.btn1.OnClicked(func() {
		if IsValidDriver(w.le1) {
			output := calccapacitance(w.le1.Text())
			w.le2.SetText(output)
		} else {
			messagebox := ui.NewMessageBox()
			messagebox.SetText(Lang.Tr("sphere.msgBox"))
			messagebox.Show()
			w.le1.Clear()
			w.le2.Clear()
		}
	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(w.label_1)
	hbox.AddWidget(w.le1)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(w.label_2)
	hbox2.AddWidget(w.le2)

	hbox3 := ui.NewHBoxLayout()
	hbox3.AddWidget(w.btn1)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddLayout(hbox2)
	vbox.AddLayout(hbox3)

	w.SetLayout(vbox)

	w.SetWindowTitle(Lang.Tr("sphere.WinTitle"))
	return w, nil
}
