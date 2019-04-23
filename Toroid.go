package main

import (
	"fmt"
	"github.com/sndnvaps/tesla_calculator/setting"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

type ToroidForm struct {
	*ui.QWidget
	btn1 *ui.QPushButton // 计算按钮
	le1  *ui.QLineEdit   //输入框
	le2  *ui.QLineEdit   //输入框
	le3  *ui.QLineEdit   //输出框

	label_1 *ui.QLabel // D1
	label_2 *ui.QLabel // D2
	label_3 *ui.QLabel // C

	picbox *ui.QLabel //Pic_box
}

/*
 *@parame
 * diameter1 -> 单位 mm
 * diameter2 -> 单位 mm
 *@return
 * 返回的是计算得到的电容 ,单位pf
 */
func TororidCalCapacitance(diameter1 string, diameter2 string) string {

	const pi = 3.14159
	const u = 25.4 // 此用于将  mm -> inch
	var k float64
	var c float64
	d1, _ := strconv.ParseFloat(diameter1, 32)
	d2, _ := strconv.ParseFloat(diameter2, 32)

	d1 = d1 / u
	d1 = d1 * math.Pow(10, 3)
	d1 = d1 + 0.5
	d1 = d1 / (math.Pow(10, 3))

	d2 = d2 / u
	d2 = d2 * math.Pow(10, 3)
	d2 = d2 + 0.5
	d2 = d2 / (math.Pow(10, 3))

	k = (2 * pi * pi * (d1 - d2) * (d2 / 2)) / (4 * pi)

	c = 2.8 * (1.2781 - d2/d1) * math.Sqrt(k)

	return fmt.Sprintf("%0.6f", c)
}

func NewToroidForm() (*ToroidForm, error) {

	w := &ToroidForm{}
	w.QWidget = ui.NewWidget()

	w.btn1 = ui.NewPushButton()
	w.btn1.SetText(Lang.Tr("toroid.calBtn"))

	w.label_1 = ui.NewLabel()
	w.label_1.SetText(Lang.Tr("toroid.inputD1"))
	w.le1 = ui.NewLineEdit()

	w.label_2 = ui.NewLabel()
	w.label_2.SetText(Lang.Tr("toroid.inputD2"))
	w.le2 = ui.NewLineEdit()

	w.label_3 = ui.NewLabel()
	w.label_3.SetText(Lang.Tr("toroid.outputCap"))
	w.le3 = ui.NewLineEdit()

	w.picbox = ui.NewLabel()

	//设置 le2 为只读模式
	w.le3.SetReadOnly(true)

	ImageBox := ui.NewPixmap()
	imgData, _ := setting.Asset("images/fig_toroid.png")
	ImageBox.LoadFromData(imgData)

	w.picbox.SetPixmap(ImageBox)

	w.btn1.OnClicked(func() {
		if (strings.Compare(w.le1.Text(), "") != 0) && (strings.Compare(w.le2.Text(), "") != 0) {
			output := TororidCalCapacitance(w.le1.Text(), w.le2.Text())
			w.le3.SetText(output)
		} else {
			messagebox := ui.NewMessageBox()
			messagebox.SetText(Lang.Tr("toroid.msgBox"))
			messagebox.Show()
			w.le1.Clear()
			w.le2.Clear()
			w.le3.Clear()
		}
	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(w.label_1)
	hbox.AddWidget(w.le1)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(w.label_2)
	hbox2.AddWidget(w.le2)

	hbox3 := ui.NewHBoxLayout()
	hbox3.AddWidget(w.label_3)
	hbox3.AddWidget(w.le3)

	hbox4 := ui.NewHBoxLayout()
	hbox4.AddWidget(w.btn1)

	hbox5 := ui.NewHBoxLayout()
	hbox5.AddWidget(w.picbox)

	vboxL := ui.NewVBoxLayout()
	vboxL.AddLayout(hbox)
	vboxL.AddLayout(hbox2)
	vboxL.AddLayout(hbox3)
	vboxL.AddLayout(hbox4)

	hboxMain := ui.NewHBoxLayout()
	hboxMain.AddLayout(vboxL)
	hboxMain.AddLayout(hbox5)

	w.SetLayout(hboxMain)

	w.SetWindowTitle(Lang.Tr("toroid.WinTitle"))
	return w, nil
}
