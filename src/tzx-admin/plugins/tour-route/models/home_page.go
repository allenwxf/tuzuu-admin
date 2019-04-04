package models

import (
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetHome_pageTable() (home_pageTable models.Table) {

	home_pageTable.Info.FieldList = []types.Field{{
			Head:     language.Get("home_page-id"),
			Field:    "id",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("home_page-img_arr"),
			Field:    "img_arr",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	home_pageTable.Info.Table = "home_page"
	home_pageTable.Info.Title = language.Get("info-home_page")
	home_pageTable.Info.Description = "Home_page"

	home_pageTable.Form.FormList = []types.Form{{
			Head:     language.Get("home_page-id"),
			Field:    "id",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("home_page-img_arr"),
			Field:    "img_arr",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	home_pageTable.Form.Table = "home_page"
	home_pageTable.Form.Title = language.Get("form-home_page")
	home_pageTable.Form.Description = "Home_page"

	home_pageTable.ConnectionDriver = "mysql"

	return
}