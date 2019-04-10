package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
	"math"
	"strconv"
	"strings"
)

//1H=1000000μH,1μH=0.0000001H

//1F=10^6uF=10^12pF
//1F=10^12pF
   
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
func SecCoilInfoCal(IFormHight, IFormDiameter, IWireDiameter string) [9]string {
	var output [9]string = [9]string{"0.0", "0.0", "0.0", "0.0", "0.0", "0.0", "0.0", "0.0", "0.0"} //默认初始化值为全部 0.0
	//当三个参数，都不为空的时候
	if (strings.Compare(IFormHight, "") != 0) && (strings.Compare(IFormDiameter, "") != 0) && (strings.Compare(IWireDiameter, "") != 0) {
		IFormHight_x, _ := strconv.ParseFloat(IFormHight, 32)
		IFormDiameter_y, _ := strconv.ParseFloat(IFormDiameter, 32)
		IWireDiameter_z, _ := strconv.ParseFloat(IWireDiameter, 32)

		turns := SecCoilTurns(IFormHight_x, IWireDiameter_z)      //箍数
		Induct := SecInduct(turns, IFormHight_x, IFormDiameter_y) //电感
		Cap := SecCap(IFormHight_x, IFormDiameter_y)              //电容
		Resonant := SecResonant(Cap, Induct)                      //Frequency
		skindepth := SkinDepth(Resonant)
		resistanceDC := ResitanceDC(IWireDiameter, IFormDiameter, turns)
		skinEffectFactor := SkinEffecFactor(IWireDiameter, skindepth)
		resistanceAC := ResitanceAC(skinEffectFactor, resistanceDC)

		QFactor := SecQFactor(Resonant, Induct, resistanceAC) //Sec Coil Quality Factor

		output[0] = turns
		output[1] = Induct
		output[2] = Cap
		output[3] = Resonant
		output[4] = skindepth
		output[5] = resistanceDC
		output[6] = skinEffectFactor
		output[7] = resistanceAC
		output[8] = QFactor

		return output
	}
	//返回 slice string
	return [9]string{"0.0", "0.0", "0.0", "0.0", "0.0", "0.0", "0.0", "0.0", "0.0"}
}

/*
 *@parame
 * FormHight -> 线圈的高度, 单位 cm
 * WireDiameter -> 次级线圈的线径, 单位 mm
 * Wire spacing = WireDiameter / 10
 *@return
 * string -> 次级线圈的箍数 （FormHight * 10 / (WireDiameter + WireDiameter/10)）//此处的0.97为修正系数，因为考虑的手工绕线的情况
 */

