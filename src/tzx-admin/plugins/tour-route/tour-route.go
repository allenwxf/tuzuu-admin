package tour_route

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	routeModel "tzx-admin/plugins/tour-route/models"
)

type TourRoute struct {
	app      *context.App
	tableCfg map[string]models.TableGenerator
}

func (tourRoute *TourRoute) InitPlugin() {

	cfg := config.Get()

	// Init database
	// TODO: support multi driver
	//for _, databaseCfg := range cfg.DATABASE {
	//	db.GetConnectionByDriver(databaseCfg.DRIVER).InitDB(map[string]config.Database{
	//		"default": databaseCfg,
	//	})
	//}
	//
	//db.GetConnectionByDriver(cfg.DATABASE[0].DRIVER).InitDB(map[string]config.Database{
	//	"default": cfg.DATABASE[0],
	//	"3rdsrc": cfg.DATABASE[1],
	//})

	// Init router
	App.app = InitRouter("/" + cfg.PREFIX)

	models.SetGenerators(map[string]models.TableGenerator{
		"guide": routeModel.GetGuideTable,
		"home_page": routeModel.GetHome_pageTable,
		"poi": routeModel.GetPoiTable,
		"route": routeModel.GetRouteTable,
		"route_distance": routeModel.GetRoute_distanceTable,
		"route_poi": routeModel.GetRoute_poiTable,
		"task": routeModel.GetTaskTable,
	})
	models.InitTableList()

	cfg.PREFIX = "/" + cfg.PREFIX
	controller.SetConfig(cfg)

}

var App = new(TourRoute)

func NewTourRoute() *TourRoute {
	return App
}

func (tourRoute *TourRoute) GetRequest() []context.Path {
	return tourRoute.app.Requests
}

func (tourRoute *TourRoute) GetHandler(url, method string) context.Handler {
	return plugins.GetHandler(url, method, tourRoute.app)
}

