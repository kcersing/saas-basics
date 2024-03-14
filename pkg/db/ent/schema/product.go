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

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("sn").Comment("编号").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("create_id").Comment("创建人id").Optional(),
		field.String("name").Comment("商品名").Optional(),
		field.Int64("pic").Comment("主图").Optional(),
		field.String("description").Comment("详情").Optional(),
		field.String("price").Comment("价格").Optional(),
		field.Int64("stock").Comment("库存").Optional(),
		field.Int64("status").Default(0).Comment("状态").Optional(),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("propertys", ProductProperty.Type),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "sn").
			Unique(),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product"},
		entsql.WithComments(true),
	}
}