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

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("order_id").Comment("订单id").Optional(),
		field.Int64("product_id").Comment("产品id").Optional(),
		field.Int64("related_user_product_id").Default(0).Comment("关联会员产品id").Optional(),
		field.Int64("data").Default(0).Comment("数据附件").Optional(),
	}
}

func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("aufk", Order.Type).
			Ref("item").Unique().
			Field("order_id"),
	}
}

func (OrderItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id").
			Unique(),
	}
}

func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_item"},
		entsql.WithComments(true),
	}
}
