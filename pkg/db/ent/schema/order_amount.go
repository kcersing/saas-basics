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

type OrderAmount struct {
	ent.Schema
}

func (OrderAmount) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
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
		edge.From("aufk", Order.Type).
			Ref("amount").Unique().
			Field("order_id"),
	}
}
func (OrderAmount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id").
			Unique(),
	}
}

func (OrderAmount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_amount"},
		entsql.WithComments(true),
	}
}
