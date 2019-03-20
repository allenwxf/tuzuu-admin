package tour_route

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	myController "tzx-admin/plugins/tour-route/controller"
)

func InitRouter(prefix string) *context.App {
	app := context.NewApp()

	app.Group(prefix, GlobalErrorHandler)
	{
		authenticator := auth.SetPrefix(prefix).SetAuthFailCallback(func(ctx *context.Context) {
			ctx.Write(302, map[string]string{
				"Location": prefix + "/login",
			}, ``)
		}).SetPermissionDenyCallback(func(ctx *context.Context) {
			controller.ShowErrorPage(ctx, "permission denied")
		})

		app.Group("", authenticator.Middleware)
		{
			//app.GET(prefix+"/ext/show/tour-route", authenticator.Middleware(myController.ShowGetTourRoute))
			//app.GET(prefix+"/ext/show/tour-route/new", authenticator.Middleware(myController.ShowNewTourRoute))
			//app.POST(prefix+"/ext/new/tour-route", authenticator.Middleware(myController.NewTourRoute))
			//app.GET(prefix+"/ext/show/tour-route/edit", authenticator.Middleware(myController.ShowEditTourRoute))
			//app.POST(prefix+"/ext/edit/tour-route", authenticator.Middleware(myController.EditTourRoute))
			//app.POST(prefix+"/ext/delete/tour-route", authenticator.Middleware(myController.DelTourRoute))
			//app.POST(prefix+"/ext/show/tour-route/order", authenticator.Middleware(myController.OrderTourRoute))

			app.GET("/ext/show/:prefix", myController.ShowGetTourRoute)
			app.GET("/ext/show/:prefix/new", myController.ShowNewTourRoute)
			app.POST("/ext/new/:prefix", myController.NewTourRoute)
			app.GET("/ext/show/:prefix/edit", myController.ShowEditTourRoute)
			app.POST("/ext/edit/:prefix", myController.EditTourRoute)
			app.POST("/ext/delete/:prefix", myController.DelTourRoute)
		}

	}

	return app
}

func GlobalErrorHandler(h context.Handler) context.Handler {
	return func(ctx *context.Context) {
		defer controller.GlobalDeferHandler(ctx)
		h(ctx)
		return
	}
}