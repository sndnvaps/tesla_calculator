package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

// 用于计算电弧的长度和功率
/* 电弧长度 L : 英寸 --> Length
 * 功率     P : 瓦特 --> Power
 * 关系:
 *  L = 1.7 * sqrt(P) = 1.7 * Math.Sqrt(P)
 *  P = (L/1.7)^2     = Math.Pow((L/1.7),2.0)
 *
 * -------------------------------------------
 * 电弧长度 L : 厘米
 * 功率     P : 瓦特
 *  P = (L/1.7*2.54)^2
 *  L = 1.7 * Sqrt(P) * 2.54
 */
func SparkLengthCal(Length, Power string) string {
	//当电弧长度不为空，而功率为空的时候
	if (strings.Compare(Length, "") != 0) && (strings.Compare(Power, "") == 0) {
		Length_x, _ := strconv.ParseFloat(Length, 32)
		p := math.Pow((Length_x / (1.7 * 2.5)), 2)
		return fmt.Sprintf("%0.6f", p)
		//当电弧长度为空，而功率不为空的时候
	} else if (strings.Compare(Length, "") == 0) && (strings.Compare(Power, "") != 0) {
		Power_x, _ := strconv.ParseFloat(Power, 32)
		l := 1.7 * 2.54 * math.Sqrt(Power_x)
		return fmt.Sprintf("%0.6f", l)
	}
	return "0.0"
}

func SparkLengthForm() {

	label := ui.NewLabel()
	label.SetText("电弧长度(cm)") //Length
	//输入框
	inputBox := ui.NewLineEdit()
	//------------------------------
	label2 := ui.NewLabel()
	label2.SetText("变压器功率(瓦特)") // LReverse
	//输入框2
	inputBox2 := ui.NewLineEdit()

	//计算结果
	CalBtn := ui.NewPushButton()
	CalBtn.SetText("计算")

	CalBtn.OnClicked(func() {
		//当两个都不为空的时候，弹出显示框
		if (strings.Compare(inputBox.Text(), "") != 0) && (strings.Compare(inputBox2.Text(), "") != 0) {
			messagebox := ui.NewMessageBox()
			messagebox.SetText("电弧长度与功率只能填写一个")
			messagebox.Show()
			inputBox.Clear()
			inputBox2.Clear()
		} else if (strings.Compare(inputBox.Text(), "") != 0) && (strings.Compare(inputBox2.Text(), "") == 0) {
			text := SparkLengthCal(inputBox.Text(), inputBox2.Text())
			inputBox2.SetText(text)
		} else if (strings.Compare(inputBox.Text(), "") == 0) && (strings.Compare(inputBox2.Text(), "") != 0) {
			text := SparkLengthCal(inputBox.Text(), inputBox2.Text())
			inputBox.SetText(text)
		}

	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(label)
	hbox.AddWidget(inputBox)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(label2)
	hbox2.AddWidget(inputBox2)

	hbox3 := ui.NewHBoxLayout()
	hbox3.AddWidget(CalBtn)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddLayout(hbox2)
	vbox.AddLayout(hbox3)

	widget := ui.NewWidget()
	widget.SetLayout(vbox)
	widget.SetWindowTitle("电弧长度计算器")
	widget.Show()
}
