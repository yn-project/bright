package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/mixin"
)

type DataFin struct {
	ent.Schema
}

func (DataFin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (DataFin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("topic_id"),
		field.String("data_id"),
		field.String("datafin"),
		field.Uint32("tx_time").Optional(),
		field.String("tx_hash").Optional(),
		field.String("state"),
		field.Uint32("retries"),
		field.String("remark").Optional(),
	}
}

func (DataFin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic_id", "data_id"),
		index.Fields("updated_at"),
	}
}
