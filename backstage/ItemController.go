package backstage

import (
	"achievement/models"
	"fmt"
	"github.com/astaxie/beego"
)

type ItemController struct {
	beego.Controller
}

//展示学期
//@router /achievement/showitem [get]
func (this *ItemController)ShowItem()  {
	item := models.NewExam().GetAllExam();
	this.Data["item"] = item
	this.Data["number"] = len(item)
	this.TplName = "backstage/term.html"
}

//添加考试信息
//@router /achievement/addexam [post]
func (this *IndexController)AddExam()  {
	error := models.NewExam().AddExam(this.GetString("type"),this.GetString("time"))
	fmt.Println(error)
	if error == nil {
		this.Data["json"] = map[string]interface{}{"name": 2, "message": "考试添加成功"}
	}else {
		this.Data["json"] = map[string]interface{}{"name": 1, "message": "考试添加失败"}
	}
	this.ServeJSON()
	this.TplName = "backstage/index.html"
}

