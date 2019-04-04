package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/chenhg5/go-admin/adapter/gin"
	"github.com/gin-contrib/cors"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
		"tzx-admin/plugins/tour-route"
)

func routeGroupV1(router *gin.RouterGroup) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE","OPTIONS"},
		AllowHeaders:  []string{"content-type", "Content-Length", "Content-Type", "Authorization", "Origin"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))
	//获取路线列表
	//router.Any("route/list", route.GetTourRoute)
	//新增路线
	//router.Any("route/add", route.AddNewTourRoute)
	//更新路线
	//router.Any("route/update", route.UpdateTourRoute)
	//删除路线
	//router.Any("route/delete", route.DelTourRoute)
}

func main() {
	r := gin.Default()
	eng := engine.Default()
	cfg := config.Config{
		TITLE: "TZX-ADMIN",
		LOGO: "<b>TZX</b>Admin",
		MINILOGO: "<b>T</b>A",
		THEME: "adminlte",
		INDEX: "/ext/show/home_page",
		DATABASE: []config.Database{
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "",
				NAME:         "tzx_sample",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "",
				NAME:         "tzx_route",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: "localhost", // the domain of cookie which be used when visiting your site.
		PREFIX: "admin",
		// STORE is important. And the directory should has permission to write.
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE: "cn",
		// debug mode
		DEBUG: true,
		// log file absolute path
		INFOLOG: "log/info.log",
		ACCESSLOG: "log/access.log",
		ERRORLOG: "log/error.log",
	}

	adminPlugin := admin.NewAdmin(datamodel.Generators)
	tourRoutePlugin := tour_route.NewTourRoute()

	eng.AddConfig(cfg).AddPlugins(adminPlugin, tourRoutePlugin).Use(r)

	r.Run(":9033")

}