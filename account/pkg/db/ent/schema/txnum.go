package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"yun.tea/block/bright/account/pkg/db/mixin"
)

type TxNum struct {
	ent.Schema
}

func (TxNum) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (TxNum) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id"),
		field.Uint32("time_at").Unique(),
		field.Uint32("num"),
	}
}

func (TxNum) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("time_at"),
	}
}
