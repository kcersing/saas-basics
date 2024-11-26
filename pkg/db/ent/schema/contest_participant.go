package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
)

type ContestParticipant struct {
	ent.Schema
}

func (ContestParticipant) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("contest_id").Comment("赛事id").Optional(),
		field.String("name").Comment("名称").Optional(),
		field.String("mobile").Comment("手机号").Optional(),

		field.String("fields").Comment("更多").Optional(),
	}
}

func (ContestParticipant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (ContestParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("contest", Contest.Type).Ref("contest_participants").Field("contest_id").Unique(),
	}
}

func (ContestParticipant) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "mobile").
			Unique(),
	}
}

func (ContestParticipant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "contest_participant"},
		entsql.WithComments(true),
	}
}
