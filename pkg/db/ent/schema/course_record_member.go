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

type CourseRecordMember struct {
	ent.Schema
}

func (CourseRecordMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("course_record_schedule_id").Comment("课程ID").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.Time("start_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("end_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("sign_start_time").Default(time.Now).Comment("上课签到时间").Optional(),
		field.Time("sign_end_time").Default(time.Now).Comment("下课签到时间").Optional(),

		field.Int64("member_product_id").Comment("会员购买课ID").Optional(),
		field.Int64("member_product_item_id").Comment("会员购买课ID").Optional(),
		field.Int64("coach_id").Comment("教练ID").Optional(),
	}
}

func (CourseRecordMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (CourseRecordMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schedule", CourseRecordSchedule.Type).StructTag(`json:"sch"`).
			Ref("members").Unique().
			Field("course_record_schedule_id"),
	}
}

func (CourseRecordMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id").
			Unique(),
	}
}

func (CourseRecordMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "course_record_member"},
		entsql.WithComments(true),
	}
}
