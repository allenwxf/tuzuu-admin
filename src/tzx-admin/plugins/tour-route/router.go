package tour_route

import (
	myController "tzx-admin/plugins/tour-route/controller"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func(ctx *context.Context) {
		ctx.Write(302, map[string]string{
			"Location": prefix + "/login",
		}, ``)
	}).SetPermissionDenyCallback(func(ctx *context.Context) {
		controller.ShowErrorPage(ctx, "permission denied")
	})

	app.GET(prefix+"/ext/show/tour-route", authenticator.Middleware(myController.ShowGetTourRoute))
	app.GET(prefix+"/ext/show/tour-route/new", authenticator.Middleware(myController.ShowNewTourRoute))
	app.POST(prefix+"/ext/new/tour-route", authenticator.Middleware(myController.NewTourRoute))
	app.GET(prefix+"/ext/show/tour-route/edit", authenticator.Middleware(myController.ShowEditTourRoute))
	app.POST(prefix+"/ext/edit/tour-route", authenticator.Middleware(myController.EditTourRoute))
	app.POST(prefix+"/ext/delete/tour-route", authenticator.Middleware(myController.DelTourRoute))
	app.POST(prefix+"/ext/show/tour-route/order", authenticator.Middleware(myController.OrderTourRoute))

	return app
}
