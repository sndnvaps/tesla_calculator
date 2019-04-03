package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

//此文件用于计算特斯拉线圈的 次级线圈的相关参数
/*
 * 次级线圈的箍数
 * 次级线圈的电感(μH)
 * 次级线圈的寄生电容(pf)
 * 在不加顶端（均压环)的状态下，谐振频率
 */

/*
 * 计算次级线圈的寄生电容
 * 计算次级线圈的箍数
 * Secondary Capacitance = (0.29* Secondary Wire Winding Height) + (0.41 * (Secondary Form Diameter/2)) + (1.94 * sqrt(((Secondary Form Diameter/2)^3)/Secondary Wire Winding Height))
 *  绕线高度 Height = Coil Turns * Magnet Wire Diameter
 *  Secondary coil Turns = (1/(Magnet Wire Diameter + 0.000001)) * Secondary Wire Winding Height * 0.97
 *  次级线圈电感 μH
 *  Secondary Inductance = ((((Secondary Coil Turns ^2) * ((Secondary Form Diameter/2)^2))/((9 * (Secondary Form Diameter / 2)) + (10 * Secondary Wire Winding Height))) * 0.001) * Secondary Inductance Adjust
 *                       = (((System.Math.Pow(Secondary Coil Turns,2.0) * (System.Math.Pow(Secondary Form Diameter/2,2.0))/(( 9 * (Secondary Form Diameter/2))) + (10 * Secondary Wire Winding Height))) * 0.001) * Secondary Inductance Adjust
 *
 *  0.27mm = 0.01063 inch
 *  0.25mm = 0.00984 inch
 *  0.32mm = 0.0125984 inch
 *  1mm = 0.0393701 inch
 *  1cm = 1/2.54 = 0.3937008 inch
 *
 *  绕线线管的直径 FormDiameter
 *
 */

/*
   public double  Output;
   private double MagnetWireDiameter;// = 0.01063; // 次级线径 0.27mm
   private double FormDiameter; //管径 输入为厘米
   private double FormHeight; //绕线的高度
   private double OutputSecCap; //用于输出次级寄生电容
   private double OutputSecInduct; //用于输出次级电感 μH
   private double OutputResonant; //用于输出谐振频率 Hz

   //1H=1000000μH,1μH=0.0000001H

   //1F=10^6uF=10^12pF
   //1F=10^12pF
*/

/* @parame
 * IFormHight -> 绕线的高度， 单位cm
 * IFormDiameter -> 管径, 单位 mm
 * IWireDiameter -> 次级线圈的线径, 单位 mm
 * @return
 * [4]string = [4]string{s1, s2, s3, s4}
 * s1 = 次级箍数
 * s2 = 次级线圈电感 (μH)
 * s3 = 次级线圈寄生电容 (pf)
 * s4 = 在不加均压环的状态下的 谐振频率 (hz)
 */
func SecCoilInfoCal(IFormHight, IFormDiameter, IWireDiameter string) [4]string {
	var output [4]string = [4]string{"0.0", "0.0", "0.0", "0.0"} //默认初始化值为全部 0.0
	//当三个参数，都不为空的时候
	if (strings.Compare(IFormHight, "") != 0) && (strings.Compare(IFormDiameter, "") != 0) && (strings.Compare(IWireDiameter, "") != 0) {
		IFormHight_x, _ := strconv.ParseFloat(IFormHight, 32)
		IFormDiameter_y, _ := strconv.ParseFloat(IFormDiameter, 32)
		IWireDiameter_z, _ := strconv.ParseFloat(IWireDiameter, 32)

		turns := SecCoilTurns(IFormHight_x, IWireDiameter_z)      //箍数
		Induct := SecInduct(turns, IFormHight_x, IFormDiameter_y) //电感
		Cap := SecCap(IFormHight_x, IFormDiameter_y)              //电容
		Resonant := SecResonant(Cap, Induct)
		output[0] = turns
		output[1] = Induct
		output[2] = Cap
		output[3] = Resonant

		return output
	}
	//返回 slice string
	return [4]string{"0.0", "0.0", "0.0", "0.0"}
}

/*
 *@parame
 * FormHight -> 线圈的高度, 单位 cm
 * WireDiameter -> 次级线圈的线径, 单位 mm
 *@return
 * string -> 次级线圈的箍数 （FormHight * 10 / (WireDiameter + 0.01)）* 0.97  //此处的0.97为修正系数，因为考虑的手工绕线的情况
 */

func SecCoilTurns(FormHight, WireDiameter float64) string {
	MagnetWireDiameter := WireDiameter //毫米
	IFormHeight := (FormHight * 10.0)  //将厘米转换成 毫米
	output := (1 / (MagnetWireDiameter + 0.01)) * IFormHeight * 0.97
	return fmt.Sprintf("%0.6f", output)
}

/*
 *@parame
 * FormHight -> 线圈的高度, 单位 cm
 * FormDiameter -> 次级线圈的直径, 单位 mm
 *@return
 * string -> 次级线圈的寄生电容
 */
