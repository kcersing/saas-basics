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

type MemberProduct struct {
	ent.Schema
}

func (MemberProduct) Fields() []ent.Field {

	return []ent.Field{
		field.String("sn").Comment("编号").Optional(),
		field.String("type").Comment("类型").Optional(),
		field.String("sub_type").Default("").Comment("次级类型").Optional(),

		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("product_id").Comment("产品ID").Optional(),
		field.Int64("venue_id").Comment("场馆ID").Optional(),
		field.Int64("order_id").Comment("订单ID").Optional(),
		field.String("name").Comment("产品名称").Optional(),
		field.Float("price").Comment("单价").Optional(),
		field.Float("fee").Comment("总价").Optional(),
		field.Int64("duration").Comment("总时长").Optional(),
		field.Int64("length").Comment("单次时长").Optional(),
		field.Int64("number").Default(0).Comment("总次数").Optional(),
		field.Int64("number_surplus").Default(0).Comment("剩余次数").Optional(),
		field.Int64("is_course").Default(1).Comment("課包 课程1不限2指定").Optional(),
		field.Int64("deadline").Comment("激活期限").Optional(),
		field.Time("validity_at").Comment("生效时间").Optional(),
		field.Time("cancel_at").Comment("作废时间").Optional(),
	}
}

func (MemberProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("members", Member.Type).Ref("member_products").Field("member_id").Unique(),

		edge.To("member_product_entry", EntryLogs.Type),
		edge.To("member_product_contents", MemberContract.Type),

		edge.To("memberCourses", MemberProductCourses.Type),
		edge.To("memberLessons", MemberProductCourses.Type),
	}
}

func (MemberProduct) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("venue_id"),
		index.Fields("member_id"),
		index.Fields("product_id"),
		index.Fields("order_id"),
	}
}

func (MemberProduct) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "member_product",
			//OnDelete: entsql.Cascade,
			Options: "AUTO_INCREMENT = 100000",
		},
		entsql.WithComments(true),
	}
}
