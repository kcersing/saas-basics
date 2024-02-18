package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Menu struct {
	ent.Schema
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:   "menu",
			Charset: "utf8",
		},
		entsql.WithComments(true),
	}
}

// Fields of the User.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("parent_id").Optional().Comment("").Nillable(),
		field.String("route_name").Optional().Comment("").Nillable(),
		field.String("route_path").Optional().Comment("").Nillable(),
		field.String("status").Optional().Comment("").Nillable(),
		field.String("menu_name").Optional().Comment("").Nillable(),
		field.String("menu_type").Optional().Comment("").Nillable(),
		field.String("icon_type").Optional().Comment("").Nillable(),
		field.String("icon").Optional().Comment("").Nillable(),
		field.String("i18n_key").Optional().Comment("").Nillable(),
		field.String("level").Optional().Comment("").Nillable(),
	}
}

func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Menu.Type).
			From("parent").
			Field("parent_id").
			Unique(),
	}
}

func (Menu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "parent_id"),
	}
}
