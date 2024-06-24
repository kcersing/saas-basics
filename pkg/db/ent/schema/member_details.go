package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"saas/pkg/db/ent/schema/mixins"
	"time"
)

type MemberDetails struct {
	ent.Schema
}

func (MemberDetails) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("member_id").Comment("会员id").Optional(),
		field.String("email").Optional().Comment("email | 邮箱号"),
		field.String("wecom").Optional().Comment("wecom | 微信号"),
		field.Int64("gender").Default(3).Comment("性别 | [0:女性;1:男性;3:保密]").Optional(),
		field.Time("birthday").Comment("出生日期").Optional(),

		field.String("identity_card").Comment("正面证件号").Optional(),
		field.String("face_identity_card").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("face_identity_card | 正面证件照片"),
		field.String("back_identity_card").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("back_identity_card | 反面证件照片"),

		field.String("face_pic").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("face_pic | 人脸照片"),

		field.String("face_eigenvalue").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("人脸特征值"),

		field.Time("face_pic_updated_time").
			Immutable().
			Default(time.Now).
			Comment("人脸更新时间"),

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

func (MemberDetails) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("info", Member.Type).
			Ref("member_details").Unique().
			Field("member_id").Unique(),
	}
}

func (MemberDetails) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id"),
	}
}

func (MemberDetails) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_details", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
