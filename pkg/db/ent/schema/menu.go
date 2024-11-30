package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type Menu struct {
	ent.Schema
}

func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("parent_id").Optional().Comment("parent menu ID | 父菜单ID"),
		field.String("path").Optional().Default("").Comment("index path | 菜单路由路径"),
		field.String("name").Comment("index name | 菜单名称"),
		field.Int64("sort").Default(0).Comment("sort | 排序编号"),
		field.Int64("disabled").Optional().Default(0).Comment("disable status | 是否停用"),
		field.Bool("ignore").Optional().Default(false).Comment("当前路由是否渲染菜单项，为 true 的话不会在菜单中显示，但可通过路由地址访问"),

		field.Int64("level").Comment("menu level | 菜单层级"),
		field.Int64("menu_type").Comment("menu type | 菜单类型 0 目录 1 菜单 2 按钮"),
		field.String("redirect").Optional().Default("").Comment("redirect path | 跳转路径 （外链）"),
		field.String("component").Optional().Default("").Comment("the path of vue file | 组件路径"),

		field.String("url").Optional().Default("").Comment("url | "),

		field.Bool("hidden").Optional().Default(true).Comment("hidden | 是否隐藏菜单"),

		//meta
		field.String("title").Comment("menu name | 菜单显示标题"),
		field.String("icon").Comment("menu icon | 菜单图标"),
		field.String("active_menu").Optional().Default("").Comment("set the active menu | 激活菜单"),
		field.Bool("affix").Optional().Default(false).Comment("affix tab | Tab 固定"),
		field.Bool("no_cache").Optional().Default(false).Comment("do not keep alive the tab | 取消页面缓存"),
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("menus"),
		edge.To("children", Menu.Type).From("parent").Unique().Field("parent_id"),
		edge.To("params", MenuParam.Type),
	}
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_menus"},
		entsql.WithComments(true),
	}
}
