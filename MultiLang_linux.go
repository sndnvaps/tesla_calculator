//+build !windows

package main

import (
	"fmt"
	"github.com/Unknwon/i18n"
	"github.com/sndnvaps/tesla_calculator/setting"
	"strings"
)

type Controller struct {
	i18n.Locale
}

var Lang *Controller

func InitI18n() error {

	langs := setting.Langs
	for _, lang := range langs {
		datapath := fmt.Sprintf("%s/conf/locale/locale_%s.ini", setting.GetCurrentDic(), strings.ToLower(lang))

		i18n.SetMessage(lang, datapath)
	}
	Lang = &Controller{
		Locale: i18n.Locale{setting.DefLang}, //set default in conf/app.ini
	}

	return i18n.ReloadLangs(langs...)
}
