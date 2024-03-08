package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type ProductProperty struct {
	ent.Schema
}

func (ProductProperty) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("类型").Optional(),
		field.String("spu_name").Comment("spu名").Optional(),
		field.Int64("duration").Comment("总时长").Optional(),
		field.Int64("length").Comment("单次时长").Optional(),
		field.Int64("count").Comment("次数").Optional(),
		field.Float("spu_price").Comment("定价").Optional(),
	}
}

func (ProductProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (ProductProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("propertys"),
	}
}

func (ProductProperty) Indexes() []ent.Index {
	return []ent.Index{}
}

func (ProductProperty) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_property"},
		entsql.WithComments(true),
	}
}
