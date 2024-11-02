package schema

import (
	"common/biz/dal/mysql/ent/schema/mixins"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"

	"entgo.io/ent/schema/field"

	"time"
)

type Face struct {
	ent.Schema
}

func (Face) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("user_id").Comment("user id").Optional(),
		field.String("identity_card").Comment("证件号").Optional(),
		field.String("face_identity_card").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("face_identity_card | 证件照片"),
		field.String("back_identity_card").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("back_identity_card | 证件照片"),

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
			Optional().
			Default(time.Now).
			Comment("人脸更新时间"),
	}
}

func (Face) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Face) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Face) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Face) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "faces"},
		entsql.WithComments(true),
	}
}
