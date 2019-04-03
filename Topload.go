package main

import (
	"github.com/visualfc/goqt/ui"
	"log"
)

type ToploadForm struct { //此为widget,
	*ui.QWidget
	btn1 *ui.QPushButton //
	btn2 *ui.QPushButton
}

func NewToploadForm() (*ToploadForm, error) {
	w := &ToploadForm{}
	w.QWidget = ui.NewWidget()

	w.SetFixedWidth(200)
	w.SetFixedHeight(80)

	w.btn1 = ui.NewPushButton()
	w.btn1.SetText("Toroid Cap")

	w.btn2 = ui.NewPushButton()
	w.btn2.SetText("Sphere Cap")

	w.btn1.OnClicked(func() {
		toroid, err := NewToroidForm()
		if err != nil {
			log.Fatalln(err)
		}
		toroid.Show()
	})

	w.btn2.OnClicked(func() {

		sphere, err := NewSphereForm()
		if err != nil {
			log.Fatalln(err)
		}
		sphere.Show()
	})

	hbox := ui.NewHBoxLayout()
	hbox.AddWidget(w.btn1)
	hbox.AddWidget(w.btn2)

	vbox := ui.NewVBoxLayout()
	vbox.AddLayout(hbox)

	w.SetLayout(vbox)
	w.SetWindowTitle("Topload Cap Calculator")

	return w, nil
}
