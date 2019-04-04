package models

import (
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetRoute_poiTable() (route_poiTable models.Table) {

	route_poiTable.Info.FieldList = []types.Field{{
			Head:     language.Get("route_poi-id"),
			Field:    "id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-route_id"),
			Field:    "route_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-poi_id"),
			Field:    "poi_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-sort"),
			Field:    "sort",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-time"),
			Field:    "time",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	route_poiTable.Info.Table = "route_poi"
	route_poiTable.Info.Title = language.Get("info-route_poi")
	route_poiTable.Info.Description = "Route_poi"

	route_poiTable.Form.FormList = []types.Form{{
			Head:     language.Get("route_poi-id"),
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-route_id"),
			Field:    "route_id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-poi_id"),
			Field:    "poi_id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-sort"),
			Field:    "sort",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_poi-time"),
			Field:    "time",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	route_poiTable.Form.Table = "route_poi"
	route_poiTable.Form.Title = language.Get("form-route_poi")
	route_poiTable.Form.Description = "Route_poi"

	route_poiTable.ConnectionDriver = "mysql"

	return
}