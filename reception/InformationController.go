package reception

import (
	"achievement/models"
	"fmt"
	"github.com/astaxie/beego"
)

type InformationController struct {
	beego.Controller
}

/*展示学生个人信息,学生的成绩*/
//@router /admin/information [get]
func (this *InformationController) Index() {
	// account,ok := this.GetSession("account").(string)
	//if !ok {
	//	return
	//}
	account := this.GetSession("account").(string)
	userInformation := 	models.NewStudent().GetInformation(account)
	exam := models.NewExam().GetExam(userInformation.Gradeid,userInformation.Clazzid)
	this.Data["exam"] = exam
	this.Data["user"] = userInformation
	this.TplName = "reception/information.html"
}

//@router /admin/getItem [post]
func (this *InformationController) GetScore() {
	eaxmid,_:=this.GetInt("item")
	studentid,_:=this.GetInt("id")
	score:=models.NewScore().GetAll(eaxmid,studentid)
	fmt.Println(score)
	this.Data["json"] = map[string]interface{}{"code": 1, "msg": score}
	this.ServeJSON()
	this.TplName = "reception/information.html"
}
