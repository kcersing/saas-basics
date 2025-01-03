package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/biz/dal/db/ent/schema/mixins"
	"saas/idl_gen/model/base"
)

type VenuePlace struct {
	ent.Schema
}

func (VenuePlace) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.Int64("classify").Comment("分类").Optional(),
		field.String("pic").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Comment("pic | 照片"),
		field.Int64("venue_id").Comment("场馆id").Optional(),
		field.Int64("number").Comment("可容纳人数").Optional(),

		field.Int64("is_show").Default(1).Comment("是否展示:1展示;2不展示").Optional(),
		field.Int64("is_accessible").Default(1).Comment("是否展示;1开放;2关闭").Optional(),
		field.Int64("is_booking").Default(1).Comment("是否预约:1可预约;2不可").Optional(),
		field.String("information").Comment("详情").Optional(),

		field.JSON("seat", []*base.Seat{}).Default([]*base.Seat{}).Comment("座位").Optional(),
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
		edge.From("venue", Venue.Type).Ref("places").Field("venue_id").Unique(),

		edge.To("products", Product.Type),
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
