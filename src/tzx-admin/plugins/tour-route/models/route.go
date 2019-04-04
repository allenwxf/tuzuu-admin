package models

import (
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetRouteTable() (routeTable models.Table) {

	routeTable.Info.FieldList = []types.Field{{
			Head:     language.Get("route-id"),
			Field:    "id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-city"),
			Field:    "city",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-name"),
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-video"),
			Field:    "video",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-imgs"),
			Field:    "imgs",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-price"),
			Field:    "price",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-guide_id"),
			Field:    "guide_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-change_time"),
			Field:    "change_time",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-time"),
			Field:    "time",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr"),
			Field:    "img_arr",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-label"),
			Field:    "label",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr2"),
			Field:    "img_arr2",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr3"),
			Field:    "img_arr3",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr4"),
			Field:    "img_arr4",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-view_count"),
			Field:    "view_count",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-sale_count"),
			Field:    "sale_count",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	routeTable.Info.Table = "route"
	routeTable.Info.Title = language.Get("info-route")
	routeTable.Info.Description = "Route"

	routeTable.Form.FormList = []types.Form{{
			Head:     language.Get("route-id"),
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-city"),
			Field:    "city",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-name"),
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-video"),
			Field:    "video",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-imgs"),
			Field:    "imgs",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-price"),
			Field:    "price",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-guide_id"),
			Field:    "guide_id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-change_time"),
			Field:    "change_time",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-time"),
			Field:    "time",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr"),
			Field:    "img_arr",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-label"),
			Field:    "label",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr2"),
			Field:    "img_arr2",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr3"),
			Field:    "img_arr3",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-img_arr4"),
			Field:    "img_arr4",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-view_count"),
			Field:    "view_count",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route-sale_count"),
			Field:    "sale_count",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	routeTable.Form.Table = "route"
	routeTable.Form.Title = language.Get("form-route")
	routeTable.Form.Description = "Route"

	routeTable.ConnectionDriver = "mysql"

	return
}