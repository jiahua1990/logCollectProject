package router

import (
	"MyGitHubProject/logCollectProject/webAdmin/controller/AppController"
	"MyGitHubProject/logCollectProject/webAdmin/controller/LogController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/list", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/apply", &AppController.AppController{}, "*:AppApply")

	beego.Router("/app/create", &AppController.AppController{}, "*:AppCreate")

	beego.Router("/log/apply", &LogController.LogController{}, "*:LogApply")
	beego.Router("/log/list", &LogController.LogController{}, "*:LogList")
	beego.Router("/log/create", &LogController.LogController{}, "*:LogCreate")
}
