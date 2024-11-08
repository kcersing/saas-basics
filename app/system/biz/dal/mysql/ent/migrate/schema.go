// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysApisColumns holds the columns for the "sys_apis" table.
	SysApisColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "path", Type: field.TypeString, Comment: "API path | API 路径"},
		{Name: "title", Type: field.TypeString, Comment: "API title | API 名称"},
		{Name: "description", Type: field.TypeString, Comment: "API description | API 描述"},
		{Name: "api_group", Type: field.TypeString, Comment: "API group | API 分组"},
		{Name: "method", Type: field.TypeString, Comment: "HTTP method | HTTP 请求类型", Default: "POST"},
	}
	// SysApisTable holds the schema information for the "sys_apis" table.
	SysApisTable = &schema.Table{
		Name:       "sys_apis",
		Columns:    SysApisColumns,
		PrimaryKey: []*schema.Column{SysApisColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "api_path_method",
				Unique:  true,
				Columns: []*schema.Column{SysApisColumns[3], SysApisColumns[7]},
			},
		},
	}
	// SysDictionariesColumns holds the columns for the "sys_dictionaries" table.
	SysDictionariesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "status", Type: field.TypeInt64, Nullable: true, Comment: "状态[0:禁用;1:正常]", Default: 1},
		{Name: "title", Type: field.TypeString, Comment: "the title shown in the ui | 展示名称 （建议配合i18n）"},
		{Name: "name", Type: field.TypeString, Unique: true, Comment: "the name of dictionary for search | 字典搜索名称"},
		{Name: "description", Type: field.TypeString, Comment: "the description of dictionary | 字典描述"},
	}
	// SysDictionariesTable holds the schema information for the "sys_dictionaries" table.
	SysDictionariesTable = &schema.Table{
		Name:       "sys_dictionaries",
		Columns:    SysDictionariesColumns,
		PrimaryKey: []*schema.Column{SysDictionariesColumns[0]},
	}
	// SysDictionaryDetailsColumns holds the columns for the "sys_dictionary_details" table.
	SysDictionaryDetailsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "status", Type: field.TypeInt64, Nullable: true, Comment: "状态[0:禁用;1:正常]", Default: 1},
		{Name: "title", Type: field.TypeString, Comment: "the title shown in the ui | 展示名称 （建议配合i18n）"},
		{Name: "key", Type: field.TypeString, Comment: "key | 键"},
		{Name: "value", Type: field.TypeString, Comment: "value | 值"},
		{Name: "dictionary_id", Type: field.TypeInt64, Nullable: true, Comment: "Dictionary ID | 字典ID"},
	}
	// SysDictionaryDetailsTable holds the schema information for the "sys_dictionary_details" table.
	SysDictionaryDetailsTable = &schema.Table{
		Name:       "sys_dictionary_details",
		Columns:    SysDictionaryDetailsColumns,
		PrimaryKey: []*schema.Column{SysDictionaryDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_dictionary_details_sys_dictionaries_dictionary_details",
				Columns:    []*schema.Column{SysDictionaryDetailsColumns[7]},
				RefColumns: []*schema.Column{SysDictionariesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "dictionarydetail_key_dictionary_id",
				Unique:  true,
				Columns: []*schema.Column{SysDictionaryDetailsColumns[5], SysDictionaryDetailsColumns[7]},
			},
		},
	}
	// SysLogsColumns holds the columns for the "sys_logs" table.
	SysLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "type", Type: field.TypeString, Comment: "type of log | 日志类型"},
		{Name: "method", Type: field.TypeString, Comment: "method of log | 日志请求方法"},
		{Name: "api", Type: field.TypeString, Comment: "api of log | 日志请求api"},
		{Name: "success", Type: field.TypeBool, Comment: "success of log | 日志请求是否成功"},
		{Name: "req_content", Type: field.TypeString, Nullable: true, Size: 2147483647, Comment: "content of request log | 日志请求内容"},
		{Name: "resp_content", Type: field.TypeString, Nullable: true, Size: 2147483647, Comment: "content of response log | 日志返回内容"},
		{Name: "ip", Type: field.TypeString, Nullable: true, Comment: "ip of log | 日志IP"},
		{Name: "user_agent", Type: field.TypeString, Nullable: true, Comment: "user_agent of log | 日志用户客户端"},
		{Name: "operator", Type: field.TypeString, Nullable: true, Comment: "operator of log | 日志操作者"},
		{Name: "time", Type: field.TypeInt64, Nullable: true, Comment: "time of log(millisecond) | 日志时间(毫秒)"},
	}
	// SysLogsTable holds the schema information for the "sys_logs" table.
	SysLogsTable = &schema.Table{
		Name:       "sys_logs",
		Columns:    SysLogsColumns,
		PrimaryKey: []*schema.Column{SysLogsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "logs_api",
				Unique:  false,
				Columns: []*schema.Column{SysLogsColumns[5]},
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "path", Type: field.TypeString, Nullable: true, Comment: "index path | 菜单路由路径", Default: ""},
		{Name: "name", Type: field.TypeString, Comment: "index name | 菜单名称"},
		{Name: "order_no", Type: field.TypeInt64, Comment: "sorting numbers | 排序编号", Default: 0},
		{Name: "disabled", Type: field.TypeInt64, Nullable: true, Comment: "disable status | 是否停用", Default: 0},
		{Name: "ignore", Type: field.TypeBool, Nullable: true, Comment: "当前路由是否渲染菜单项，为 true 的话不会在菜单中显示，但可通过路由地址访问", Default: false},
		{Name: "parent_id", Type: field.TypeInt64, Nullable: true, Comment: "parent menu ID | 父菜单ID"},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[8]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenuParamsColumns holds the columns for the "sys_menu_params" table.
	SysMenuParamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "type", Type: field.TypeString, Comment: "pass parameters via params or query | 参数类型"},
		{Name: "key", Type: field.TypeString, Comment: "the key of parameters | 参数键"},
		{Name: "value", Type: field.TypeString, Comment: "the value of parameters | 参数值"},
		{Name: "menu_params", Type: field.TypeInt64, Nullable: true},
	}
	// SysMenuParamsTable holds the schema information for the "sys_menu_params" table.
	SysMenuParamsTable = &schema.Table{
		Name:       "sys_menu_params",
		Columns:    SysMenuParamsColumns,
		PrimaryKey: []*schema.Column{SysMenuParamsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menu_params_sys_menus_params",
				Columns:    []*schema.Column{SysMenuParamsColumns[6]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "type", Type: field.TypeString, Comment: "类型[1:用户user;2:会员member]"},
		{Name: "to_user_id", Type: field.TypeString, Comment: "该消息接受者ID"},
		{Name: "from_user_id", Type: field.TypeString, Comment: "该消息发送者ID"},
		{Name: "content", Type: field.TypeString, Comment: "消息内容"},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "messages_to_user_id_from_user_id",
				Unique:  false,
				Columns: []*schema.Column{MessagesColumns[4], MessagesColumns[5]},
			},
		},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, Comment: "primary key"},
		{Name: "created_at", Type: field.TypeTime, Comment: "created time"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "last update time"},
		{Name: "status", Type: field.TypeInt64, Nullable: true, Comment: "状态[0:禁用;1:正常]", Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "role name | 角色名"},
		{Name: "value", Type: field.TypeString, Unique: true, Comment: "role value for permission control in front end | 角色值，用于前端权限控制"},
		{Name: "default_router", Type: field.TypeString, Comment: "default menu : dashboard | 默认登录页面", Default: "dashboard"},
		{Name: "remark", Type: field.TypeString, Comment: "remark | 备注", Default: ""},
		{Name: "order_no", Type: field.TypeInt64, Comment: "order number | 排序编号", Default: 0},
		{Name: "apis", Type: field.TypeJSON, Comment: "apis"},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
	}
	// RoleMenusColumns holds the columns for the "role_menus" table.
	RoleMenusColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt64},
		{Name: "menu_id", Type: field.TypeInt64},
	}
	// RoleMenusTable holds the schema information for the "role_menus" table.
	RoleMenusTable = &schema.Table{
		Name:       "role_menus",
		Columns:    RoleMenusColumns,
		PrimaryKey: []*schema.Column{RoleMenusColumns[0], RoleMenusColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_menus_role_id",
				Columns:    []*schema.Column{RoleMenusColumns[0]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_menus_menu_id",
				Columns:    []*schema.Column{RoleMenusColumns[1]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysApisTable,
		SysDictionariesTable,
		SysDictionaryDetailsTable,
		SysLogsTable,
		SysMenusTable,
		SysMenuParamsTable,
		MessagesTable,
		SysRolesTable,
		RoleMenusTable,
	}
)

func init() {
	SysApisTable.Annotation = &entsql.Annotation{
		Table: "sys_apis",
	}
	SysDictionariesTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionaries",
	}
	SysDictionaryDetailsTable.ForeignKeys[0].RefTable = SysDictionariesTable
	SysDictionaryDetailsTable.Annotation = &entsql.Annotation{
		Table: "sys_dictionary_details",
	}
	SysLogsTable.Annotation = &entsql.Annotation{
		Table: "sys_logs",
	}
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_menus",
	}
	SysMenuParamsTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenuParamsTable.Annotation = &entsql.Annotation{
		Table: "sys_menu_params",
	}
	MessagesTable.Annotation = &entsql.Annotation{
		Table: "messages",
	}
	SysRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_roles",
	}
	RoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	RoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
}
