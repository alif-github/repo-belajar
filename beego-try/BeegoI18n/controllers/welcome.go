package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

var i18nCookieSet = false

type WelcomeController struct {
	beego.Controller
	i18n.Locale
}

func (c *WelcomeController) Post() {
	c.Data["langTemplateKey"] = c.GetString("lang")
	c.TplName = "welcome.tpl"
}

// @Title Get Data Welcome
// @Tag Login API
// @Description login
// @router /welcome [get]
func (c *WelcomeController) Get() {
	if !i18nCookieSet {
		c.Ctx.SetCookie("lang_choice", "fr-FR")
		i18nCookieSet = true
		_ = c.Ctx.Output.Body([]byte("i18n cookie successfully set!"))
	} else {
		c.Data["langTemplateKey"] = c.Ctx.GetCookie("lang_choice")
	}
	c.TplName = "welcome.tpl"
}
