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

type ProductProperty struct {
	ent.Schema
}

func (ProductProperty) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Comment("类型").Optional(),
		field.String("name").Comment("商品名").Optional(),
		field.Int64("stock").Comment("库存").Optional(),
		field.Int64("deadline").Comment("激活期限").Optional(),
		field.Int64("duration").Comment("有效期").Optional(),
		field.Int64("length").Comment("单次时长").Optional(),
		field.JSON("sales", []*base.Sales{}).Comment("售卖信息[售价等]").Optional(),
		field.Int64("is_sales").Comment("销售方式 1会员端").Optional(),
		field.Time("sign_sales_at").Comment("开始售卖时间").Optional(),
		field.Time("end_sales_at").Comment("结束售卖时间").Optional(),
		field.String("pic").Comment("主图").Optional(),
		field.Text("description").Comment("详情").Optional(),
	}
}

func (ProductProperty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (ProductProperty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tag", DictionaryDetail.Type),
		edge.To("contracts", Contract.Type),
	}
}

func (ProductProperty) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(),
	}
}

func (ProductProperty) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_property", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
