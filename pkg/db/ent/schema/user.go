package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("username").Optional().Comment("姓名").Nillable(),
		field.String("password").Optional().Comment("密码").Nillable(),
		field.String("phone_number").Optional().Comment("联系方式").Nillable(),
		field.Int("gender").Optional().Default(0).Comment("性别").Nillable(),
		field.Int("age").Optional().Comment("年龄").Nillable(),
		field.String("introduce").Optional().Comment("介绍").Nillable(),

		field.String("account_id").Optional().Comment("").Nillable(),
		field.String("avatar_blob_id").Optional().Comment("").Nillable(),
		field.String("open_id").Optional().Comment("").Nillable(),
		field.Int("balance").Optional().Comment("").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
