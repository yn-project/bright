package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"yun.tea/block/bright/account/pkg/db/mixin"
)

type BlockNum struct {
	ent.Schema
}

func (BlockNum) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (BlockNum) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id"),
		field.Uint32("time_at").Unique(),
		field.Uint64("height"),
	}
}

func (BlockNum) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("time_at"),
		index.Fields("updated_at"),
	}
}
