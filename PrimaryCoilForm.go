package main

import (
	//"errors"
	"fmt"
	"github.com/sndnvaps/tesla_calculator/setting"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

//Unit = 25.4, mm -> inch , 25.4mm / Unit = 1 inch
var Unit float64 = 25.4 //

//PrimaryCoilForm
type PrimaryCoilForm struct {
	*ui.QWidget
	btn1 *ui.QPushButton // 计算按钮
	le1  *ui.QLineEdit   //输入框 ->线圈直径D单位mm
	le2  *ui.QLineEdit   //输入框 -> 线圈匝数 N
	le3  *ui.QLineEdit   //输入框 -> 绕线线径 W, 单位mm
	le4  *ui.QLineEdit   //输入框 -> 匝间距 S, 单位mm

	le5 *ui.QLineEdit //输出框
	le6 *ui.QLineEdit //输出框
	le7 *ui.QLineEdit //输出框
	le8 *ui.QLineEdit //输出框

	label1 *ui.QLabel
	label2 *ui.QLabel
	label3 *ui.QLabel
	label4 *ui.QLabel
	label5 *ui.QLabel
	label6 *ui.QLabel
	label7 *ui.QLabel
	label8 *ui.QLabel

	picbox *ui.QLabel //用于放图片
}

//func FormHeight(N, S, W string) string
func FormHeight(N, S, W string) string {
	n, _ := strconv.ParseFloat(N, 64)
	s, _ := strconv.ParseFloat(S, 64)
	w, _ := strconv.ParseFloat(W, 64)
	h := n * (w*1 + s*1)
	return fmt.Sprintf("%0.6f", h) //返回的单位为mm
}

//func WireLong(N, D string) string
func WireLong(N, D string) string {
	n, _ := strconv.ParseFloat(N, 64)
	d, _ := strconv.ParseFloat(D, 64)
	w1 := n * math.Pi * d / 1000.0
	return fmt.Sprintf("%0.6f", w1)
}

//func CalcCap(H, D string) string
func CalcCap(H, D string) string {
	h, _ := strconv.ParseFloat(H, 64) //单位mm
	d, _ := strconv.ParseFloat(D, 64) //单位mm

	h = h / Unit //转换成 inch
	d = d / Unit //转换成 inch
	r := d / 2.0 //求半径

	cs := 5.08 * r * (0.0563*(h/r) + 0.08 + 0.38*math.Sqrt(1/(h/r)))

	return fmt.Sprintf("%0.6f", cs)
}

//func CalcInductance(N, D, H string) string
func CalcInductance(N, D, H string) string {
	n, _ := strconv.ParseFloat(N, 64)
	h, _ := strconv.ParseFloat(H, 64) //单位mm
	d, _ := strconv.ParseFloat(D, 64) //单位mm

	h = h / Unit //转换成 inch
	d = d / Unit //转换成 inch
	r := d / 2.0 //求半径

	i := (n * n * r * r) / (9*r + 10*h)

	return fmt.Sprintf("%0.6f", i)
}

//func CalPrimaryCoilInfo(D, N, W, S string) [4]string
func CalPrimaryCoilInfo(D, N, W, S string) [4]string {
	var output [4]string //用于存放输出结果
	H := FormHeight(N, S, W)
	output[0] = H
	output[1] = WireLong(N, D)
	output[2] = CalcInductance(N, D, H)
	output[3] = CalcCap(H, D)

	return output
}

//func NewPrimaryCoilForm() (*PrimaryCoilForm, error)
func NewPrimaryCoilForm() (*PrimaryCoilForm, error) {

	w := &PrimaryCoilForm{}
	w.QWidget = ui.NewWidget()

	w.btn1 = ui.NewPushButton()
	w.btn1.SetText(Lang.Tr("primary.PriCalBtn"))

	w.label1 = ui.NewLabel()
	w.label1.SetText(Lang.Tr("primary.PriFormDia"))
	w.le1 = ui.NewLineEdit()

	w.label2 = ui.NewLabel()
	w.label2.SetText(Lang.Tr("primary.PriTurns"))
	w.le2 = ui.NewLineEdit()

	w.label3 = ui.NewLabel()
	w.label3.SetText(Lang.Tr("primary.PriDia"))
	w.le3 = ui.NewLineEdit()

	w.label4 = ui.NewLabel()
	w.label4.SetText(Lang.Tr("primary.PriSpace"))
	w.le4 = ui.NewLineEdit()

	w.label5 = ui.NewLabel()
	w.label5.SetText(Lang.Tr("primary.PriFormHeigh"))
	w.le5 = ui.NewLineEdit()

	w.label6 = ui.NewLabel()
	w.label6.SetText(Lang.Tr("primary.PriLength"))
	w.le6 = ui.NewLineEdit()

	w.label7 = ui.NewLabel()
	w.label7.SetText(Lang.Tr("primary.PriInductance"))
	w.le7 = ui.NewLineEdit()

	w.label8 = ui.NewLabel()
	w.label8.SetText(Lang.Tr("primary.PriParasiticCap"))
	w.le8 = ui.NewLineEdit()

	w.picbox = ui.NewLabel()

	//设置为只读
	w.le5.SetReadOnly(true)
	w.le6.SetReadOnly(true)
	w.le7.SetReadOnly(true)
	w.le8.SetReadOnly(true)

	ImageBox := ui.NewPixmap()
	imgData, _ := setting.Asset("images/helix_fig.png") //先加载图片
	ImageBox.LoadFromData(imgData)

	w.picbox.SetPixmap(ImageBox)

	w.btn1.OnClicked(func() {
		if (strings.Compare(w.le1.Text(), "") != 0) && (strings.Compare(w.le2.Text(), "") != 0) &&
			(strings.Compare(w.le3.Text(), "") != 0) && (strings.Compare(w.le4.Text(), "") != 0) {
			output := CalPrimaryCoilInfo(w.le1.Text(), w.le2.Text(), w.le3.Text(), w.le4.Text())
			w.le5.SetText(output[0])
			w.le6.SetText(output[1])
			w.le7.SetText(output[2])
			w.le8.SetText(output[3])
		} else {
			messagebox := ui.NewMessageBox()
			messagebox.SetText(Lang.Tr("primary.PriAboutbox"))
			messagebox.Show()
			w.le1.Clear()
			w.le2.Clear()
			w.le3.Clear()
			w.le4.Clear()
			w.le4.Clear()
			w.le5.Clear()
			w.le6.Clear()
			w.le7.Clear()
			w.le8.Clear()
		}
	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(w.label1)
	hbox.AddWidget(w.le1)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(w.label2)
	hbox2.AddWidget(w.le2)

	hbox3 := ui.NewHBoxLayout()
	hbox3.AddWidget(w.label3)
	hbox3.AddWidget(w.le3)

	hbox4 := ui.NewHBoxLayout()
	hbox4.AddWidget(w.label4)
	hbox4.AddWidget(w.le4)

	hbox5 := ui.NewHBoxLayout()
	hbox5.AddWidget(w.btn1)

	hbox6 := ui.NewHBoxLayout()
	hbox6.AddWidget(w.label5)
	hbox6.AddWidget(w.le5)

	hbox7 := ui.NewHBoxLayout()
	hbox7.AddWidget(w.label6)
	hbox7.AddWidget(w.le6)

	hbox8 := ui.NewHBoxLayout()
	hbox8.AddWidget(w.label7)
	hbox8.AddWidget(w.le7)

	hbox9 := ui.NewHBoxLayout()
	hbox9.AddWidget(w.label8)
	hbox9.AddWidget(w.le8)

	hbox10 := ui.NewHBoxLayout()
	hbox10.AddWidget(w.picbox)

	vboxL := ui.NewVBoxLayout()
	vboxL.AddLayout(hbox)
	vboxL.AddLayout(hbox2)
	vboxL.AddLayout(hbox3)
	vboxL.AddLayout(hbox4)
	vboxL.AddLayout(hbox5)
	vboxL.AddLayout(hbox6)
	vboxL.AddLayout(hbox7)
	vboxL.AddLayout(hbox8)
	vboxL.AddLayout(hbox9)

	hboxMain := ui.NewHBoxLayout()
	hboxMain.AddLayout(vboxL)
	hboxMain.AddLayout(hbox10)

	w.SetLayout(hboxMain)

	w.SetWindowTitle(Lang.Tr("primary.PriWinTitle"))
	return w, nil

}
