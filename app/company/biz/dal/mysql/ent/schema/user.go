package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().Comment("user's login name | 登录名"),
		field.String("password").Comment("password | 密码"),
		field.String("nickname").Optional().Comment("nickname | 昵称"),
		field.String("side_mode").Optional().Default("dark").Comment("template mode | 布局方式"),
		field.String("base_color").Optional().Default("#fff").Comment("base color of template | 后台页面色调"),
		field.String("active_color").Optional().Default("#1890ff").Comment("active color of template | 当前激活的颜色设定"),
		field.Int64("role_id").Optional().Default(2).Comment("role id | 角色ID"),
		field.String("mobile").Unique().Comment("mobile number | 手机号"),
		field.String("email").Optional().Comment("email | 邮箱号"),
		field.String("wecom").Optional().Comment("wecom | 微信号"),

		field.String("job").Optional().Comment("职业"),
		field.String("organization").Optional().Comment("部门"),

		field.Int64("default_venue_id").Optional().Comment("登陆后默认场馆ID"),

		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("avatar | 头像路径"),

		field.Int64("gender").Default(3).Comment("性别 | [0:女性;1:男性;3:保密]").Optional(),
		field.Time("birthday").Comment("出生日期").Optional(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("token", Token.Type).Unique(),
		//edge.To("created_orders", Order.Type),
		edge.To("user_entry", EntryLogs.Type),
		edge.To("user_face", Face.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "u_users", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
