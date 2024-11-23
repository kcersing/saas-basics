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

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("商品名").Optional(),
		field.String("pic").Comment("主图").Optional(),
		field.String("description").Comment("详情").Optional(),
		field.Float("price").Comment("价格").Optional(),
		field.Int64("stock").Comment("库存").Optional(),
		field.Int64("create_id").Comment("创建人id").Optional(),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("propertys", ProductProperty.Type),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
