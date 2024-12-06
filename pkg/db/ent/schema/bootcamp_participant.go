package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
)

type BootcampParticipant struct {
	ent.Schema
}

func (BootcampParticipant) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("bootcamp_id").Comment("赛事id").Optional(),
		field.String("name").Comment("名称").Optional(),
		field.String("mobile").Comment("手机号").Optional(),
		field.Text("fields").Comment("更多").Optional(),
		field.Int64("order_id").Default(0).Comment("订单ID").Optional(),
		field.String("order_sn").Default("").Comment("订单编号").Optional(),
		field.Float("fee").Comment("费用").Optional(),
	}
}

func (BootcampParticipant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (BootcampParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bootcamp", Bootcamp.Type).Ref("bootcamp_participants").Field("bootcamp_id").Unique(),
	}
}

func (BootcampParticipant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "mobile").
			Unique(),
	}
}

func (BootcampParticipant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "bootcamp_participant"},
		entsql.WithComments(true),
	}
}
