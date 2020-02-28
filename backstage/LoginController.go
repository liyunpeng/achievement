package backstage

import (
	"achievement/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	beego.Controller
}

//后台登陆
//@router /backlogin [get]
func (this *LoginController)Login()  {
	this.TplName = "backstage/login.html"
}

//@router /backlogin/judge [post]
func (this *LoginController)JudgeLogin()  {
	valid := validation.Validation{};
	valid.Required(this.GetString("password"),"password").Message("请您输入密码")
	valid.Required(this.GetString("captcha"),"captcha").Message("请您输入验证码")
	if valid.HasErrors() {
		for _,err := range valid.Errors {
			beego.Alert(err.Key+"+++++++"+err.Message)
		}
		this.Data["json"] = map[string]interface{}{"name": 0, "message": "你输入的账号或密码不正确"}
	}
	h := md5.New()
	h.Write([]byte(this.GetString("password")))
	user := models.NewUser().LoginJudge(this.GetString("school"),hex.EncodeToString(h.Sum(nil)))
	if len(user) == 0 {

		this.Data["json"] = map[string]interface{}{"name": 0, "message": "你输入的账号或密码不正确"}
	}else {
		if user[0].Type == 3 {
			this.Data["json"] = map[string]interface{}{"name": -1, "message": "您没有权限登陆"}
		}else {
			this.SetSession("type",this.GetString("school"))
			this.Data["json"] = map[string]interface{}{"name": 1, "message": "登陆成功"}
		}
	}
	this.ServeJSON()
	this.TplName = "backstage/login.html"
}

//退出
//@router /leavelogin [get]
func (this *LoginController)LeaveLogin()  {
	this.DestroySession()
	this.TplName = "backstage/login.html"
}