// 毫米转 英寸 需要除 25.4
func SecCap(FormHight, FormDiameter float64) string {
	var u float64 = 25.4                       //
	IFormHeight := (float64)(FormHight * 10.0) //将厘米转换成 毫米
	IFormDiameter := (float64)(FormDiameter)   // 毫米
	Radius := IFormDiameter / 2                //线管的半径，单位毫米
	cs := 5.08 * Radius / u * (0.0563*((IFormHeight/u)/(Radius/u)) + 0.08 + 0.38*math.Sqrt(1/((IFormHeight/u)/(Radius/u))))
	return fmt.Sprintf("%0.6f", cs)
}

/*
 *@parame
 * Turns -> 次级线圈的箍数，单位箍
 * FormHight -> 线圈的高度, 单位 cm
 * FormDiameter -> 次级线圈的直径, 单位 mm
 *@return
 * string -> 次级线圈的电感(μH)
 */
func SecInduct(Turns string, FormHight, FormDiameter float64) string {
	ITurns, _ := strconv.ParseFloat(Turns, 32)                                                                                  //
	IFormHeight := (float64)(FormHight * (1 / 2.54))                                                                            //将厘米转换成 英寸
	IFormDiameter := (float64)(FormDiameter * (1 / 25.4))                                                                       //将毫米转换成 英寸
	output := (((math.Pow(ITurns, 2.0)) * (math.Pow(IFormDiameter/2.0, 2))) / ((9 * (IFormDiameter / 2)) + (10 * IFormHeight))) // * 0.001 * 13.72; //miss Secondary Induct Adjust
	return fmt.Sprintf("%0.6f", output)
}

/*
 *@parame
 * SecCap -> 次级线圈的寄生电容，单位 pf
 * SecInduct -> 次级线圈的电感， 单位 μH
 *@return
 * string -> 次级线圈在不加均压环的状态下的 谐振频率 hz
 */
//1/(2*pi*sqrt(L*C))
func SecResonant(SecCap, SecInduct string) string {
	seccap_x, _ := strconv.ParseFloat(SecCap, 64)
	secinduct_y, _ := strconv.ParseFloat(SecInduct, 64)
	output := (float64)(1 / (2 * math.Pi * math.Sqrt((seccap_x*math.Pow(10, -12.0))*(secinduct_y*math.Pow(10, -6.0))))) //谐振频率 (hz)
	return fmt.Sprintf("%0.6f", output)
}
func SecCoilInfoForm() {

	label := ui.NewLabel()
	label.SetText("次级线圈的绕线高度(cm)") //Length
	//输入框
	inputBox := ui.NewLineEdit()
	//------------------------------
	label2 := ui.NewLabel()
	label2.SetText("次级线圈线径(mm)") // LReverse
	//输入框2
	inputBox2 := ui.NewLineEdit()

	label3 := ui.NewLabel()
	label3.SetText("次级线圈直径(mm)")
	//输入框3
	inputBox3 := ui.NewLineEdit()

	label4 := ui.NewLabel()
	label4.SetText("次级箍数")
	inputbox4 := ui.NewLineEdit()
	inputbox4.SetReadOnly(true)

	label5 := ui.NewLabel()
	label5.SetText("次级线圈电感(μH)")
	inputbox5 := ui.NewLineEdit()
	inputbox5.SetReadOnly(true)

	label6 := ui.NewLabel()
	label6.SetText("次级线圈寄生电容(pf)")
	inputbox6 := ui.NewLineEdit()
	inputbox6.SetReadOnly(true)

	label7 := ui.NewLabel()
	label7.SetText("不加均压环的谐振频率(Hz)")
	inputbox7 := ui.NewLineEdit()
	inputbox7.SetReadOnly(true)

	//计算结果
	CalBtn := ui.NewPushButton()
	CalBtn.SetText("计算")

	CalBtn.OnClicked(func() {
		//当两个都不为空的时候，弹出显示框
		if (strings.Compare(inputBox.Text(), "") != 0) && (strings.Compare(inputBox2.Text(), "") != 0) && (strings.Compare(inputBox3.Text(), "") != 0) {
			output := SecCoilInfoCal(inputBox.Text(), inputBox3.Text(), inputBox2.Text())
			inputbox4.SetText(output[0])
			inputbox5.SetText(output[1])
			inputbox6.SetText(output[2])
			inputbox7.SetText(output[3])

		} else {
			messagebox := ui.NewMessageBox()
			messagebox.SetText("必须要同时输入三个参数！")
			messagebox.Show()
			inputBox.Clear()
			inputBox2.Clear()
			inputBox3.Clear()
			inputbox4.Clear()
			inputbox5.Clear()
			inputbox6.Clear()
			inputbox7.Clear()
		}

	})

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
	hbox4.AddWidget(CalBtn)

	hbox5 := ui.NewHBoxLayout()
	hbox5.AddWidget(label4)
	hbox5.AddWidget(inputbox4)

	hbox6 := ui.NewHBoxLayout()
	hbox6.AddWidget(label5)
	hbox6.AddWidget(inputbox5)

	hbox7 := ui.NewHBoxLayout()
	hbox7.AddWidget(label6)
	hbox7.AddWidget(inputbox6)

	hbox8 := ui.NewHBoxLayout()
	hbox8.AddWidget(label7)
	hbox8.AddWidget(inputbox7)

	vbox := ui.NewVBoxLayout()
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
	widget.SetWindowTitle("次级线圈参数计算")
	widget.Show()
}
