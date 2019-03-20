package models

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/modules/language"
)

func GetRoutepoiTable() (routepoiTable models.Table) {

	routepoiTable.Info.FieldList = []types.Field{{
			Head:     "Id",
			Field:    "Id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-RouteId"),
			Field:    "RouteId",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-PoiId"),
			Field:    "PoiId",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-Sort"),
			Field:    "Sort",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-Time"),
			Field:    "Time",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	routepoiTable.Info.Table = "routepoi"
	routepoiTable.Info.Title = "Routepoi"
	routepoiTable.Info.Description = "管理poi route列表"

	routepoiTable.Form.FormList = []types.Form{{
			Head:     "Id",
			Field:    "Id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-RouteId"),
			Field:    "RouteId",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-PoiId"),
			Field:    "PoiId",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-Sort"),
			Field:    "Sort",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("RoutePoi-Time"),
			Field:    "Time",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	routepoiTable.Form.Table = "routepoi"
	routepoiTable.Form.Title = "Routepoi"
	routepoiTable.Form.Description = "管理poi route信息"

	routepoiTable.ConnectionDriver = "mysql"

	return
}