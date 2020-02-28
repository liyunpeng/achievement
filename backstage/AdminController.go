package backstage

import (
	"achievement/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type AdminController struct {
	beego.Controller
}

//成员管理
//@router /achievement/showadmin [get]
func (this *AdminController)ShowAdmin()  {
	this.TplName = "backstage/setMenager.html"
}

//展示所有管理员
//@router /achievement/showmenager [get]
func (this *IdiomController)Idiom()  {
	page,pageerr := this.GetInt("page")
	limit,limiterr := this.GetInt("limit")
	var users []models.User
	if pageerr == nil && limiterr == nil && len(users) == 0  {
		users = models.NewUser().GetMenager(page,limit);
	}else {
		users = models.NewUser().GetMenager(1,10);
	}
	this.Data["json"] = map[string]interface{}{"code": 0, "count": models.NewUser().GetMenagerConut(),"data":users}
	this.ServeJSON()
	this.TplName = "backstage/setMenager.html"
}

//删除管理员
//@router /achievement/deletemenager [post]
func (this *IdiomController)Deletemenager()  {
	user := this.GetString("Id")
	userarrsy := strings.Split(user,",")
	for _,value := range userarrsy{
		id,_ :=strconv.Atoi(value)
		models.NewUser().DeleteMenagerConut(id)
	}
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "删除成功"}
	this.ServeJSON()
	this.TplName = "backstage/setMenager.html"
}

//添加管理员
//@router /achievement/addmenager [post]
func (this *IdiomController)Addtemenager()  {
	account := this.GetString("account")
	password := this.GetString("password")
	name := this.GetString("name")
	user := models.NewUser().AddMenagerExit(account)
	if len(user) == 1 {
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": "添加的账号已经存在"}
	}else {
		h := md5.New()
		h.Write([]byte(password))
		models.NewUser().AddMenager(account,hex.EncodeToString(h.Sum(nil)),name)
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": "添加成功"}
	}
	this.ServeJSON()
	this.TplName = "backstage/setMenager.html"
}
