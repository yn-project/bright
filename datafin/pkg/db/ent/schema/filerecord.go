package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/mixin"
)

type FileRecord struct {
	ent.Schema
}

func (FileRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (FileRecord) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("file_name"),
		field.String("topic_id"),
		field.Uint32("record_num").Default(0),
		field.String("sha1_sum"),
		field.String("state"),
		field.String("remark").Optional(),
	}
}

func (FileRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("file_name"),
		index.Fields("topic_id"),
		index.Fields("updated_at"),
	}
}
