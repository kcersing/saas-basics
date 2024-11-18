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

type OrderPay struct {
	ent.Schema
}

func (OrderPay) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Float("remission").Comment("减免").Optional(),
		field.Float("pay").Comment("实际付款").Optional(),
		field.String("note").Comment("备注").Optional(),
		field.String("pay_way").Comment("支付方式").Optional(),
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
		edge.From("order", Order.Type).
			Ref("pay").
			Field("order_id").Unique(),
	}
}

func (OrderPay) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}

func (OrderPay) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_pay", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
