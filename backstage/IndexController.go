package backstage

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

//后台页面展示
//@router /achievement/index [get]
func (this *IndexController)Index() {
	this.TplName = "backstage/index.html"
}

