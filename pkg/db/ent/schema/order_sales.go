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

type OrderSales struct {
	ent.Schema
}

func (OrderSales) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Int64("sales_id").Comment("销售id").Optional(),
	}
}

func (OrderSales) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderSales) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Order.Type).
			Ref("sales").Unique().
			Field("order_id"),
	}
}

func (OrderSales) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id").
			Unique(),
	}
}

func (OrderSales) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_sales"},
		entsql.WithComments(true),
	}
}