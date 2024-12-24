package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"saas/biz/dal/db/ent/schema/mixins"
)

type Sms struct {
	ent.Schema
}

func (Sms) Fields() []ent.Field {
	return []ent.Field{
		field.String("mobile").Comment("手机号"),
		field.String("biz_id").Comment("BizId"),
		field.String("code").Comment("内容"),
		field.String("template").Comment("短信模板"),
	}
}

func (Sms) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Sms) Edges() []ent.Edge {
	return nil
}

func (Sms) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Sms) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_sms"},
		entsql.WithComments(true),
	}
}
