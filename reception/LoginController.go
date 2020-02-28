package reception

import (
	"achievement/models"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type LoginContorller struct {
	beego.Controller
}

var cpt *captcha.Captcha

func init()  {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

/*返回登录首页*/
// @router /login [get]
func (c *LoginContorller)Login() {
	c.XSRFExpire = 7200
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "reception/index.html"
}

/*登录验证*/
//@router /login/judge [post]
func (c *LoginContorller)Post() {
	valid := validation.Validation{};
	valid.MaxSize(c.GetString("school"),11,"school").Message("您输入的学号不正确")
	valid.AlphaNumeric(c.GetString("school"),"number").Message("请您输入的学号必须是数字字符")
	valid.Match(c.GetString("password"),regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`),"password").Message("您输入的密码格式不正确")
	valid.Required(c.GetString("password"),"password").Message("请您输入密码")
	valid.Required(c.GetString("captcha"),"captcha").Message("请您输入验证码")
	captchaBoolean := cpt.Verify(c.GetString("captchaId"),c.GetString("captcha"));

	if valid.HasErrors() {
		for _,err := range valid.Errors {
			beego.Alert(err.Key+"+++++++"+err.Message)
		}
		return
	}
	h := md5.New()
	h.Write([]byte(c.GetString("password")))
	user := models.NewUser().ReceptionLoginJudge(c.GetString("school"),hex.EncodeToString(h.Sum(nil)))
	if !captchaBoolean || len(user) == 0 {
		if !captchaBoolean {
			c.Data["json"] = map[string]interface{}{"name": -1, "message": "你输入的验证码不正确"}
		}else {
			c.Data["json"] = map[string]interface{}{"name": 0, "message": "你输入的账号或密码不正确"}
		}
	}else {
		c.SetSession("account",c.GetString("school"))
		c.Data["json"] = map[string]interface{}{"name": 1, "message": "登陆成功"}
	}
	c.ServeJSON()
	c.TplName = "reception/index.html"
}

/*github第三方登录*/
//@router /github [get]
func (c *LoginContorller)GithubCallback()  {
	data := make(url.Values)
	fmt.Println("%s+=================", c.GetStrings("code"))
	data["code"] = c.GetStrings("code")
	data["client_id"] = []string{"e2cd16bc8436ab104785"}
	data["client_secret"] = []string{"b7c2a499a71869b2110208c35be5e5f71d77b307"}
	res,err := http.PostForm("https://github.com/login/oauth/access_token",data)
	if err != nil {
		//fmt.Println(err)
		c.TplName = "reception/index.html"
	}
	result,err := ioutil.ReadAll(res.Body)
	if err != nil {
		//fmt.Println(err)
		c.TplName = "reception/index.html"
	}
	//分割字符串access_token=4baa536a65c5b798e6776d7d84704e5bb94df98c&scope=user%3Aemail&token_type=bearer
	assessToken := strings.Split(strings.Split(string(result),"&")[0],"=")[1]
	fmt.Printf("%s---------------------", assessToken)
	//fmt.Println(res)


	u, _ := url.Parse("https://api.github.com/user")
	q := u.Query()
	q.Set("access_token", assessToken)
	u.RawQuery = q.Encode()
	res , er := http.Get(u.String())
	defer res.Body.Close()
	if er != nil {
		//fmt.Println(err)
		c.TplName = "reception/index.html"
	}
	result,error := ioutil.ReadAll(res.Body)
	if error != nil {
		//fmt.Println(err)
		c.TplName = "reception/index.html"
	}
	informtion :=make(map[string]string)
	json.Unmarshal(result, &informtion)
	fmt.Printf("%s+++++++++++++",informtion["message"])
	if informtion["message"] != "" {
		c.TplName = "reception/index.html"
	}
	fmt.Println(string(result))
	account := "20161514301"
	userInformation := 	models.NewStudent().GetInformation(account)
	exam := models.NewExam().GetExam(userInformation.Gradeid,userInformation.Clazzid)
	c.Data["exam"] = exam
	c.Data["user"] = userInformation
	c.TplName = "reception/information.html"
}

func (c *LoginContorller)LoginRedirect()  {
	fmt.Println("")
}