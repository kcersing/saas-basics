package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type OrderAmount struct {
	ent.Schema
}

func (OrderAmount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").NonNegative().Comment("订单id").Optional(),
		field.Float("total").Comment("总金额").Optional(),
		field.Float("remission").Comment("减免").Optional(),
		field.Float("pay").Comment("实际付款").Optional(),
	}
}

func (OrderAmount) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderAmount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("amount").
			Field("order_id").Unique(),
	}
}
func (OrderAmount) Indexes() []ent.Index {
	return []ent.Index{}
}

func (OrderAmount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_amount"},
		entsql.WithComments(true),
	}
}
