package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"log"
	"xAdmin/config"
	orm "xAdmin/database"
	"xAdmin/models"
	"xAdmin/router"
)

var Cron = cron.New()

func init() {
	//0 */3 * * * ?    0 15,30,45 * * * ?
	//Cron.AddFunc("0 15,30,45 * * * ?", auth.CornTask)
	//Cron.AddFunc("@daily", auth.CornTaskGKC)

	//Cron.Start()
}

// @title tco-admin API
// @version 0.0.1
// @description GKC管理系统的接口文档

// @description
// @license.name MIT
// @license.url
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	gin.SetMode(gin.DebugMode)

	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}
	//models.RunInit()
	//auth.AddStatement()
	r := router.InitRouter()
	//models.SaveAddressAssets()
	//go Test()
	//auth.CornTaskGKC()

	// 定期更新过期订单状态
	go router.CheckInvestmentExpiration()

	defer orm.Eloquent.Close()
	if err := r.Run(config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port); err != nil {
		log.Fatal(err)
	}
}
