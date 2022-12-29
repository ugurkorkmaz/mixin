package mixin

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// UUID is a mixin that adds a UUID field to an entity.
type Id struct {
	mixin.Schema
}

// Fields of the UUID mixin.
func (Id) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ID"),
				entproto.Field(1),
			).
			Default(uuid.New).
			Immutable().
			Comment("The unique identifier of the entity."),
	}
}
