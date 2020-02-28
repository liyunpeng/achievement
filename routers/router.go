package routers

import (
	"achievement/backstage"
	"achievement/reception"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	/*前台路由*/
	beego.InsertFilter("/reception/admin/*", beego.BeforeRouter, func(ctx *context.Context) { //前台路由过滤
		_, ok := ctx.Input.Session("account").(string)
		if !ok {
			ctx.Redirect(302, "/reception/login")
			fmt.Println("验证不通过")
		}
		fmt.Println("我验证了路由")
	})

	ns := beego.NewNamespace("reception",
		beego.NSInclude(
			&reception.LoginContorller{},
			&reception.InformationController{},
			&reception.ForgetController{},
		),
	)
	beego.AddNamespace(ns)



	/*后台路由*/
	beego.InsertFilter("/back/achievement/*", beego.BeforeRouter, func(ctx *context.Context) { //前台路由过滤
		_, ok := ctx.Input.Session("type").(string)
		if !ok {
			ctx.Redirect(302, "/back/backlogin")
		}
	})
	back := beego.NewNamespace("back",
		beego.NSInclude(
			&backstage.LoginController{},
			&backstage.IndexController{},
			&backstage.ScroeController{},
			&backstage.IdiomController{},
			&backstage.ItemController{},
			&backstage.AdminController{},
			&backstage.TipsController{},
		),
	)
	beego.AddNamespace(back)
}
