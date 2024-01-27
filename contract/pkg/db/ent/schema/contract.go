package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"yun.tea/block/bright/contract/pkg/db/mixin"
)

type Contract struct {
	ent.Schema
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name"),
		field.String("address"),
		field.String("remark").Optional(),
		field.String("version"),
	}
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{}
}
