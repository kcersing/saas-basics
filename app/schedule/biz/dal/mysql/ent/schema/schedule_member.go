package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type ScheduleMember struct {
	ent.Schema
}

func (ScheduleMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("schedule_id").Comment("课程ID").Optional(),
		field.String("schedule_name").Comment("课程名称").Optional(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("member_product_id").Comment("会员购买课ID").Optional(),
		field.Int64("member_product_property_id").Comment("会员购买课ID").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.Time("start_time").Default(time.Now).Comment("开始时间").Optional(),
		field.Time("end_time").Default(time.Now).Comment("结束时间").Optional(),
		field.Time("sign_start_time").Default(time.Now).Comment("上课签到时间").Optional(),
		field.Time("sign_end_time").Default(time.Now).Comment("下课签到时间").Optional(),

		field.String("member_name").Comment("会员名称").Optional(),
		field.String("member_product_name").Comment("会员产品名称").Optional(),
		field.String("member_product_property_name").Comment("会员产品属性名称").Optional(),

		field.String("remark").Comment("备注").Optional(),
	}
}

func (ScheduleMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (ScheduleMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("schedule", Schedule.Type).
			Ref("members").Unique().
			Field("schedule_id"),
	}
}

func (ScheduleMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id"),
		index.Fields("member_id"),
		index.Fields("schedule_id"),
	}
}

func (ScheduleMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "s_schedule_member", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
