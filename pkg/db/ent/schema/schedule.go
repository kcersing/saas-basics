package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
)

type Schedule struct {
	ent.Schema
}

func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("类型").Optional(),
		field.String("name").Comment("名称").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("property_id").Comment("课程").Optional(),
		field.Int64("place_id").Comment("场地ID").Optional(),
		field.Int64("num").Comment("上课人数").Optional(),

		field.String("date").Comment("日期").Optional(),
		field.Time("start_time").Comment("开始时间").Optional(),
		field.Time("end_time").Comment("开始时间").Optional(),
		field.Float("price").Default(0).Comment("课程价格").Optional(),
		field.String("remark").Comment("备注").Optional(),
	}
}

func (Schedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", ScheduleMember.Type),
		edge.To("coachs", ScheduleCoach.Type),
	}
}

func (Schedule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id"),
		index.Fields("property_id"),
		index.Fields("start_time"),
		index.Fields("end_time"),
	}
}

func (Schedule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "schedule"},
		entsql.WithComments(true),
	}
}
