package models

import (
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetTaskTable() (taskTable models.Table) {

	taskTable.Info.FieldList = []types.Field{{
			Head:     language.Get("task-id"),
			Field:    "id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-poi_id"),
			Field:    "poi_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-type"),
			Field:    "type",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-interaction"),
			Field:    "interaction",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-task_copy"),
			Field:    "task_copy",
			TypeName: "text",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-button_copy"),
			Field:    "button_copy",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-sort"),
			Field:    "sort",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-time"),
			Field:    "time",
			TypeName: "bigint",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-img_arr"),
			Field:    "img_arr",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-status"),
			Field:    "status",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-class"),
			Field:    "class",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-opt_arr"),
			Field:    "opt_arr",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	taskTable.Info.Table = "task"
	taskTable.Info.Title = language.Get("info-task")
	taskTable.Info.Description = "Task"

	taskTable.Form.FormList = []types.Form{{
			Head:     language.Get("task-id"),
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-poi_id"),
			Field:    "poi_id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-type"),
			Field:    "type",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-interaction"),
			Field:    "interaction",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-task_copy"),
			Field:    "task_copy",
			TypeName: "text",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-button_copy"),
			Field:    "button_copy",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-sort"),
			Field:    "sort",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-time"),
			Field:    "time",
			TypeName: "bigint",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-img_arr"),
			Field:    "img_arr",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-status"),
			Field:    "status",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-class"),
			Field:    "class",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     language.Get("task-opt_arr"),
			Field:    "opt_arr",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	taskTable.Form.Table = "task"
	taskTable.Form.Title = language.Get("form-task")
	taskTable.Form.Description = "Task"

	taskTable.ConnectionDriver = "mysql"

	return
}