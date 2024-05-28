package schema

import (
	"saas/pkg/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrderSales struct {
	ent.Schema
}

func (OrderSales) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("sales_id").Comment("销售id").Optional(),
		field.Int64("ratio").Comment("").Optional(),
	}
}

func (OrderSales) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (OrderSales) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("aufk", Order.Type).
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
