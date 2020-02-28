package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["achievement/backstage:AdminController"] = append(beego.GlobalControllerRouter["achievement/backstage:AdminController"],
        beego.ControllerComments{
            Method: "ShowAdmin",
            Router: `/achievement/showadmin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "AddIdiom",
            Router: `/achievement/addidiom`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "Addtemenager",
            Router: `/achievement/addmenager`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "AddScore",
            Router: `/achievement/addscore`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "Deletemenager",
            Router: `/achievement/deletemenager`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "AddExcel",
            Router: `/achievement/excel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "SelectScore",
            Router: `/achievement/selectallscore`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "ShowIdiom",
            Router: `/achievement/showidiom`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IdiomController"] = append(beego.GlobalControllerRouter["achievement/backstage:IdiomController"],
        beego.ControllerComments{
            Method: "Idiom",
            Router: `/achievement/showmenager`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IndexController"] = append(beego.GlobalControllerRouter["achievement/backstage:IndexController"],
        beego.ControllerComments{
            Method: "AddExam",
            Router: `/achievement/addexam`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:IndexController"] = append(beego.GlobalControllerRouter["achievement/backstage:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/achievement/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:ItemController"] = append(beego.GlobalControllerRouter["achievement/backstage:ItemController"],
        beego.ControllerComments{
            Method: "ShowItem",
            Router: `/achievement/showitem`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:LoginController"] = append(beego.GlobalControllerRouter["achievement/backstage:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/backlogin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:LoginController"] = append(beego.GlobalControllerRouter["achievement/backstage:LoginController"],
        beego.ControllerComments{
            Method: "JudgeLogin",
            Router: `/backlogin/judge`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:LoginController"] = append(beego.GlobalControllerRouter["achievement/backstage:LoginController"],
        beego.ControllerComments{
            Method: "LeaveLogin",
            Router: `/leavelogin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:ScroeController"] = append(beego.GlobalControllerRouter["achievement/backstage:ScroeController"],
        beego.ControllerComments{
            Method: "ShowScore",
            Router: `/achievement/showscore`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["achievement/backstage:TipsController"] = append(beego.GlobalControllerRouter["achievement/backstage:TipsController"],
        beego.ControllerComments{
            Method: "ShowTips",
            Router: `/achievement/showitips`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
