package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/biz/dal/db/ent/schema/mixins"
)

type MemberProfile struct {
	ent.Schema
}

func (MemberProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("mobile_ascription").Default(0).Optional().Comment("手机号归属"),
		field.String("father_name").Comment("父亲名称").Optional(),
		field.String("mother_name").Comment("母亲名称").Optional(),
		field.Int64("grade").Default(0).Comment("年级").Optional(),
		field.Int64("intention").Default(0).Comment("意向").Optional(),
		field.Int64("source").Default(0).Comment("来源").Optional(),
	}
}

func (MemberProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Member.Type).
			Ref("member_profile").
			Field("member_id").
			Unique(),
	}
}

func (MemberProfile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id"),
	}
}

func (MemberProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_profile", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
