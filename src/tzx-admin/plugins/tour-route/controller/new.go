package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/modules/menu"
	"strings"
	"net/http"
	"github.com/chenhg5/go-admin/plugins/admin/controller"
	"github.com/chenhg5/go-admin/plugins/admin/modules/file"
	)

// 显示新建
func ShowNewTourRoute(ctx *context.Context) {
	//prefix := "tour-route"
	prefix := ctx.Query("prefix")
	params := models.GetParam(ctx.Request.URL.Query())

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(controller.Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(
		tmpl,
		tmplName,
		user,
		types.Panel{
			Content: template.Get(controller.Config.THEME).Form().
				SetPrefix(controller.Config.PREFIX).
				SetContent(models.GetNewFormList(models.TableList[prefix].Form.FormList)).
				SetUrl(controller.Config.PREFIX + "/ext/new/" + prefix).
				SetToken(auth.TokenHelper.AddToken()).
				SetTitle("新建").
				SetInfoUrl(controller.Config.PREFIX + "/ext/show/" + prefix + params.GetRouteParamStr()).
				GetContent(),
			Description: models.TableList[prefix].Form.Description,
			Title:       models.TableList[prefix].Form.Title,
		},
		controller.Config,
		menu.GetGlobalMenu(user).SetActiveClass(
			strings.Replace(ctx.Path(), controller.Config.PREFIX, "", 1)))
	ctx.Html(http.StatusOK, buf.String())
}

// 新建
func NewTourRoute(ctx *context.Context) {
	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg": "新增失败",
		})
		return
	}

	//prefix := "tour-route"
	prefix := ctx.Query("prefix")

	form := ctx.Request.MultipartForm

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		file.GetFileEngine("local").Upload(form)
	}
	models.TableList[prefix].InsertDataFromDatabase("3rdsrc", (*form).Value)

	models.RefreshTableList()

	previous := ctx.FormValue("_previous_")
	prevUrlArr := strings.Split(previous, "?")
	params := models.GetParamFromUrl(previous)

	panelInfo := models.TableList[prefix].GetDataFromDatabase("3rdsrc", prevUrlArr[0], params)

	editUrl := controller.Config.PREFIX + "/ext/show/" + prefix + "/edit" + params.GetRouteParamStr()
	newUrl := controller.Config.PREFIX + "/ext/show/" + prefix + "/new" + params.GetRouteParamStr()
	deleteUrl := controller.Config.PREFIX + "/ext/delete/" + prefix

	dataTable := template.Get(controller.Config.THEME).
		DataTable().
		SetInfoList(panelInfo.InfoList).
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

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(controller.Config.THEME).GetTemplate(true)
	buffer := template.Excecute(
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
			strings.Replace(previous, controller.Config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buffer.String())
	ctx.AddHeader("X-PJAX-URL", previous)
}
