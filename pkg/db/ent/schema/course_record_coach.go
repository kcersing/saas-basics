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

type CourseRecordCoach struct {
	ent.Schema
}

func (CourseRecordCoach) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("coach_id").Comment("教练ID").Optional(),
		field.Int64("course_record_schedule_id").Comment("课程ID").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.Time("start_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("end_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("sign_start_time").Default(time.Now).Comment("上课签到时间").Optional(),
		field.Time("sign_end_time").Default(time.Now).Comment("下课签到时间").Optional(),
	}
}

func (CourseRecordCoach) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (CourseRecordCoach) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schedule", CourseRecordSchedule.Type).
			Ref("coachs").Unique().
			Field("course_record_schedule_id"),
	}
}

func (CourseRecordCoach) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("coach_id").
			Unique(),
	}
}

func (CourseRecordCoach) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "course_record_coach"},
		entsql.WithComments(true),
	}
}
