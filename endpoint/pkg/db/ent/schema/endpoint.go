package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"yun.tea/block/bright/endpoint/pkg/db/mixin"
)

type Endpoint struct {
	ent.Schema
}

func (Endpoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Endpoint) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name"),
		field.String("address"),
		field.String("state").Optional(),
		field.Uint32("rps").Default(10),
		field.String("remark").Optional(),
	}
}

func (Endpoint) Indexes() []ent.Index {
	return []ent.Index{}
}
