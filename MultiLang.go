package main

import (
	"fmt"
	"github.com/Unknwon/i18n"
	"github.com/sndnvaps/tesla_calculator/setting"
)

// Controller struct -> i18n.Locale
type Controller struct {
	i18n.Locale
}

// Lang for i18n.Tr()
var Lang *Controller

//Init i18n, Set default Language
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
