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

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_sn").Comment("订单编号").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("status").Default(0).Comment("状态").Optional(),
		field.String("source").Default("").Comment("订单来源").Optional(),
		field.String("device").Default("").Comment("设备来源").Optional(),
		field.Time("completion_at").Comment("订单完成时间").Optional(),
		field.Int64("create_id").Comment("创建人id").Optional(),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amount", OrderAmount.Type).StructTag(`json:"amounts"`),
		edge.To("item", OrderItem.Type),
		edge.To("pay", OrderPay.Type),
		edge.To("sales", OrderSales.Type),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id", "order_sn").
			Unique(),
	}
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order"},
		entsql.WithComments(true),
	}
}