func SecCoilTurns(FormHight, WireDiameter float64) string {
	MagnetWireDiameter := WireDiameter //毫米
	IFormHeight := (FormHight * 10.0)  //将厘米转换成 毫米
	output := (1 / (MagnetWireDiameter + MagnetWireDiameter/10)) * IFormHeight
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
	cs := 5.08 * (Radius / u) * (0.0563*((IFormHeight/u)/(Radius/u)) + 0.08 + 0.38*math.Sqrt(1/((IFormHeight/u)/(Radius/u))))
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

/*
 *@parame
 * Fres -> Frequency of sec coil(hz)
 * return
 * string -> skin depth (m)
 * output = math.Sqrt((0.0179*math.Pow(10,-6)/(math.Pi*math.Pow(10,-7))))*(1/(math.Sqrt(Fres)))
 */
func SkinDepth(Fres string) string {

	fres, _ := strconv.ParseFloat(Fres, 64)
	output := math.Sqrt((0.0179*math.Pow(10, -6))/(4*math.Pi*math.Pi*math.Pow(10, -7))) * (1 / (math.Sqrt(fres)))
	return fmt.Sprintf("%0.8f", output)
}

/*
 *@parame
 * WireDia -> mm
 * CoilDia -> mm
 *
 */
func ResitanceDC(WireDia, CoilDia, Turns string) string {
	CoilDia_x, _ := strconv.ParseFloat(CoilDia, 64)
	Turns_x, _ := strconv.ParseFloat(Turns, 64)
	WireDia_x, _ := strconv.ParseFloat(WireDia, 64)
	output := (4 * (CoilDia_x / 1000) * Turns_x * (0.0179 * math.Pow(10, -6))) / ((WireDia_x / 1000) * (WireDia_x / 1000))
	return fmt.Sprintf("%0.8f", output)
}

/*
 *@parame
 * WireDia -> mm
 * SiknDepth -> m
 *return
 * output ->
 *   if(SkinDepth * 1000) > (WireDia/2) ->= 1
 *    else
       (((WireDia/1000)^2)/(4*(((WireDia/1000)*SkinDepth) - (SkinDepth^2))))
 *
*/
func SkinEffecFactor(WireDia, SkinDepth string) string {
	WireDia_x, _ := strconv.ParseFloat(WireDia, 64)
	SkinDepth_x, _ := strconv.ParseFloat(SkinDepth, 64)
	var output float64 = 0.0
	if (SkinDepth_x * 1000) > (WireDia_x / 2) {
		output = 1.0
	} else {
		output = math.Pow((WireDia_x/1000.0), 2) / (4 * (((WireDia_x / 1000.0) * SkinDepth_x) - (math.Pow(SkinDepth_x, 2))))
	}
	return fmt.Sprintf("%0.6f", output)
}

/*
 *@parame
 * skinEffectFactor ->
 * resitanceDC-> ohm
 *@return
 * output -> ohm
 *
 */
func ResitanceAC(skinEffectFactor, ResistanceDC string) string {

	skineffect, _ := strconv.ParseFloat(skinEffectFactor, 64)
	resistanceDC, _ := strconv.ParseFloat(ResistanceDC, 64)
	output := resistanceDC * skineffect * 3
	return fmt.Sprintf("%0.4f", output)
}

/*
 *@parame
 * Fres -> hz
 * Induct -> uH
 * ResitanceAC -> ohm
 *@return
 * output -> Quality Factor of Secondaly coil -> Q
 *   output = (2 * math.Pi * Fres *(Induct/1000000))/ResitanceAC
 */
func SecQFactor(Fres, Induct, ResitanceAC string) string {
	fres, _ := strconv.ParseFloat(Fres, 64)
	induct, _ := strconv.ParseFloat(Induct, 64)
	resitanceAC, _ := strconv.ParseFloat(ResitanceAC, 64)
	output := (2 * math.Pi * fres * (induct / 1000000)) / resitanceAC
	return fmt.Sprintf("%0.4f", output)
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
	outputbox1 := ui.NewLineEdit()
	outputbox1.SetReadOnly(true)

	label5 := ui.NewLabel()
	label5.SetText("次级线圈电感(μH)")
	outputbox2 := ui.NewLineEdit()
	outputbox2.SetReadOnly(true)

	label6 := ui.NewLabel()
	label6.SetText("次级线圈寄生电容(pf)")
	outputbox3 := ui.NewLineEdit()
	outputbox3.SetReadOnly(true)

	label7 := ui.NewLabel()
	label7.SetText("不加均压环的谐振频率(Hz)")
	outputbox4 := ui.NewLineEdit()
	outputbox4.SetReadOnly(true)

	label8 := ui.NewLabel()
	label8.SetText("趋肤深度(m)")
	outputbox5 := ui.NewLineEdit()
	outputbox5.SetReadOnly(true)

	label9 := ui.NewLabel()
	label9.SetText("次级电阻值(ohm)")
	outputbox6 := ui.NewLineEdit()
	outputbox6.SetReadOnly(true)

	label10 := ui.NewLabel()
	label10.SetText("趋肤效应因子")
	outputbox7 := ui.NewLineEdit()
	outputbox7.SetReadOnly(true)

	label11 := ui.NewLabel()
	label11.SetText("交变阻抗(ohm)")
	outputbox8 := ui.NewLineEdit()
	outputbox8.SetReadOnly(true)

	label12 := ui.NewLabel()
	label12.SetText("Q值(Q)")
	outputbox9 := ui.NewLineEdit()
	outputbox9.SetReadOnly(true)

	//计算结果
	CalBtn := ui.NewPushButton()
	CalBtn.SetText("计算")

	CalBtn.OnClicked(func() {
		//当两个都不为空的时候，弹出显示框
		if (strings.Compare(inputBox.Text(), "") != 0) && (strings.Compare(inputBox2.Text(), "") != 0) && (strings.Compare(inputBox3.Text(), "") != 0) {
			output := SecCoilInfoCal(inputBox.Text(), inputBox3.Text(), inputBox2.Text())
			outputbox1.SetText(output[0])
			outputbox2.SetText(output[1])
			outputbox3.SetText(output[2])
			outputbox4.SetText(output[3])
			outputbox5.SetText(output[4])
			outputbox6.SetText(output[5])
			outputbox7.SetText(output[6])
			outputbox8.SetText(output[7])
			outputbox9.SetText(output[8])

		} else {
			messagebox := ui.NewMessageBox()
			messagebox.SetText("必须要同时输入三个参数！")
			messagebox.Show()
			inputBox.Clear()
			inputBox2.Clear()
			inputBox3.Clear()
			outputbox1.Clear()
			outputbox2.Clear()
			outputbox3.Clear()
			outputbox4.Clear()
			outputbox5.Clear()
			outputbox6.Clear()
			outputbox7.Clear()
			outputbox8.Clear()
			outputbox9.Clear()
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
	hbox5.AddWidget(outputbox1)

	hbox6 := ui.NewHBoxLayout()
	hbox6.AddWidget(label5)
	hbox6.AddWidget(outputbox2)

	hbox7 := ui.NewHBoxLayout()
	hbox7.AddWidget(label6)
	hbox7.AddWidget(outputbox3)

	hbox8 := ui.NewHBoxLayout()
	hbox8.AddWidget(label7)
	hbox8.AddWidget(outputbox4)

	hbox9 := ui.NewHBoxLayout()
	hbox9.AddWidget(label8)
	hbox9.AddWidget(outputbox5)

	hbox10 := ui.NewHBoxLayout()
	hbox10.AddWidget(label9)
	hbox10.AddWidget(outputbox6)

	hbox11 := ui.NewHBoxLayout()
	hbox11.AddWidget(label10)
	hbox11.AddWidget(outputbox7)

	hbox12 := ui.NewHBoxLayout()
	hbox12.AddWidget(label11)
	hbox12.AddWidget(outputbox8)

	hbox13 := ui.NewHBoxLayout()
	hbox13.AddWidget(label12)
	hbox13.AddWidget(outputbox9)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddLayout(hbox2)
	vbox.AddLayout(hbox3)
	vbox.AddLayout(hbox4)
	vbox.AddLayout(hbox5)
	vbox.AddLayout(hbox6)
	vbox.AddLayout(hbox7)
	vbox.AddLayout(hbox8)
	vbox.AddLayout(hbox9)
	vbox.AddLayout(hbox10)
	vbox.AddLayout(hbox11)
	vbox.AddLayout(hbox12)
	vbox.AddLayout(hbox13)

	widget := ui.NewWidget()
	widget.SetLayout(vbox)
	widget.SetWindowTitle("次级线圈参数计算")
	widget.Show()
}
