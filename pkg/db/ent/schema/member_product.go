package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type MemberProduct struct {
	ent.Schema
}

func (MemberProduct) Fields() []ent.Field {

	return []ent.Field{
		field.String("sn").Comment("编号").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("product_id").Comment("产品ID").Optional(),
		field.Float("name").Comment("产品名称").Optional(),
		field.Float("price").Comment("产品价格").Optional(),
		field.Time("validity_at").Comment("生效时间").Optional(),
		field.Time("cancel_at").Comment("作废时间").Optional(),
		field.Int64("status").Default(0).Comment("状态").Optional(),
	}
}

func (MemberProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Member.Type).
			Ref("member_products").Unique().
			Field("member_id").Unique(),

		edge.To("member_product_propertys", MemberProductProperty.Type),
	}
}

func (MemberProduct) Indexes() []ent.Index {
	return []ent.Index{}
}

func (MemberProduct) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_product"},
		entsql.WithComments(true),
	}
}
