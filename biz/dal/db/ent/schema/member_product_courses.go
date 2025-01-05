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

type MemberProductCourses struct {
	ent.Schema
}

func (MemberProductCourses) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Default("").Comment("类型").Optional(),
		field.String("name").Default("").Comment("课名").Optional(),
		field.Int64("number").Default(0).Comment("节数").Optional(),
		field.Int64("member_product_id").Default(0).Comment("产品名称").Optional(),
		field.Int64("courses_id").Default(0).Comment("课名称").Optional(),
	}
}

func (MemberProductCourses) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberProductCourses) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("nodeC", MemberProduct.Type).
			Ref("memberCourses").
			Field("member_product_id").Unique(),

		edge.From("nodeL", MemberProduct.Type).
			Ref("memberLessons").
			Field("member_product_id").Unique(),
	}
}

func (MemberProductCourses) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

func (MemberProductCourses) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_product_courses", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
