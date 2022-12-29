package mixin

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CreatedAt mixin for the created_at field.
type CreatedAt struct {
	mixin.Schema
}

// Fields of the CreatedAt mixin.
func (CreatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("created_at").
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entproto.Field(2),
			).
			Default(time.Now).
			Comment("The time when the entity was created."),
	}
}
