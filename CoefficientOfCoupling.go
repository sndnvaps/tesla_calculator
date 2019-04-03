package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

type CoefficientForm struct {
	*ui.QWidget
	btn1 *ui.QPushButton // 计算按钮
	le1  *ui.QLineEdit   //输入框 ->正向测量电感
	le2  *ui.QLineEdit   //输入框 -> 反向测量电感
	le3  *ui.QLineEdit   //输入框 -> 线圈电感L1
	le4  *ui.QLineEdit   //输入框 -> 线圈电感L2

	le5 *ui.QLineEdit //输出框
	le6 *ui.QLineEdit //输出框

	label_1 *ui.QLabel //
	label_2 *ui.QLabel
	label_3 *ui.QLabel
	label_4 *ui.QLabel //
	label_5 *ui.QLabel
	label_6 *ui.QLabel
	
	picbox *ui.QLabel //用于放图片
}



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


func NewCoefficientForm() (*CoefficientForm, error) {

	w := &CoefficientForm{}
	w.QWidget = ui.NewWidget()
    //w.SetFixedWidth(700)
    //w.SetFixedHeight(280)

    w.btn1 = ui.NewPushButton()
    w.btn1.SetText("计算")

    w.label_1 = ui.NewLabel()
    w.label_1.SetText("正向测试电感值")
    w.le1 = ui.NewLineEdit()
    //QWidget::setGeometry(int,int,int,int)
    //w.le1.SetGeometryWithXYWidthHeight(150,10,113,20)

    w.label_2 = ui.NewLabel()
    w.label_2.SetText("反向测试电感值")
    w.le2 = ui.NewLineEdit()
    //w.le2.SetGeometryWithXYWidthHeight(150,40,113,20)

    w.label_3 = ui.NewLabel()
    w.label_3.SetText("线圈电感L1")
    //w.label_3.SetGeometryWithXYWidthHeight(20,70,111,20)
    w.le3 = ui.NewLineEdit()
    //w.le3.SetGeometryWithXYWidthHeight(150,70,113,20)

    w.label_4 = ui.NewLabel()
    w.label_4.SetText("线圈电感L2")
    w.le4 = ui.NewLineEdit()
    //w.le4.SetGeometryWithXYWidthHeight(150,100,113,20)

    w.label_5 = ui.NewLabel()
    w.label_5.SetText("互感系数")
    w.le5 = ui.NewLineEdit()
    //w.le5.SetGeometryWithXYWidthHeight(150,170,113,20)

    w.label_6 = ui.NewLabel()
    w.label_6.SetText("耦合系数")
    w.le6 = ui.NewLineEdit()
    //w.le6.SetGeometryWithXYWidthHeight(150,200,113,20)

    w.picbox = ui.NewLabel()


	//设置为只读
	w.le5.SetReadOnly(true)
	w.le6.SetReadOnly(true)

	ImageBox := ui.NewPixmap()
    imgData, _ := Asset("images/CouplingDegree.png")
    ImageBox.LoadFromData(imgData)

	w.picbox.SetPixmap(ImageBox)

	w.btn1.OnClicked(func() {
		if (strings.Compare(w.le1.Text(), "") != 0) && (strings.Compare(w.le2.Text(), "") != 0) &&
			(strings.Compare(w.le3.Text(), "") != 0) && (strings.Compare(w.le4.Text(), "") != 0) {
			coc := CoefficientOfCouplingCal(w.le1.Text(), w.le2.Text(), w.le3.Text(), w.le4.Text())
			Inductance  := GetMutualInductance(w.le1.Text(), w.le2.Text())
			w.le5.SetText(coc)
			w.le6.SetText(Inductance)
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
    hbox4.AddWidget(w.label_4)
    hbox4.AddWidget(w.le4)

    hbox5 := ui.NewHBoxLayout()
    hbox5.AddWidget(w.btn1)

    hbox6 := ui.NewHBoxLayout()
    hbox6.AddWidget(w.label_5)
    hbox6.AddWidget(w.le5)

    hbox7 := ui.NewHBoxLayout()
    hbox7.AddWidget(w.label_6)
    hbox7.AddWidget(w.le6)

    hbox8 := ui.NewHBoxLayout()
    hbox8.AddWidget(w.picbox)

    vboxL := ui.NewVBoxLayout()
	vboxL.AddLayout(hbox)
	vboxL.AddLayout(hbox2)
	vboxL.AddLayout(hbox3)
	vboxL.AddLayout(hbox4)
	vboxL.AddLayout(hbox5)
	vboxL.AddLayout(hbox6)
	vboxL.AddLayout(hbox7)

    hboxMain := ui.NewHBoxLayout()
    hboxMain.AddLayout(vboxL)
    hboxMain.AddLayout(hbox8)

	w.SetLayout(hboxMain)

	w.SetWindowTitle("耦合系数计算器")
	return w, nil
}
