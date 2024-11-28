package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
)

type VenuePlace struct {
	ent.Schema
}

func (VenuePlace) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),

		field.String("pic").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("pic | 照片"),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("number").Comment("可容纳人数").Optional(),
		field.String("information").Comment("详情").Optional(),
	}
}

func (VenuePlace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (VenuePlace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("venue", Venue.Type).
			Ref("places").Field("venue_id").Unique(),
	}
}

func (VenuePlace) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id"),
	}
}

func (VenuePlace) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "venue_place", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
