package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/biz/dal/db/ent/schema/mixins"
)

type Contract struct {
	ent.Schema
}

func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Comment("name | 名称"),
		field.String("content").Optional().Comment("content | 内容"),
	}
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Contract) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).Ref("contracts"),
	}
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Contract) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "contracts", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
