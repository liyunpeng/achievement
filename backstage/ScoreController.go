package backstage

import (
	"achievement/models"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"strconv"
)

type ScroeController struct {
	beego.Controller
}

//展示成绩
//@router /achievement/showscore [get]
func (this *ScroeController)ShowScore()  {
	this.Data["clazz"] = models.NewClazz().GetClazz()
	this.TplName ="backstage/score.html"
}

//成绩录入
//@router /achievement/addscore [post]
func (this *IdiomController)AddScore()  {
	file, h, error := this.GetFile("file")                  //获取上传的文件
	if error == nil {
		if h.Header["Content-Type"][0] == "application/wps-office.xlsx" {
			path := "./upload/"+h.Filename
			file.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
			error:=this.SaveToFile("file", path)                    //存文件
			if error == nil {
				fmt.Println(path)
				xlsx, _ := excelize.OpenFile(path)
				index := xlsx.GetSheetIndex("Sheet1")
				rows := xlsx.GetRows("Sheet" + strconv.Itoa(index))
				score :=  make([]models.Score,len(rows)-1)
				for index, row := range rows {
					if index != 0 {
						scores := models.Score{}
						scores.Studentid = models.NewStudent().GetId(row[0])
						scores.Examid = models.NewExam().GetId(row[1])
						scores.Courseid = models.NewScore().GetId(row[2])
						scores.Score,_ = strconv.Atoi(row[3])
						score[index-1] = scores
					}
				}
				error :=models.NewScore().AddBatchSCore(score)
				if error == nil {
					this.Data["json"] = map[string]interface{}{"name": 0, "message": "上传文件成功"}
					this.Delete()
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
	this.TplName = "backstage/score.html"
}

//成绩查询
//@router /achievement/selectallscore [get]
func (this *IdiomController)SelectScore()  {
	page,pageerr := this.GetInt("page")
	limit,limiterr := this.GetInt("limit")
	class,_ := this.GetInt("class")
	school := this.GetString("school")
	var score []models.ScoreInformation
	if pageerr == nil && limiterr == nil && len(score) == 0  {
		score = models.NewScore().SelectAll(page,limit,class,school);
	}else {
		score = models.NewScore().SelectAll(1,10,class,school);
	}
	this.Data["json"] = map[string]interface{}{"code": 0, "count": models.NewScore().SelectCount(),"data":score}
	this.ServeJSON()
	this.TplName = "backstage/score.html"
}