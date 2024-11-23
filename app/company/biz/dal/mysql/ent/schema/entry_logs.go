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

type EntryLogs struct {
	ent.Schema
}

func (EntryLogs) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Comment("场馆id").Optional(),

		field.Int64("user_id").Default(0).Comment("用户id").Optional(),

		field.Int64("member_id").Default(0).Comment("会员id").Optional(),
		field.Int64("member_product_id").Comment("用户产品id").Optional(),
		field.Int64("member_property_id").Comment("属性id").Optional(),

		field.Time("entry_time").
			Comment("进场时间").
			Optional(),
		field.Time("leaving_time").
			Comment("离场时间").
			Optional(),
	}
}

func (EntryLogs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (EntryLogs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("venues", Venue.Type).Ref("venue_entry").Field("venue_id").Unique(),
		edge.From("members", Member.Type).Ref("member_entry").Field("member_id").Unique(),
		edge.From("users", User.Type).Ref("user_entry").Field("user_id").Unique(),
		edge.From("member_products", MemberProduct.Type).Ref("member_product_entry").Field("member_product_id").Unique(),
	}
}

func (EntryLogs) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id", "member_id", "user_id"),
	}
}

func (EntryLogs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "entry_logs", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
