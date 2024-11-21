package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent/schema/index"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ScheduleCoach struct {
	ent.Schema
}

func (ScheduleCoach) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("coach_id").Comment("教练ID").Optional(),
		field.Int64("schedule_id").Comment("课程ID").Optional(),
		field.String("schedule_name").Comment("课程名称").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.Time("start_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("end_time").Default(time.Now).Comment("结束时间").Optional(),
		field.Time("sign_start_time").Default(time.Now).Comment("上课签到时间").Optional(),
		field.Time("sign_end_time").Default(time.Now).Comment("下课签到时间").Optional(),

		field.String("coach_name").Comment("教练名称").Optional(),
	}
}

func (ScheduleCoach) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (ScheduleCoach) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schedule", Schedule.Type).
			Ref("coachs").Unique().
			Field("schedule_id"),
	}
}

func (ScheduleCoach) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id"),
		index.Fields("coach_id"),
		index.Fields("schedule_id"),
	}
}

func (ScheduleCoach) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "s_schedule_coach", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
