package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
	}
}

func (Venue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Venue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("places", VenuePlace.Type),
		edge.To("venue_entry", EntryLogs.Type),
		edge.From("member_property_venues", MemberProductProperty.Type).Ref("venues"),
		edge.From("property_venues", ProductProperty.Type).Ref("venues"),
	}
}

func (Venue) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Venue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "venue", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
