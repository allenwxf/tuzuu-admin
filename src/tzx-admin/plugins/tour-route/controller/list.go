package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"strings"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"net/http"
	"github.com/chenhg5/go-admin/context"
	)

func ShowGetTourRoute(ctx *context.Context) {
	//prefix := "tour-route"
	prefix := ctx.Query("prefix")
	params := models.GetParam(ctx.Request.URL.Query())

	user := auth.Auth(ctx)
	menu.GlobalMenu.SetActiveClass(
		strings.Replace(ctx.Path(), controller.Config.PREFIX, "", 1))

	newUrl := controller.Config.PREFIX + "/ext/show/" + prefix + "/new" + params.GetRouteParamStr()
	editUrl := controller.Config.PREFIX + "/ext/show/" + prefix + "/edit" + params.GetRouteParamStr()
	deleteUrl := controller.Config.PREFIX + "/ext/delete/" + prefix + params.GetRouteParamStr()
	//orderUrl := controller.Config.PREFIX + "/tour-route/order"

	panelInfo := models.TableList[prefix].GetDataFromDatabase("3rdsrc", ctx.Path(), params)

	dataTable := template.Get(controller.Config.THEME).
		DataTable().
		SetInfoList(panelInfo.InfoList).
		SetFilters(models.TableList[prefix].GetFiltersMap()).
		SetInfoUrl(controller.Config.PREFIX + "/show/" + prefix).
		SetThead(panelInfo.Thead).
		SetEditUrl(editUrl).
		SetNewUrl(newUrl).
		SetDeleteUrl(deleteUrl)

	table := dataTable.GetContent()

	box := template.Get(controller.Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(panelInfo.Paginator.GetContent()).
		GetContent()

	tmpl, tmplName := template.Get(controller.Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(
		tmpl,
		tmplName,
		user,
		types.Panel{
			Content:     box,
			Description: panelInfo.Description,
			Title:       panelInfo.Title,
		},
		controller.Config,
		menu.GetGlobalMenu(user).SetActiveClass(
			strings.Replace(ctx.Path(), controller.Config.PREFIX, "", 1)))
	ctx.Html(http.StatusOK, buf.String())

}