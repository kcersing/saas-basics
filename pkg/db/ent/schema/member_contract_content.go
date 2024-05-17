package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type MemberContractContent struct {
	ent.Schema
}

func (MemberContractContent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_contract_id").Comment("合同ID").Optional(),
		field.String("content").Optional().Comment("content | 内容"),
	}
}

func (MemberContractContent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberContractContent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contract", MemberContract.Type).Ref("content").Field("member_contract_id").Unique(),
	}
}

func (MemberContractContent) Indexes() []ent.Index {
	return []ent.Index{}
}

func (MemberContractContent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_contract_content"},
		entsql.WithComments(true),
	}
}
