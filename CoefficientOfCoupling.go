package main

import (
	"errors"
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

	file := ui.NewFileWithName(":/forms/CoefficientForm.ui")
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

	w.label_1 = ui.NewLabelFromDriver(formWidget.FindChild("label_1"))
	w.label_2 = ui.NewLabelFromDriver(formWidget.FindChild("label_2"))
	w.label_3 = ui.NewLabelFromDriver(formWidget.FindChild("label_3"))
	w.label_4 = ui.NewLabelFromDriver(formWidget.FindChild("label_4"))
	w.label_5 = ui.NewLabelFromDriver(formWidget.FindChild("label_5"))
	w.label_6 = ui.NewLabelFromDriver(formWidget.FindChild("label_6"))

	w.picbox = ui.NewLabelFromDriver(formWidget.FindChild("pic_label_1"))

	//设置为只读
	w.le5.SetReadOnly(true)
	w.le6.SetReadOnly(true)

	ImageBox := ui.NewPixmap()
	ImageBox.Load(":/images/CouplingDegree.png") //先加载图片 CouplingDegree

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

	layout := ui.NewVBoxLayout()
	layout.AddWidget(formWidget)
	w.SetLayout(layout)

	w.SetWindowTitle("耦合系数计算器")
	return w, nil
}
