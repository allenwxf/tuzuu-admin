package models

import (
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetRoute_distanceTable() (route_distanceTable models.Table) {

	route_distanceTable.Info.FieldList = []types.Field{{
			Head:     language.Get("route_distance-id"),
			Field:    "id",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-poi_set"),
			Field:    "poi_set",
			TypeName: "longtext",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-distance"),
			Field:    "distance",
			TypeName: "double",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-duration"),
			Field:    "duration",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-route_id"),
			Field:    "route_id",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-time"),
			Field:    "time",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	route_distanceTable.Info.Table = "route_distance"
	route_distanceTable.Info.Title = language.Get("info-route_distance")
	route_distanceTable.Info.Description = "Route_distance"

	route_distanceTable.Form.FormList = []types.Form{{
			Head:     language.Get("route_distance-id"),
			Field:    "id",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-poi_set"),
			Field:    "poi_set",
			TypeName: "longtext",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-distance"),
			Field:    "distance",
			TypeName: "double",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-duration"),
			Field:    "duration",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-route_id"),
			Field:    "route_id",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("route_distance-time"),
			Field:    "time",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	route_distanceTable.Form.Table = "route_distance"
	route_distanceTable.Form.Title = language.Get("form-route_distance")
	route_distanceTable.Form.Description = "Route_distance"

	route_distanceTable.ConnectionDriver = "mysql"

	return
}