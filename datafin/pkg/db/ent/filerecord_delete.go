// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"yun.tea/block/bright/datafin/pkg/db/ent/filerecord"
	"yun.tea/block/bright/datafin/pkg/db/ent/predicate"
)

// FileRecordDelete is the builder for deleting a FileRecord entity.
type FileRecordDelete struct {
	config
	hooks    []Hook
	mutation *FileRecordMutation
}

// Where appends a list predicates to the FileRecordDelete builder.
func (frd *FileRecordDelete) Where(ps ...predicate.FileRecord) *FileRecordDelete {
	frd.mutation.Where(ps...)
	return frd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (frd *FileRecordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(frd.hooks) == 0 {
		affected, err = frd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			frd.mutation = mutation
			affected, err = frd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(frd.hooks) - 1; i >= 0; i-- {
			if frd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = frd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, frd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (frd *FileRecordDelete) ExecX(ctx context.Context) int {
	n, err := frd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (frd *FileRecordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: filerecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: filerecord.FieldID,
			},
		},
	}
	if ps := frd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, frd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// FileRecordDeleteOne is the builder for deleting a single FileRecord entity.
type FileRecordDeleteOne struct {
	frd *FileRecordDelete
}

// Exec executes the deletion query.
func (frdo *FileRecordDeleteOne) Exec(ctx context.Context) error {
	n, err := frdo.frd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{filerecord.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (frdo *FileRecordDeleteOne) ExecX(ctx context.Context) {
	frdo.frd.ExecX(ctx)
}