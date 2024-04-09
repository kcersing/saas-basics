package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
	"time"
)

type CourseRecordSchedule struct {
	ent.Schema
}

func (CourseRecordSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("类型").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("place_id").Comment("场地ID").Optional(),
		field.Int64("num").Comment("上课人数").Optional(),
		field.Time("start_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("end_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Float("price").Default(0).Comment("课程价格").Optional(),
	}
}

func (CourseRecordSchedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (CourseRecordSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", CourseRecordMember.Type),
		edge.To("coachs", CourseRecordCoach.Type),
	}
}

func (CourseRecordSchedule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id").
			Unique(),
	}
}

func (CourseRecordSchedule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "course_record_schedule"},
		entsql.WithComments(true),
	}
}
