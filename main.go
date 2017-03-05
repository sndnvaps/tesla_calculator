package main

import (
	"fmt"
	"os"
	//"time"
	"runtime"
	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.RunEx(os.Args, main_ui)
}

func version() {
	info := fmt.Sprintf("GoQt Version %v \nQt Version %v\ngo verison %v %v/%v",
		ui.Version(),
		ui.QtVersion(),
		runtime.Version(), runtime.GOOS, runtime.GOARCH)

	lable := ui.NewLabel()
	lable.SetText(info)

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(lable)

	widget := ui.NewWidget()
	widget.SetLayout(hbox)
	widget.Show()
}

func main_ui() {
	btn := ui.NewPushButton()
	btn.SetText("计算初始参数")

	clear := ui.NewPushButton()
	clear.SetText("计算顶端电容")

	SecCoilInfoBtn := ui.NewPushButton()
	SecCoilInfoBtn.SetText("计算次级参数")

	ArcLengthBtn := ui.NewPushButton()
	ArcLengthBtn.SetText("估算电弧长度")

	CoefficientOfCouplinBtn := ui.NewPushButton()
	CoefficientOfCouplinBtn.SetText("初始和次级耦合度计算")
	/*
		edit := ui.NewPlainTextEdit()
		edit.SetReadOnly(true)
	*/
	btn.OnClicked(func() {
		/*
			for i := 0; i < 10; i++ {
				go func(i int) {
					start := time.Now()
					<-time.After(1e7)
					offset := time.Now().Sub(start)
					ui.Async(func() {
						edit.AppendPlainText(fmt.Sprintf("Task %v %v", i, offset))
					})
				}(i)
			}
		*/
		version()
	})

	clear.OnClicked(func() {
		//edit.Clear()
		version()
	})

	SecCoilInfoBtn.OnClicked(func() {
		version()
	})

	ArcLengthBtn.OnClicked(func() {
		SparkLengthForm()
	})

	CoefficientOfCouplinBtn.OnClicked(func() {
		CoefficientOfCouplinForm()
	})


	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(btn)
	hbox.AddWidget(clear)

	hbox2 := ui.NewHBoxLayout()
	hbox2.AddWidget(SecCoilInfoBtn)
	hbox2.AddWidget(ArcLengthBtn)

	hbox3 := ui.NewHBoxLayout()
	hbox3.AddWidget(CoefficientOfCouplinBtn)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)
	vbox.AddLayout(hbox2)
	vbox.AddLayout(hbox3)
	//vbox.AddWidget(edit)

	widget := ui.NewWidget()
	widget.SetLayout(vbox)
	widget.Show()
}
