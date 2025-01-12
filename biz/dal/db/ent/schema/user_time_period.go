package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/biz/dal/db/ent/schema/mixins"
	"saas/idl_gen/model/base"
)

type UserTimePeriod struct {
	ent.Schema
}

func (UserTimePeriod) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").Comment("日期").Optional(),
		field.JSON("period", base.Period{}).Comment("时间段").Optional(),
		field.Int64("user_id").Comment("員工id").Optional(),
		field.Int64("venue_id").Comment("id").Optional(),
	}
}

func (UserTimePeriod) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (UserTimePeriod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("user_time_period").Field("user_id").Unique(),
	}
}

func (UserTimePeriod) Indexes() []ent.Index {
	return []ent.Index{}
}

func (UserTimePeriod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_user_time_period", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
