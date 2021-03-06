package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type HasOwnerMixin struct {
	mixin.Schema
}

func (HasOwnerMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("ownerID").GoType(idPtrType).Optional().Nillable(),
		field.String("ownerType").Optional().Nillable(),
		field.String("owningUserID").GoType(idPtrType).Optional().Nillable(),
		field.UUID("ownerUID", uuid.New()).Optional().Nillable(),
	}
}

func (HasOwnerMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ownerType", "ownerID"),
	}
}

func (HasOwnerMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owningUser", User.Type).Field("owningUserID").Unique(),
	}
}
