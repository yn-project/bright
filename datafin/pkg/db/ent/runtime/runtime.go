// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/ent/schema"
	"yun.tea/block/bright/datafin/pkg/db/ent/topic"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	topicMixin := schema.Topic{}.Mixin()
	topic.Policy = privacy.NewPolicies(topicMixin[0], schema.Topic{})
	topic.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := topic.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	topicMixinFields0 := topicMixin[0].Fields()
	_ = topicMixinFields0
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescCreatedAt is the schema descriptor for created_at field.
	topicDescCreatedAt := topicMixinFields0[0].Descriptor()
	// topic.DefaultCreatedAt holds the default value on creation for the created_at field.
	topic.DefaultCreatedAt = topicDescCreatedAt.Default.(func() uint32)
	// topicDescUpdatedAt is the schema descriptor for updated_at field.
	topicDescUpdatedAt := topicMixinFields0[1].Descriptor()
	// topic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	topic.DefaultUpdatedAt = topicDescUpdatedAt.Default.(func() uint32)
	// topic.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	topic.UpdateDefaultUpdatedAt = topicDescUpdatedAt.UpdateDefault.(func() uint32)
	// topicDescDeletedAt is the schema descriptor for deleted_at field.
	topicDescDeletedAt := topicMixinFields0[2].Descriptor()
	// topic.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	topic.DefaultDeletedAt = topicDescDeletedAt.Default.(func() uint32)
	// topicDescID is the schema descriptor for id field.
	topicDescID := topicFields[0].Descriptor()
	// topic.DefaultID holds the default value on creation for the id field.
	topic.DefaultID = topicDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)