package main

import (
	"encoding/gob"
	"github.com/Spieny/AquaMall/common"
	"github.com/Spieny/AquaMall/models"
	_ "github.com/Spieny/AquaMall/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	//添加方法到map,用于前端html调用
	beego.AddFuncMap("timestampToDate", common.TimestampToDate)
	models.DB.LogMode(true)
	beego.AddFuncMap("formatImage", common.FormatImage)
	beego.AddFuncMap("mul", common.Mul)
	beego.AddFuncMap("formatAttribute", common.FormatAttribute)
	beego.AddFuncMap("setting", models.GetSettingByColumn)

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true, // 允许Cookie
	}))

	//注册模型
	gob.Register(models.Administrator{})

	//配置redis用于存储session
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"

	//本地启动，请设置如下
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
}
