package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type MemberProductPropertyVenue struct {
	ent.Schema
}

func (MemberProductPropertyVenue) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("member_product_property_id").Comment("").Optional(),
	}
}

func (MemberProductPropertyVenue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberProductPropertyVenue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", MemberProductProperty.Type).
			Ref("member_product_property_venues").
			Field("member_product_property_id").Unique(),
	}
}

func (MemberProductPropertyVenue) Indexes() []ent.Index {
	return []ent.Index{}
}

func (MemberProductPropertyVenue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_product_property_venue"},
		entsql.WithComments(true),
	}
}
