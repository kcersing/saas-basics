package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type ProductPropertyVenue struct {
	ent.Schema
}

func (ProductPropertyVenue) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("product_property_id").Comment("").Optional(),
	}
}

func (ProductPropertyVenue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (ProductPropertyVenue) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (ProductPropertyVenue) Indexes() []ent.Index {
	return []ent.Index{}
}

func (ProductPropertyVenue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_property_venue"},
		entsql.WithComments(true),
	}
}
