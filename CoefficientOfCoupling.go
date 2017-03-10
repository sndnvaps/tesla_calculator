package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

/*  用于计算特斯拉线圈的互感系统及耦合系数
 *  //MutualInductance = (LForward - LReverse)/4
 *  //CouplingDegree =  MutualInductance / sqrt (L1*L2)
 *  //源代码参考自 http://bbs.kechuang.org/read-kc-tid-58673-1-1.html
 */
func CoefficientOfCouplingCal(LForward, LReverse, L1, L2 string) string {
	if (strings.Compare(LForward, "") != 0) && (strings.Compare(LReverse, "") != 0) && (strings.Compare(L1, "") != 0) && (strings.Compare(L2, "") != 0) {
		LForward_x, _ := strconv.ParseFloat(LForward, 32)
		LReverse_y, _ := strconv.ParseFloat(LReverse, 32)
		L1_x, _ := strconv.ParseFloat(L1, 32)
		L2_y, _ := strconv.ParseFloat(L2, 32)
		res := (math.Abs(LForward_x-LReverse_y) / 4) / (math.Sqrt(L1_x * L2_y))
		return fmt.Sprintf("%0.6f", res)
	}

	return "0.0"
}

func GetMutualInductance(LForward, LReverse string) string {
	if (strings.Compare(LForward, "") != 0) && (strings.Compare(LReverse, "") != 0) {
		LForward_x, _ := strconv.ParseFloat(LForward, 32)
		LReverse_y, _ := strconv.ParseFloat(LReverse, 32)
		mi := (math.Abs(LForward_x-LReverse_y) / 4)
		return fmt.Sprintf("%0.6f", mi)
	}
	return "0.0"
}

func CoefficientOfCouplinForm() {

	label := ui.NewLabel()
	label.SetText("正向测试电感值") //LForward
	//输入框
	inputBox := ui.NewLineEdit()
	//------------------------------
	label2 := ui.NewLabel()
	label2.SetText("反向测试电感值") // LReverse
	//输入框2
	inputBox2 := ui.NewLineEdit()

	label3 := ui.NewLabel()
	label3.SetText("线圈电感L1")
	//输入框3
	inputBox3 := ui.NewLineEdit()

	label4 := ui.NewLabel()
	label4.SetText("线圈电感L2")
	//输入框4
	inputBox4 := ui.NewLineEdit()

	//计算结果
	CalBtn := ui.NewPushButton()
	CalBtn.SetText("计算")

	//
	label5 := ui.NewLabel()
	label5.SetText("互感系数")

	label6 := ui.NewLabel()
	label6.SetText("耦合系数")

	//用于显示输出的结果 显示耦合系数
	CouplingDegree := ui.NewLineEdit()
	CouplingDegree.SetReadOnly(true) //设置为只读模式

	//用于显示互感系数
	MutualInductance := ui.NewLineEdit()
	MutualInductance.SetReadOnly(true)

	//点击，计算按钮的时候，把 inputbox + input2的内容， 连接起来，并显示在 CouplingDegree里面
	CalBtn.OnClicked(func() {
		CouplingDegree.Clear()
		MutualInductance.Clear()
		text := CoefficientOfCouplingCal(inputBox.Text(), inputBox2.Text(), inputBox3.Text(), inputBox4.Text())
		text2 := GetMutualInductance(inputBox.Text(), inputBox2.Text())
		CouplingDegree.SetText(text)
		MutualInductance.SetText(text2)
	})

	//------------------开始画图部分---------------------
	//准备显示图片
	picboxLabel := ui.NewLabel()

	ImageBox := ui.NewPixmap()
	ImageBox.Load(":/images/CouplingDegree.jpg") //先加载图片

	picboxLabel.SetPixmap(ImageBox)
	//---------------结束画图部分----------------------

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(label)
	hbox.AddWidget(inputBox)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(label2)
	hbox2.AddWidget(inputBox2)

	hbox3 := ui.NewHBoxLayout()
	hbox3.AddWidget(label3)
	hbox3.AddWidget(inputBox3)

	hbox4 := ui.NewHBoxLayout()
	hbox4.AddWidget(label4)
	hbox4.AddWidget(inputBox4)

	hbox5 := ui.NewHBoxLayout()
	hbox5.AddWidget(label5)
	hbox5.AddWidget(MutualInductance)

	hbox6 := ui.NewHBoxLayout()
	hbox6.AddWidget(label6)
	hbox6.AddWidget(CouplingDegree)

	hbox7 := ui.NewHBoxLayout()
	hbox7.AddWidget(CalBtn)

	hbox8 := ui.NewHBoxLayout()
	hbox8.AddWidget(picboxLabel) //显示图片

	vbox := ui.NewVBoxLayout()
	vbox.AddStretchWithStretch(1)
	vbox.AddLayout(hbox)
	vbox.AddLayout(hbox2)
	vbox.AddLayout(hbox3)
	vbox.AddLayout(hbox4)
	vbox.AddLayout(hbox5)
	vbox.AddLayout(hbox6)
	vbox.AddLayout(hbox7)
	vbox.AddLayout(hbox8)

	widget := ui.NewWidget()
	widget.SetLayout(vbox)
	widget.Show()
}
