// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/google/uuid"
	"yun.tea/block/bright/user/pkg/db/ent/schema"
	"yun.tea/block/bright/user/pkg/db/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(userMixin[0], schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() uint32)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() uint32)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() uint32)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields0[2].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(func() uint32)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2" // Version of ent codegen.
)
