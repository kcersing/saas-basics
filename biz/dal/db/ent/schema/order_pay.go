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

type OrderPay struct {
	ent.Schema
}

func (OrderPay) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Float("remission").Comment("减免").Optional(),
		field.Float("pay").Comment("实际付款").Optional(),
		field.String("note").Comment("备注").Optional(),
		field.Time("pay_at").Comment("支付时间").Optional(),
		field.String("pay_way").Comment("支付方式").Optional(),
		field.String("pay_sn").Comment("支付单号").Optional(),
		field.String("prepay_id").Comment("预支付交易会话标识").Optional(),
		field.JSON("pay_extra", []string{}).Comment("支付额外信息").Optional(),
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
