package mixin

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// UpdatedAt mixin for the updated_at field.
type UpdatedAt struct {
	mixin.Schema
}

// Fields of the UpdatedAt mixin.
func (UpdatedAt) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("updated_at").
			Annotations(
				entgql.OrderField("UPDATED_AT"),
				entproto.Field(3),
			).
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("The time when the entity was updated."),
	}
}
