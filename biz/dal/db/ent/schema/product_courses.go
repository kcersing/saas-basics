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

type ProductCourses struct {
	ent.Schema
}

func (ProductCourses) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Default("").Comment("类型").Optional(),
		field.String("name").Default("").Comment("课名").Optional(),
		field.Int64("number").Default(0).Comment("节数").Optional(),
		field.Int64("product_id").Default(0).Comment("产品名称").Optional(),
		field.Int64("courses_id").Default(0).Comment("课名称").Optional(),
	}
}

func (ProductCourses) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (ProductCourses) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("product", Product.Type).
			Ref("courses").
			Field("product_id").Unique(),
	}
}

func (ProductCourses) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

func (ProductCourses) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_courses", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
