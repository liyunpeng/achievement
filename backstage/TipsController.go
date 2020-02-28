package backstage

import "github.com/astaxie/beego"

type TipsController struct {
	beego.Controller
}

//消息提示
//@router /achievement/showitips [get]
func (this *TipsController)ShowTips()  {
	this.TplName = "backstage/welcome.html"
}
