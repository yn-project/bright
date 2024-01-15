package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"yun.tea/block/bright/account/pkg/db/mixin"
)

type Account struct {
	ent.Schema
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("address"),
		field.String("balance").Optional(),
		field.Bool("enable").Default(false),
		field.Bool("is_root").Default(false),
		field.String("remark").Optional(),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{}
}
