package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type Venue struct {
	ent.Schema
}

func (Venue) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("address").Comment("地址 省/市/区").Optional(),
		field.String("address_detail").Comment("详细地址").Optional(),
		field.String("latitude").Comment("维度").Optional(),
		field.String("longitude").Comment("经度").Optional(),
		field.String("mobile").Comment("联系电话").Optional(),
		field.String("pic").Comment("场馆照片").Optional(),
		field.String("information").Comment("详情").Optional(),
		field.Int64("status").Default(0).Comment("状态").Optional(),
	}
}

func (Venue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Venue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("places", VenuePlace.Type),
	}
}

func (Venue) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Venue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "venue"},
		entsql.WithComments(true),
	}
}
