package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
	return []ent.Edge{}
}

func (UserScheduling) Indexes() []ent.Index {
	return []ent.Index{}
}

func (UserScheduling) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user_scheduling", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
