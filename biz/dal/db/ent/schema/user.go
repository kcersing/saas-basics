package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/biz/dal/db/ent/schema/mixins"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("mobile").Unique().Comment("mobile number | 手机号"),
		field.String("name").Optional().Comment("name | 姓名"),
		field.Int64("gender").Default(3).Comment("性别 | [0:女性;1:男性;3:保密]").Optional(),

		field.String("username").Unique().Comment("user's login name | 登录名"),
		field.String("password").Comment("password | 密码"),

		field.String("functions").Comment("functions | 职能"),
		field.Int64("type").Default(1).Comment("账号类别1普通 2管理员").Optional(),

		field.Int64("job_time").Default(1).Comment("job time | [1:全职;2:兼职;]").Optional(),

		field.Int64("role_id").Optional().Default(0).Comment("role id | 角色ID"),

		field.Int64("default_venue_id").Optional().Comment("登陆后默认场馆ID"),

		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("avatar | 头像路径"),

		field.String("detail").Comment("详情").Optional(),
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
		edge.To("token", Token.Type).Unique(),
		edge.To("tags", DictionaryDetail.Type),
		edge.To("created_orders", Order.Type),
		edge.To("user_entry", EntryLogs.Type),
		edge.To("venues", Venue.Type),
		//edge.To("user_face", Face.Type),

		edge.To("user_time_period", UserScheduling.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "mobile").
			Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_users", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
