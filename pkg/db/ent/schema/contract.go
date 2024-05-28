package schema

import (
	"saas/pkg/db/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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
	return nil
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Contract) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_contracts"},
		entsql.WithComments(true),
	}
}
