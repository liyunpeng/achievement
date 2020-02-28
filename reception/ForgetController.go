package reception

import (
	"achievement/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
)

type ForgetController struct {
	beego.Controller
}

//@router /forget [get]
func (this *ForgetController) Index() {
	//this.TplName = "reception/forget.html"
	//c.TplName = "reception/index.html"
	//err := this.Email("2181550591@qq.com", "xsywrvvxqzwxebbi", "smtp.qq.com:25", "3438196289@qq.com", "大橘猫", "sssss", "html")
	//r := rand.New(rand.NewSource(time.Now().Unix()))

	this.TplName = "reception/forget.html"
}

//err := SendMail("发送的邮箱", "发送的邮箱密码", "smtp.qq.com:25", "目标邮箱", "邮件标题", "邮件内容", "html")
func (this *ForgetController)Email(user, password, host, to, subject, body, mailtype string) error  {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

//@router /forget/look [post]
func (this *ForgetController) Forget() {
	qq := this.GetString("qq")
	code,_ := this.GetInt("code")
	password := this.GetString("password")
	if code == this.GetSession("code") {
		//md5加密
		h := md5.New()
		h.Write([]byte(password))
		student :=models.NewStudent().GetAccount(qq)
		err := models.NewStudent().ModifyBatchStudent(student.Number,hex.EncodeToString(h.Sum(nil)))
		fmt.Println(err)
		if err == nil {
			this.Data["json"] = map[string]interface{}{"name": 1, "message": "更新秘码成功"}
		}else {
			this.Data["json"] = map[string]interface{}{"name": 1, "message": "更新秘码成失败"}
		}
	}else {
		this.Data["json"] = map[string]interface{}{"name": 1, "message": "验证码错误"}
	}
	this.ServeJSON()
	this.TplName = "reception/forget.html"
}

//@router /forget/getcode [post]
func (this *ForgetController) GetCode() {
	qq := this.GetString("qq")
	//生成随即数
	code := rand.Intn(9000)+1000
	//发送邮件
	err := this.Email("2181550591@qq.com", "xsywrvvxqzwxebbi", "smtp.qq.com:25", qq, "大橘猫",strconv.Itoa(code), "html")
	if err == nil {
		//保存验证码
		this.SetSession("code",code)
		fmt.Println(this.GetSession("code"))
		this.Data["json"] = map[string]interface{}{"name": 1, "message": "获取验证成功"}
	}else {
		this.Data["json"] = map[string]interface{}{"name": 1, "message": "获取验证失败"}
	}
	this.ServeJSON()
	this.TplName = "reception/forget.html"
}