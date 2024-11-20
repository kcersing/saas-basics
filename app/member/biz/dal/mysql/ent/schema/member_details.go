package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MemberDetails struct {
	ent.Schema
}

func (MemberDetails) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("member_id").Comment("会员id").Optional(),
		field.String("email").Optional().Comment("email | 邮箱号"),
		field.String("wecom").Optional().Comment("wecom | 微信号"),
		field.Int64("gender").Default(3).Comment("性别 | [0:女性;1:男性;3:保密]").Optional(),
		field.Time("birthday").Comment("出生日期").Optional(),
	}
}

func (MemberDetails) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("info", Member.Type).
			Ref("member_details").
			Field("member_id").Unique(),
	}
}

func (MemberDetails) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id"),
	}
}

func (MemberDetails) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_details", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
