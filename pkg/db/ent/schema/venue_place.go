package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type VenuePlace struct {
	ent.Schema
}

func (VenuePlace) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("pic").Comment("照片").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("status").Default(0).Comment("状态").Optional(),
	}
}

func (VenuePlace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (VenuePlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("venue", Venue.Type).
			Ref("places").Field("venue_id").Unique(),
	}
}

func (VenuePlace) Indexes() []ent.Index {
	return []ent.Index{}
}

func (VenuePlace) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "venue_place"},
		entsql.WithComments(true),
	}
}
