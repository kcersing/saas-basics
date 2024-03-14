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

type OrderPay struct {
	ent.Schema
}

func (OrderPay) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.String("pay_sn").Comment("支付编号").Optional(),
		field.Float("remission").Comment("减免").Optional(),
		field.Float("pay").Comment("实际付款").Optional(),
		field.Int64("create_id").Comment("操作人id").Optional(),
	}
}

func (OrderPay) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderPay) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Order.Type).
			Ref("pay").Unique().
			Field("order_id"),
	}
}

func (OrderPay) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id").
			Unique(),
	}
}

func (OrderPay) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_pay"},
		entsql.WithComments(true),
	}
}