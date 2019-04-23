package main

import (
	"fmt"
	"github.com/Unknwon/i18n"
	"github.com/sndnvaps/tesla_calculator/setting"
)

type Controller struct {
	i18n.Locale
}

var Lang *Controller

func InitI18n() error {

	langs := setting.Langs
	for _, lang := range langs {
		datapath := fmt.Sprintf("conf/locale/locale_%s.ini", lang)
		data, _ := setting.Asset(datapath)
		i18n.SetMessage(lang, data)
	}
	Lang = &Controller{
		Locale: i18n.Locale{setting.DefLang},
	}
	return i18n.ReloadLangs(langs...)
}
