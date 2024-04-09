package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
	"time"
)

type EntryLogs struct {
	ent.Schema
}

func (EntryLogs) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("user_id").Comment("用户id").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),

		field.Int64("member_product_id").Comment("用户产品id").Optional(),
		field.Int64("member_property_id").Comment("场馆id").Optional(),

		field.Time("entry_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("进场时间").
			Optional(),
		field.Time("leaving_time").
			Default(time.Now).
			UpdateDefault(time.Now).
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
	return nil
}

func (EntryLogs) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id", "member_id", "user_id"),
	}
}

func (EntryLogs) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "entry_logs"},
		entsql.WithComments(true),
	}
}
