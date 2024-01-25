package main

import (
	_ "BeegoI18n/routers"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

// @Title Started From Here!
func main() {
	_ = beego.AddFuncMap("i18n", i18n.Tr)

	//--- Logger
	log := logs.NewLogger(1000)
	_ = log.SetLogger(logs.AdapterConsole, `{"color":true}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
	log.Error("Reading File")
	beego.Run()
}
