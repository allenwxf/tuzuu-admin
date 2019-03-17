package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/modules/auth"
	"net/http"
)

func DelTourRoute(ctx *context.Context) {
	prefix := "tour-route"

	models.TableList[prefix].
		DeleteDataFromDatabase("3rdsrc", ctx.FormValue("id"))

	newToken := auth.TokenHelper.AddToken()

	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "删除成功", // TODO: 配置为根据语言返回内容
		"data": newToken,
	})
	return
}
