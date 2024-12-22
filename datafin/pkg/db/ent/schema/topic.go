package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/mixin"
)

type Topic struct {
	ent.Schema
}

func (Topic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("topic_id"),
		field.String("name"),
		field.String("contract"),
		field.String("type"),
		field.Bool("change_able"),
		field.Bool("on_chain"),
		field.String("remark").Optional(),
	}
}

func (Topic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic_id", "contract").Unique(),
		index.Fields("updated_at"),
	}
}
