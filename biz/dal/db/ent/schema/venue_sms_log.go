package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/biz/dal/db/ent/schema/mixins"
)

type VenueSmsLog struct {
	ent.Schema
}

func (VenueSmsLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.String("mobile").Comment("手机号"),
		field.String("biz_id").Comment("BizId"),
		field.String("code").Comment("验证码"),
		field.String("content").Comment("内容"),
		field.Int64("notify_type").Comment("通知类型[1会员;2员工]").Optional(),
		field.String("template").Comment("短信模板"),
	}
}

func (VenueSmsLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (VenueSmsLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("venue", Venue.Type).Ref("smslog").Field("venue_id").Unique(),
	}
}

func (VenueSmsLog) Indexes() []ent.Index {
	return []ent.Index{}
}

func (VenueSmsLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_venue_sms_log"},
		entsql.WithComments(true),
	}
}
