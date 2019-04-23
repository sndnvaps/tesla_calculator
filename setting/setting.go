// Copyright 2019 sndnvaps
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package setting

import (
	"fmt"

	"github.com/Unknwon/com"
	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	Langs, Names []string
	DefLang      string
    LangsMap map[string]string
    LangsMapNum map[string]int
)

func init() {

	var err error
    CFG_DATA, _ := Asset("conf/app.ini")
	Cfg, err = ini.Load(CFG_DATA)
	if err != nil {
		panic(fmt.Errorf("fail to load config file '%s': %v", CFG_PATH, err))
	}
	if com.IsFile(CFG_CUSTOM_PATH) {
		if err = Cfg.Append(CFG_CUSTOM_PATH); err != nil {
			panic(fmt.Errorf("fail to load config file '%s': %v", CFG_CUSTOM_PATH, err))
		}
	}

        LangsMap = make(map[string]string)
        LangsMapNum = make(map[string]int)
        
	Langs = Cfg.Section("i18n").Key("langs").Strings(",")
	Names = Cfg.Section("i18n").Key("names").Strings(",")

        for key , val := range Langs {
               LangsMap[Names[key]] =  val
               LangsMapNum[Langs[key]] = key
        }

	DefLang = Cfg.Section("i18n").Key("defaultLang").String()
}
