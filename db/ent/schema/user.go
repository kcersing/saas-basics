package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").Optional().Comment("姓名").Nillable(),
		field.String("password").Optional().Comment("密码").Nillable(),
		field.Int("gender").Optional().Default(0).Comment("性别").Nillable(),
		field.Int("age").Optional().Comment("年龄").Nillable(),
		field.String("introduce").Optional().Comment("介绍").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
