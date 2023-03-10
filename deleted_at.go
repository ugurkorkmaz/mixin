package mixin

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// DeletedAt mixin for the deleted_at field.
type DeletedAt struct {
	mixin.Schema
}

// Fields of the DeletedAt mixin.
func (DeletedAt) Fields() []ent.Field {
	return []ent.Field{
		field.
			Time("deleted_at").
			Annotations(
				entgql.OrderField("DELETED_AT"),
				entproto.Field(4),
			).
			Optional().
			Nillable().
			Comment("The time when the entity was deleted."),
	}
}
