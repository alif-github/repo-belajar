package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/i18n"
	"strings"
	_ "test_api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	var (
		langs string
		errs  error
	)

	defer func() {
		if errs != nil {
			log := logs.NewLogger(1000)
			_ = log.SetLogger(logs.AdapterConsole, `{"color":true}`)
			log.EnableFuncCallDepth(true)
			log.Error(errs.Error())
		}
	}()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//--- Set Localize
	_ = beego.AddFuncMap("i18n", i18n.Tr)

	//--- Localize
	langs, errs = beego.AppConfig.String("langs")
	if errs != nil {
		return
	}

	//--- Get Localize
	langsArr := strings.Split(langs, "|")
	for _, lang := range langsArr {
		if errs = i18n.SetMessage(lang, "conf/"+lang+".ini"); errs != nil {
			return
		}
	}

	beego.Run()
}
