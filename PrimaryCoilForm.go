package main

import (
	"errors"
	"fmt"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

/*
<!--
var d=6, r=0, n=10, w = 0.30, s = 0.30, h=6, i=0, u=1, cs=0, wl=0, pi=3.14159;

function gv() {
d=document.helixcalc.diameter.value;
n=document.helixcalc.turns.value;
w=document.helixcalc.wirediameter.value;
s=document.helixcalc.turnspacing.value;
h=n*(w*1+s*1);
h=h*1e+3;
h=parseInt(h);
h=h/1e+3;
r=d/2;
}

function pv() {
document.helixcalc.diameter.value=d;
document.helixcalc.turns.value=n;
document.helixcalc.wirediameter.value=w;
document.helixcalc.turnspacing.value=s;
document.helixcalc.height.value=h;
r=d/2;
calcinductance();
calccap();
}

function calcinductance() {
gv();
i=(n*n*r/u*r/u)/(r/u*9+h/u*10);
i=i*1e+3;
i=parseInt(i);
i=i/1e+3;
wl=n*pi*d;
if(u==25.4){wl=wl/1000};
wl=wl*1e+3;
wl=parseInt(wl);
wl=wl/1e+3;
document.helixcalc.inductance.value=i;
document.helixcalc.height.value=h;
document.helixcalc.wirelength.value=wl;
}

function calccap() {
cs=5.08*r/u*(0.0563*((h/u)/(r/u))+0.08+0.38*Math.sqrt(1/((h/u)/(r/u))));
cs=cs*1e+3;
cs=parseInt(cs);
cs=cs/1e+3;
document.helixcalc.selfcap.value=cs
}

function changes() {
gv();
calcinductance();
pv();
}

function selmm() {
if (u==1)
	{
		u=25.4;
		gv();
		d=d*u;
d=d*1e+3;
d=parseInt(d+0.5);
d=d/1e+3;
		w=w*u;
w=w*1e+3;
w=parseInt(w+0.5);
w=w/1e+3;
		s=s*u;
s=s*1e+3;
s=parseInt(s+0.5);
s=s/1e+3;
		pv();
	}

document.helixcalc.unit1.value="mm";
document.helixcalc.unit2.value="mm";
document.helixcalc.unit3.value="mm";
document.helixcalc.unit4.value="mm";
document.helixcalc.unit5.value="metre";
}

function selinches() {
if (u==25.4)
	{
		gv();
		d=d/u;
d=d*1e+3;
d=parseInt(d+0.5);
d=d/1e+3;
		w=w/u;
w=w*1e+3;
w=parseInt(w+0.5);
w=w/1e+3;
		s=s/u;
s=s*1e+3;
s=parseInt(s+0.5);
s=s/1e+3;
		u=1;
		pv();
	}
document.helixcalc.unit1.value="inches";
document.helixcalc.unit2.value="inches";
document.helixcalc.unit3.value="inches";
document.helixcalc.unit4.value="inches";
document.helixcalc.unit5.value="inches";
}
// -->

*/

var Unit float64 = 25.4 //

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

	label_1 *ui.QLabel //
	label_2 *ui.QLabel
	label_3 *ui.QLabel
	label_4 *ui.QLabel //
	label_5 *ui.QLabel
	label_6 *ui.QLabel
	label_7 *ui.QLabel //
	label_8 *ui.QLabel

	picbox *ui.QLabel //用于放图片
}

//返回螺线管高度
func FormHeight(N, S, W string) string {
	n, _ := strconv.ParseFloat(N, 64)
	s, _ := strconv.ParseFloat(S, 64)
	w, _ := strconv.ParseFloat(W, 64)
	h := n * (w*1 + s*1)
	return fmt.Sprintf("%0.6f", h) //返回的单位为mm
}

//参数列表：
/*N -> 线圈匝数
 *D -> 螺丝管直径， 单位mm
 *返回值：
 * 返回的结果为，使用的线长，单位m
 *
 */
func WireLong(N, D string) string {
	n, _ := strconv.ParseFloat(N, 64)
	d, _ := strconv.ParseFloat(D, 64)
	w1 := n * math.Pi * d / 1000.0
	return fmt.Sprintf("%0.6f", w1)
}

//
func CalcCap(H, D string) string {
	h, _ := strconv.ParseFloat(H, 64) //单位mm
	d, _ := strconv.ParseFloat(D, 64) //单位mm

	h = h / Unit //转换成 inch
	d = d / Unit //转换成 inch
	r := d / 2.0 //求半径

	cs := 5.08 * r * (0.0563*(h/r) + 0.08 + 0.38*math.Sqrt(1/(h/r)))

	return fmt.Sprintf("%0.6f", cs)
}

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

/*
 *
 *
 *[4]string { h, long, i, cap } -> {螺线管高度mm, 漆包线长度m, 电感量(L), 寄生电容(pf)  }
 *
 *
 */

func CalPrimaryCoilInfo(D, N, W, S string) [4]string {
	var output [4]string //用于存放输出结果
	H := FormHeight(N, S, W)
	output[0] = H
	output[1] = WireLong(N, D)
	output[2] = CalcInductance(N, D, H)
	output[3] = CalcCap(H, D)

	return output
}
func NewPrimaryCoilForm() (*PrimaryCoilForm, error) {
	w := &PrimaryCoilForm{}
	w.QWidget = ui.NewWidget()

	file := ui.NewFileWithName(":/forms/PrimaryCoilForm.ui")
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
	w.le3 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_3"))
	w.le4 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_4"))

	w.le5 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_5"))
	w.le6 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_6"))
	w.le7 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_7"))
	w.le8 = ui.NewLineEditFromDriver(formWidget.FindChild("lineEdit_8"))

	w.label_1 = ui.NewLabelFromDriver(formWidget.FindChild("label_1"))
	w.label_2 = ui.NewLabelFromDriver(formWidget.FindChild("label_2"))
	w.label_3 = ui.NewLabelFromDriver(formWidget.FindChild("label_3"))
	w.label_4 = ui.NewLabelFromDriver(formWidget.FindChild("label_4"))
	w.label_5 = ui.NewLabelFromDriver(formWidget.FindChild("label_5"))
	w.label_6 = ui.NewLabelFromDriver(formWidget.FindChild("label_6"))
	w.label_7 = ui.NewLabelFromDriver(formWidget.FindChild("label_7"))
	w.label_8 = ui.NewLabelFromDriver(formWidget.FindChild("label_8"))

	w.picbox = ui.NewLabelFromDriver(formWidget.FindChild("pic_label_1"))

	//设置为只读
	w.le5.SetReadOnly(true)
	w.le6.SetReadOnly(true)
	w.le7.SetReadOnly(true)
	w.le8.SetReadOnly(true)

	ImageBox := ui.NewPixmap()
	ImageBox.Load(":/images/helix_fig.png") //先加载图片

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
			messagebox.SetText("必须要同时输入4个数值")
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

	layout := ui.NewVBoxLayout()
	layout.AddWidget(formWidget)
	w.SetLayout(layout)

	w.SetWindowTitle("初级线圈参数计算")
	return w, nil

}
