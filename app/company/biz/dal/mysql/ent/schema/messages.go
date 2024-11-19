package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Messages struct {
	ent.Schema
}

func (Messages) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("类型[1:用户user;2:会员member]"),
		field.String("to_user_id").Comment("该消息接受者ID"),
		field.String("from_user_id").Comment("该消息发送者ID"),
		field.String("content").Comment("消息内容"),
	}
}

func (Messages) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Messages) Edges() []ent.Edge {
	return nil
}

func (Messages) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("to_user_id", "from_user_id"),
	}
}

func (Messages) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "messages"},
		entsql.WithComments(true),
	}
}
