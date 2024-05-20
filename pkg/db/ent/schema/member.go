package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"saas/pkg/db/ent/schema/mixins"
)

type Member struct {
	ent.Schema
}

func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("password").Optional().Comment("password | 密码"),
		field.String("name").Optional().Comment("name | 名称"),
		field.String("mobile").Optional().Comment("mobile number | 手机号"),
		field.String("email").Optional().Comment("email | 邮箱号"),
		field.String("wecom").Optional().Comment("wecom | 微信号"),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("avatar | 头像路径"),
		field.Int64("condition").
			Default(1).
			Optional().
			Comment("状态[0:潜在;1:正式;2:到期]"),
		field.Int64("create_id").
			Optional().
			Comment("创建人"),
	}
}

func (Member) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("member_details", MemberDetails.Type),
		edge.To("member_notes", MemberNote.Type),
		edge.To("member_orders", Order.Type),
		edge.To("member_products", MemberProduct.Type),
		edge.To("member_entry", EntryLogs.Type),
		edge.To("member_contents", MemberContract.Type),
	}
}

func (Member) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member"},
		entsql.WithComments(true),
	}
}
