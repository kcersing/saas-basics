package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/biz/dal/db/ent/schema/mixins"
)

type UserScheduling struct {
	ent.Schema
}

func (UserScheduling) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("user_id", []string{}).Default(0).Comment("会员ID").Optional(),
		field.String("date").Optional().Comment("日期"),
	}
}

func (UserScheduling) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (UserScheduling) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("token", Token.Type).Unique(),
		edge.To("tags", DictionaryDetail.Type),
		edge.To("created_orders", Order.Type),
		edge.To("user_entry", EntryLogs.Type),
		edge.To("venues", Venue.Type),
		//edge.To("user_face", Face.Type),
	}
}

func (UserScheduling) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "mobile").
			Unique(),
	}
}

func (UserScheduling) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user_scheduling", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
