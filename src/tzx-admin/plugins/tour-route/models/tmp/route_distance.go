package model

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetRoute_distanceTable() (route_distanceTable models.Table) {

	route_distanceTable.Info.FieldList = []types.Field{{
			Head:     "Id",
			Field:    "id",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Poi_set",
			Field:    "poi_set",
			TypeName: "longtext",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Distance",
			Field:    "distance",
			TypeName: "double",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Duration",
			Field:    "duration",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Route_id",
			Field:    "route_id",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Time",
			Field:    "time",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	route_distanceTable.Info.Table = "route_distance"
	route_distanceTable.Info.Title = "Route_distance"
	route_distanceTable.Info.Description = "Route_distance"

	route_distanceTable.Form.FormList = []types.Form{{
			Head:     "Id",
			Field:    "id",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Poi_set",
			Field:    "poi_set",
			TypeName: "longtext",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Distance",
			Field:    "distance",
			TypeName: "double",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Duration",
			Field:    "duration",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Route_id",
			Field:    "route_id",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Time",
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
	route_distanceTable.Form.Title = "Route_distance"
	route_distanceTable.Form.Description = "Route_distance"

	route_distanceTable.ConnectionDriver = "mysql"

	return
}