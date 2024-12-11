package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/biz/dal/db/ent/schema/mixins"
)

type CommunityParticipant struct {
	ent.Schema
}

func (CommunityParticipant) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("community_id").Comment("社群id").Optional(),
		field.String("name").Comment("名称").Optional(),
		field.String("mobile").Comment("手机号").Optional(),
		field.Text("fields").Comment("更多").Optional(),
		field.Int64("order_id").Default(0).Comment("订单ID").Optional(),
		field.String("order_sn").Default("").Comment("订单编号").Optional(),
		field.Float("fee").Comment("费用").Optional(),
		field.Int64("member_id").Default(0).Comment("会员ID").Optional(),
	}
}

func (CommunityParticipant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (CommunityParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("community", Community.Type).Ref("community_participants").Field("community_id").Unique(),
		edge.From("members", Member.Type).Ref("member_communitys"),
	}
}

func (CommunityParticipant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "mobile").
			Unique(),
	}
}

func (CommunityParticipant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "community_participant"},
		entsql.WithComments(true),
	}
}
