package backstage

import (
	"achievement/models"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"strconv"
)

type IdiomController struct {
	beego.Controller
}

//展示成员
//@router /achievement/showidiom [get]
func (this *IdiomController)ShowIdiom()  {
	this.Data["grade"] = models.NewGrade().GetGrade()
	this.Data["class"] = models.NewClazz().GetClazz()
	this.TplName = "backstage/people.html"
}



//添加成员
//@router /achievement/addidiom [post]
func (this *IdiomController)AddIdiom()  {
	clazzid,_ := this.GetInt("clazzid")
	gradeid,_ :=this.GetInt("gradeid")
	err := models.NewStudent().AddStudent(this.GetString("number"),this.GetString("name"),this.GetString("sex"),this.GetString("phone"),this.GetString("qq"),clazzid,gradeid)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"name": 1, "message": err}
	}else {
		this.Data["json"] = map[string]interface{}{"name": 1, "message": "添加成功"}
	}
	this.ServeJSON()
	this.TplName = "backstage/people.html"
}

// 上传Excel 表格
//@router /achievement/excel [post]
func (this *IdiomController)AddExcel()  {
	file, h, error := this.GetFile("file")                  //获取上传的文件
	if error == nil {
		if h.Header["Content-Type"][0] == "application/wps-office.xlsx" {
			path := "./upload/"+h.Filename
			file.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
			error:=this.SaveToFile("file", path)                    //存文件
			if error == nil {
				xlsx, _ := excelize.OpenFile(path)
				index := xlsx.GetSheetIndex("Sheet1")
				rows := xlsx.GetRows("Sheet" + strconv.Itoa(index))
				this.Data["json"] = map[string]interface{}{"name": 0, "message": xlsx.GetRows("Sheet1")}
				student :=  make([]models.Student,len(rows)-1)
				clazz := models.NewClazz().GetClazz()
				grade := models.NewGrade().GetGrade()
				for index, row := range rows {
					if index != 0 {
						students := models.Student{}
						students.Number = row[0];
						students.Name = row[1];
						students.Sex = row[2];
						students.Phone = row[3];
						students.Qq = row[4];
						for _, grades := range grade {
							if grades.GradeName == row[6] {
								students.Gradeid = grades.Id
							}
						}
						for _, class := range clazz {
							if class.ClazzName == row[5] && students.Gradeid == class.Gradeid {
								students.Clazzid = class.Id
							}
						}
						student[index-1] = students
					}
				}
				error :=models.NewStudent().AddBatchStudent(student)
				if error == nil {
					this.Data["json"] = map[string]interface{}{"name": 0, "message": "上传文件成功"}
				}else{
					this.Data["json"] = map[string]interface{}{"name": 0, "message": error}
				}
			}else {
				this.Data["json"] = map[string]interface{}{"name": 0, "message": "上传文件失败,请重新上传"}
			}
		}else {
			this.Data["json"] = map[string]interface{}{"name": 0, "message": "您上传文件的文件格式不正确,请上传Excel文件"}
		}
	}else {
		this.Data["json"] = map[string]interface{}{"name": 0, "message": "上传文件失败,请重新上传"}
	}
	this.ServeJSON()
	this.TplName = "backstage/people.html"
}
