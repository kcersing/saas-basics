package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/biz/dal/db/ent/schema/mixins"
)

type VenueSms struct {
	ent.Schema
}

func (VenueSms) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("venue_id").Default(0).Comment("场馆id").Optional(),
		field.Int64("notice_count").Default(0).Comment("通知短信数量"),
		field.Int64("used_notice").Default(0).Comment("已用通知"),
	}
}

func (VenueSms) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (VenueSms) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("venue", Venue.Type).Ref("sms").Field("venue_id").Unique(),
	}
}

func (VenueSms) Indexes() []ent.Index {
	return []ent.Index{}
}

func (VenueSms) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_venue_sms"},
		entsql.WithComments(true),
	}
}
