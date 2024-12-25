package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/biz/dal/db/ent/schema/mixins"
	"saas/idl_gen/model/base"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Default("").Comment("类型").Optional(),
		field.String("name").Default("").Comment("商品名").Optional(),
		field.Int64("stock").Default(0).Comment("库存").Optional(),
		field.Int64("deadline").Default(0).Comment("激活期限").Optional(),
		field.Int64("duration").Default(0).Comment("有效期").Optional(),
		field.Int64("length").Default(0).Comment("课程课时").Optional(),

		field.Float("price").Default(0).Comment("售价").Optional(),
		field.Int64("times").Default(0).Comment("次数").Optional(),
		field.Int64("is_lessons").Default(1).Comment("团课预约 1支持2不支持").Optional(),

		field.JSON("sales", []*base.Sales{}).Comment("售卖信息[售价等]").Optional(),
		field.Int64("is_sales").Default(1).Comment("销售方式 1会员端").Optional(),

		field.Time("sign_sales_at").Comment("开始售卖时间").Optional(),
		field.Time("end_sales_at").Comment("结束售卖时间").Optional(),
		field.String("pic").Default("").Comment("主图").Optional(),
		field.Text("description").Default("").Comment("详情").Optional(),
	}
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tag", DictionaryDetail.Type),
		edge.To("contracts", Contract.Type),

		edge.To("lessons", Product.Type).
			From("products"),
	}
}

func (Product) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
