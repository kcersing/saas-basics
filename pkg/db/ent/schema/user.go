package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/rs/xid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("account_id").
			GoType(xid.ID{}).
			DefaultFunc(xid.New),
		field.String("username").Optional().Comment("姓名").Nillable(),
		field.String("password").Optional().Comment("密码").Nillable(),
		field.String("mobile").Optional().Comment("联系方式").Nillable(),
		field.String("gender").Optional().Default("0").Comment("性别").Nillable(),
		field.String("age").Optional().Comment("年龄").Nillable(),
		field.String("introduce").Optional().Comment("介绍").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "account_id"),
	}
}
