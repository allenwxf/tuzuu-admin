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

// 显示编辑
func ShowEditTourRoute(ctx *context.Context) {
	//prefix := "tour-route"
	prefix := ctx.Query("prefix")

	formData, title, description := models.TableList[prefix].
		GetDataFromDatabaseWithId("3rdsrc", ctx.Query("id"))
	params := models.GetParam(ctx.Request.URL.Query())

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(controller.Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(
		tmpl,
		tmplName,
		user,
		types.Panel{
			Content: template.Get(controller.Config.THEME).Form().
				SetContent(formData).
				SetPrefix(controller.Config.PREFIX).
				SetUrl(controller.Config.PREFIX + "/ext/edit/" + prefix).
				SetToken(auth.TokenHelper.AddToken()).
				SetInfoUrl(controller.Config.PREFIX + "/ext/show/" + prefix + params.GetRouteParamStr()).
				GetContent(),
			Description: description,
			Title: title,
		},
		controller.Config,
		menu.GetGlobalMenu(user).SetActiveClass(
			strings.Replace(ctx.Path(), controller.Config.PREFIX, "", 1)))
	ctx.Html(http.StatusOK, buf.String())
}

// 编辑
func EditTourRoute(ctx *context.Context) {
	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "编辑失败",
		})
		return
	}

	//prefix := "tour-route"
	prefix := ctx.Query("prefix")

	form := ctx.Request.MultipartForm

	menu.GlobalMenu.SetActiveClass(
		strings.Replace(ctx.Path(), controller.Config.PREFIX, "", 1))

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		file.GetFileEngine("local").Upload(form)
	}

	models.TableList[prefix].UpdateDataFromDatabase("3rdsrc", (*form).Value)

	models.RefreshTableList()

	previous := ctx.FormValue("_previous_")
	prevUrlArr := strings.Split(previous, "?")
	params := models.GetParamFromUrl(previous)

	previous = controller.Config.PREFIX + "/ext/show/" + prefix + params.GetRouteParamStr()
	editUrl := controller.Config.PREFIX + "/ext/show/" + prefix + "/edit" + params.GetRouteParamStr()
	newUrl := controller.Config.PREFIX + "/ext/show/" + prefix + "/new" + params.GetRouteParamStr()
	deleteUrl := controller.Config.PREFIX + "/ext/delete/" + prefix

	panelInfo := models.TableList[prefix].GetDataFromDatabase("3rdsrc", prevUrlArr[0], params)

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
			strings.Replace(previous, controller.Config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buf.String())
	ctx.AddHeader("X-PJAX-URL", previous)
}
