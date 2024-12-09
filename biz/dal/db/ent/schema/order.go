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

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_sn").Comment("订单编号").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("member_id").Comment("会员id").Optional(),
		field.String("nature").Comment("业务类型").Optional(),
		field.String("product_type").Comment("产品类型").Optional(),
		field.Int64("status").Default(1).Comment("状态 | [1:待付款]").Optional(),
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
		edge.To("amount", OrderAmount.Type),
		edge.To("item", OrderItem.Type),
		edge.To("pay", OrderPay.Type),
		edge.To("order_contents", MemberContract.Type),
		edge.To("sales", OrderSales.Type),

		edge.From("order_venues", Venue.Type).Ref("venue_orders").Field("venue_id").Unique(),
		edge.From("order_members", Member.Type).Ref("member_orders").Field("member_id").Unique(),
		edge.From("order_creates", User.Type).Ref("created_orders").Field("create_id").Unique(),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_sn"),
		index.Fields("venue_id"),
		index.Fields("member_id"),
		index.Fields("completion_at"),
	}
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
