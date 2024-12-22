package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"yun.tea/block/bright/account/pkg/db/mixin"
	"yun.tea/block/bright/proto/bright/basetype"
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
		field.String("pri_key"),
		field.String("balance").Optional(),
		field.Uint64("nonce").Optional().Default(0),
		field.String("state").Default(basetype.AccountState_AccountUnkonwn.String()),
		field.Bool("is_root").Default(false),
		field.String("remark").Optional(),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at", "pri_key").Unique(),
		index.Fields("updated_at"),
	}
}
