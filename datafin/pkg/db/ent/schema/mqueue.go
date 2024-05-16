package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/mixin"
)

type Mqueue struct {
	ent.Schema
}

func (Mqueue) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Mqueue) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").Unique(),
		field.String("remark").Optional(),
		field.String("topic_name").Unique(),
	}
}

func (Mqueue) Indexes() []ent.Index {
	return []ent.Index{}
}
