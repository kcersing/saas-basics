package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type Banner struct {
	ent.Schema
}

func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称"),
		field.String("pic").Comment("图片"),
		field.String("link").Comment("跳转链接"),
		field.Int64("is_show").Default(1).Comment("是否展示 1 展示 2 不展示").Optional(),
	}
}

func (Banner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Banner) Edges() []ent.Edge {
	return nil
}

func (Banner) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Banner) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "banner"},
		entsql.WithComments(true),
	}
}
