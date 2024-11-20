package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type MemberNode struct {
	ent.Schema
}

func (MemberNode) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),

		field.Float("money_sum").Default(3).Comment("消费总金额").Optional(),

		field.Int64("product_id").Default(0).Comment("首次的产品").Optional(),
		field.String("product_name").Comment("首次的产品").Optional(),

		field.Int64("product_venue").Default(0).Comment("首次消费场馆").Optional(),
		field.String("product_venue_name").Comment("首次消费场馆").Optional(),

		field.Int64("entry_sum").Default(0).Comment("进馆总次数").Optional(),

		field.Time("entry_last_time").Optional().Comment("最后一次进馆时间"),

		field.Time("entry_deadline_time").Optional().Comment("进馆最后期限时间"),

		field.Time("class_last_time").Optional().Comment("最后一次上课时间"),

		field.Int64("relation_uid").Default(0).Comment("关联员工").Optional(),
		field.String("relation_uname").Comment("关联员工").Optional(),

		field.Int64("relation_mid").Default(0).Comment("关联会员").Optional(),
		field.String("relation_mame").Comment("关联会员").Optional(),

		field.Int64("create_id").
			Optional().
			Comment("创建人"),
		field.String("create_name").Comment("创建人").Optional(),
	}
}

func (MemberNode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (MemberNode) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("content", MemberContractContent.Type),

		//edge.From("member", Member.Type).Ref("member_contents").Field("member_id").Unique(),
		//edge.From("order", Order.Type).Ref("order_contents").Field("order_id").Unique(),
		//edge.From("member_product", MemberProduct.Type).Ref("member_product_contents").Field("member_product_id").Unique(),
	}
}

func (MemberNode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
		index.Fields("member_id"),
		index.Fields("member_product_id"),
	}
}

func (MemberNode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_contract", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
