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

type Contest struct {
	ent.Schema
}

func (Contest) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").Comment("比赛名称").Optional(),
		field.Time("sign_number").Comment("报名人数").Optional(),
		field.Time("sign_start_at").Comment("报名开始时间").Optional(),
		field.Time("sign_end_at").Comment("报名结束时间").Optional(),
		field.Time("number").Comment("参赛人数").Optional(),
		field.Time("start_at").Comment("比赛开始时间").Optional(),
		field.Time("end_at").Comment("比赛结束时间").Optional(),
		field.String("pic").Comment("比赛图片").Optional(),
		field.String("sponsor").Comment("主办方").Optional(),
		field.Float("fee").Comment("费用").Optional(),
		field.Int64("is_cancel").Comment("是否支持取消报名 0支持 1不支持").Default(0).Optional(),
		field.Int64("cancel_time").Comment("取消时间").Default(0).Optional(),
		field.String("detail").Comment("详情").Optional(),
		field.String("sign_fields").Comment("报名信息").Optional(),
	}
}

func (Contest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Contest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participant", ContestParticipant.Type),
	}
}

func (Contest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "sign_start_at", "sign_end_at", "start_at", "end_at").
			Unique(),
	}
}

func (Contest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "contest"},
		entsql.WithComments(true),
	}
}
