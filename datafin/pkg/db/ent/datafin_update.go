// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"yun.tea/block/bright/datafin/pkg/db/ent/datafin"
	"yun.tea/block/bright/datafin/pkg/db/ent/predicate"
)

// DataFinUpdate is the builder for updating DataFin entities.
type DataFinUpdate struct {
	config
	hooks     []Hook
	mutation  *DataFinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DataFinUpdate builder.
func (dfu *DataFinUpdate) Where(ps ...predicate.DataFin) *DataFinUpdate {
	dfu.mutation.Where(ps...)
	return dfu
}

// SetCreatedAt sets the "created_at" field.
func (dfu *DataFinUpdate) SetCreatedAt(u uint32) *DataFinUpdate {
	dfu.mutation.ResetCreatedAt()
	dfu.mutation.SetCreatedAt(u)
	return dfu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dfu *DataFinUpdate) SetNillableCreatedAt(u *uint32) *DataFinUpdate {
	if u != nil {
		dfu.SetCreatedAt(*u)
	}
	return dfu
}

// AddCreatedAt adds u to the "created_at" field.
func (dfu *DataFinUpdate) AddCreatedAt(u int32) *DataFinUpdate {
	dfu.mutation.AddCreatedAt(u)
	return dfu
}

// SetUpdatedAt sets the "updated_at" field.
func (dfu *DataFinUpdate) SetUpdatedAt(u uint32) *DataFinUpdate {
	dfu.mutation.ResetUpdatedAt()
	dfu.mutation.SetUpdatedAt(u)
	return dfu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (dfu *DataFinUpdate) AddUpdatedAt(u int32) *DataFinUpdate {
	dfu.mutation.AddUpdatedAt(u)
	return dfu
}

// SetDeletedAt sets the "deleted_at" field.
func (dfu *DataFinUpdate) SetDeletedAt(u uint32) *DataFinUpdate {
	dfu.mutation.ResetDeletedAt()
	dfu.mutation.SetDeletedAt(u)
	return dfu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dfu *DataFinUpdate) SetNillableDeletedAt(u *uint32) *DataFinUpdate {
	if u != nil {
		dfu.SetDeletedAt(*u)
	}
	return dfu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (dfu *DataFinUpdate) AddDeletedAt(u int32) *DataFinUpdate {
	dfu.mutation.AddDeletedAt(u)
	return dfu
}

// SetTopicID sets the "topic_id" field.
func (dfu *DataFinUpdate) SetTopicID(s string) *DataFinUpdate {
	dfu.mutation.SetTopicID(s)
	return dfu
}

// SetDataID sets the "data_id" field.
func (dfu *DataFinUpdate) SetDataID(s string) *DataFinUpdate {
	dfu.mutation.SetDataID(s)
	return dfu
}

// SetDatafin sets the "datafin" field.
func (dfu *DataFinUpdate) SetDatafin(s string) *DataFinUpdate {
	dfu.mutation.SetDatafin(s)
	return dfu
}

// SetTxTime sets the "tx_time" field.
func (dfu *DataFinUpdate) SetTxTime(u uint32) *DataFinUpdate {
	dfu.mutation.ResetTxTime()
	dfu.mutation.SetTxTime(u)
	return dfu
}

// SetNillableTxTime sets the "tx_time" field if the given value is not nil.
func (dfu *DataFinUpdate) SetNillableTxTime(u *uint32) *DataFinUpdate {
	if u != nil {
		dfu.SetTxTime(*u)
	}
	return dfu
}

// AddTxTime adds u to the "tx_time" field.
func (dfu *DataFinUpdate) AddTxTime(u int32) *DataFinUpdate {
	dfu.mutation.AddTxTime(u)
	return dfu
}

// ClearTxTime clears the value of the "tx_time" field.
func (dfu *DataFinUpdate) ClearTxTime() *DataFinUpdate {
	dfu.mutation.ClearTxTime()
	return dfu
}

// SetTxHash sets the "tx_hash" field.
func (dfu *DataFinUpdate) SetTxHash(s string) *DataFinUpdate {
	dfu.mutation.SetTxHash(s)
	return dfu
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (dfu *DataFinUpdate) SetNillableTxHash(s *string) *DataFinUpdate {
	if s != nil {
		dfu.SetTxHash(*s)
	}
	return dfu
}

// ClearTxHash clears the value of the "tx_hash" field.
func (dfu *DataFinUpdate) ClearTxHash() *DataFinUpdate {
	dfu.mutation.ClearTxHash()
	return dfu
}

// SetState sets the "state" field.
func (dfu *DataFinUpdate) SetState(s string) *DataFinUpdate {
	dfu.mutation.SetState(s)
	return dfu
}

// SetRetries sets the "retries" field.
func (dfu *DataFinUpdate) SetRetries(u uint32) *DataFinUpdate {
	dfu.mutation.ResetRetries()
	dfu.mutation.SetRetries(u)
	return dfu
}

// AddRetries adds u to the "retries" field.
func (dfu *DataFinUpdate) AddRetries(u int32) *DataFinUpdate {
	dfu.mutation.AddRetries(u)
	return dfu
}

// SetRemark sets the "remark" field.
func (dfu *DataFinUpdate) SetRemark(s string) *DataFinUpdate {
	dfu.mutation.SetRemark(s)
	return dfu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dfu *DataFinUpdate) SetNillableRemark(s *string) *DataFinUpdate {
	if s != nil {
		dfu.SetRemark(*s)
	}
	return dfu
}

// ClearRemark clears the value of the "remark" field.
func (dfu *DataFinUpdate) ClearRemark() *DataFinUpdate {
	dfu.mutation.ClearRemark()
	return dfu
}

// Mutation returns the DataFinMutation object of the builder.
func (dfu *DataFinUpdate) Mutation() *DataFinMutation {
	return dfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dfu *DataFinUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := dfu.defaults(); err != nil {
		return 0, err
	}
	if len(dfu.hooks) == 0 {
		affected, err = dfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataFinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dfu.mutation = mutation
			affected, err = dfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dfu.hooks) - 1; i >= 0; i-- {
			if dfu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dfu *DataFinUpdate) SaveX(ctx context.Context) int {
	affected, err := dfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dfu *DataFinUpdate) Exec(ctx context.Context) error {
	_, err := dfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dfu *DataFinUpdate) ExecX(ctx context.Context) {
	if err := dfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dfu *DataFinUpdate) defaults() error {
	if _, ok := dfu.mutation.UpdatedAt(); !ok {
		if datafin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized datafin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := datafin.UpdateDefaultUpdatedAt()
		dfu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dfu *DataFinUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DataFinUpdate {
	dfu.modifiers = append(dfu.modifiers, modifiers...)
	return dfu
}

func (dfu *DataFinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   datafin.Table,
			Columns: datafin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: datafin.FieldID,
			},
		},
	}
	if ps := dfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dfu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldCreatedAt,
		})
	}
	if value, ok := dfu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldCreatedAt,
		})
	}
	if value, ok := dfu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldUpdatedAt,
		})
	}
	if value, ok := dfu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldUpdatedAt,
		})
	}
	if value, ok := dfu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldDeletedAt,
		})
	}
	if value, ok := dfu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldDeletedAt,
		})
	}
	if value, ok := dfu.mutation.TopicID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldTopicID,
		})
	}
	if value, ok := dfu.mutation.DataID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldDataID,
		})
	}
	if value, ok := dfu.mutation.Datafin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldDatafin,
		})
	}
	if value, ok := dfu.mutation.TxTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldTxTime,
		})
	}
	if value, ok := dfu.mutation.AddedTxTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldTxTime,
		})
	}
	if dfu.mutation.TxTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: datafin.FieldTxTime,
		})
	}
	if value, ok := dfu.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldTxHash,
		})
	}
	if dfu.mutation.TxHashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: datafin.FieldTxHash,
		})
	}
	if value, ok := dfu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldState,
		})
	}
	if value, ok := dfu.mutation.Retries(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldRetries,
		})
	}
	if value, ok := dfu.mutation.AddedRetries(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldRetries,
		})
	}
	if value, ok := dfu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldRemark,
		})
	}
	if dfu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: datafin.FieldRemark,
		})
	}
	_spec.Modifiers = dfu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, dfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{datafin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DataFinUpdateOne is the builder for updating a single DataFin entity.
type DataFinUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DataFinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (dfuo *DataFinUpdateOne) SetCreatedAt(u uint32) *DataFinUpdateOne {
	dfuo.mutation.ResetCreatedAt()
	dfuo.mutation.SetCreatedAt(u)
	return dfuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dfuo *DataFinUpdateOne) SetNillableCreatedAt(u *uint32) *DataFinUpdateOne {
	if u != nil {
		dfuo.SetCreatedAt(*u)
	}
	return dfuo
}

// AddCreatedAt adds u to the "created_at" field.
func (dfuo *DataFinUpdateOne) AddCreatedAt(u int32) *DataFinUpdateOne {
	dfuo.mutation.AddCreatedAt(u)
	return dfuo
}

// SetUpdatedAt sets the "updated_at" field.
func (dfuo *DataFinUpdateOne) SetUpdatedAt(u uint32) *DataFinUpdateOne {
	dfuo.mutation.ResetUpdatedAt()
	dfuo.mutation.SetUpdatedAt(u)
	return dfuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (dfuo *DataFinUpdateOne) AddUpdatedAt(u int32) *DataFinUpdateOne {
	dfuo.mutation.AddUpdatedAt(u)
	return dfuo
}

// SetDeletedAt sets the "deleted_at" field.
func (dfuo *DataFinUpdateOne) SetDeletedAt(u uint32) *DataFinUpdateOne {
	dfuo.mutation.ResetDeletedAt()
	dfuo.mutation.SetDeletedAt(u)
	return dfuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dfuo *DataFinUpdateOne) SetNillableDeletedAt(u *uint32) *DataFinUpdateOne {
	if u != nil {
		dfuo.SetDeletedAt(*u)
	}
	return dfuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (dfuo *DataFinUpdateOne) AddDeletedAt(u int32) *DataFinUpdateOne {
	dfuo.mutation.AddDeletedAt(u)
	return dfuo
}

// SetTopicID sets the "topic_id" field.
func (dfuo *DataFinUpdateOne) SetTopicID(s string) *DataFinUpdateOne {
	dfuo.mutation.SetTopicID(s)
	return dfuo
}

// SetDataID sets the "data_id" field.
func (dfuo *DataFinUpdateOne) SetDataID(s string) *DataFinUpdateOne {
	dfuo.mutation.SetDataID(s)
	return dfuo
}

// SetDatafin sets the "datafin" field.
func (dfuo *DataFinUpdateOne) SetDatafin(s string) *DataFinUpdateOne {
	dfuo.mutation.SetDatafin(s)
	return dfuo
}

// SetTxTime sets the "tx_time" field.
func (dfuo *DataFinUpdateOne) SetTxTime(u uint32) *DataFinUpdateOne {
	dfuo.mutation.ResetTxTime()
	dfuo.mutation.SetTxTime(u)
	return dfuo
}

// SetNillableTxTime sets the "tx_time" field if the given value is not nil.
func (dfuo *DataFinUpdateOne) SetNillableTxTime(u *uint32) *DataFinUpdateOne {
	if u != nil {
		dfuo.SetTxTime(*u)
	}
	return dfuo
}

// AddTxTime adds u to the "tx_time" field.
func (dfuo *DataFinUpdateOne) AddTxTime(u int32) *DataFinUpdateOne {
	dfuo.mutation.AddTxTime(u)
	return dfuo
}

// ClearTxTime clears the value of the "tx_time" field.
func (dfuo *DataFinUpdateOne) ClearTxTime() *DataFinUpdateOne {
	dfuo.mutation.ClearTxTime()
	return dfuo
}

// SetTxHash sets the "tx_hash" field.
func (dfuo *DataFinUpdateOne) SetTxHash(s string) *DataFinUpdateOne {
	dfuo.mutation.SetTxHash(s)
	return dfuo
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (dfuo *DataFinUpdateOne) SetNillableTxHash(s *string) *DataFinUpdateOne {
	if s != nil {
		dfuo.SetTxHash(*s)
	}
	return dfuo
}

// ClearTxHash clears the value of the "tx_hash" field.
func (dfuo *DataFinUpdateOne) ClearTxHash() *DataFinUpdateOne {
	dfuo.mutation.ClearTxHash()
	return dfuo
}

// SetState sets the "state" field.
func (dfuo *DataFinUpdateOne) SetState(s string) *DataFinUpdateOne {
	dfuo.mutation.SetState(s)
	return dfuo
}

// SetRetries sets the "retries" field.
func (dfuo *DataFinUpdateOne) SetRetries(u uint32) *DataFinUpdateOne {
	dfuo.mutation.ResetRetries()
	dfuo.mutation.SetRetries(u)
	return dfuo
}

// AddRetries adds u to the "retries" field.
func (dfuo *DataFinUpdateOne) AddRetries(u int32) *DataFinUpdateOne {
	dfuo.mutation.AddRetries(u)
	return dfuo
}

// SetRemark sets the "remark" field.
func (dfuo *DataFinUpdateOne) SetRemark(s string) *DataFinUpdateOne {
	dfuo.mutation.SetRemark(s)
	return dfuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dfuo *DataFinUpdateOne) SetNillableRemark(s *string) *DataFinUpdateOne {
	if s != nil {
		dfuo.SetRemark(*s)
	}
	return dfuo
}

// ClearRemark clears the value of the "remark" field.
func (dfuo *DataFinUpdateOne) ClearRemark() *DataFinUpdateOne {
	dfuo.mutation.ClearRemark()
	return dfuo
}

// Mutation returns the DataFinMutation object of the builder.
func (dfuo *DataFinUpdateOne) Mutation() *DataFinMutation {
	return dfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dfuo *DataFinUpdateOne) Select(field string, fields ...string) *DataFinUpdateOne {
	dfuo.fields = append([]string{field}, fields...)
	return dfuo
}

// Save executes the query and returns the updated DataFin entity.
func (dfuo *DataFinUpdateOne) Save(ctx context.Context) (*DataFin, error) {
	var (
		err  error
		node *DataFin
	)
	if err := dfuo.defaults(); err != nil {
		return nil, err
	}
	if len(dfuo.hooks) == 0 {
		node, err = dfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataFinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dfuo.mutation = mutation
			node, err = dfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dfuo.hooks) - 1; i >= 0; i-- {
			if dfuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dfuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dfuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DataFin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DataFinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dfuo *DataFinUpdateOne) SaveX(ctx context.Context) *DataFin {
	node, err := dfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dfuo *DataFinUpdateOne) Exec(ctx context.Context) error {
	_, err := dfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dfuo *DataFinUpdateOne) ExecX(ctx context.Context) {
	if err := dfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dfuo *DataFinUpdateOne) defaults() error {
	if _, ok := dfuo.mutation.UpdatedAt(); !ok {
		if datafin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized datafin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := datafin.UpdateDefaultUpdatedAt()
		dfuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dfuo *DataFinUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DataFinUpdateOne {
	dfuo.modifiers = append(dfuo.modifiers, modifiers...)
	return dfuo
}

func (dfuo *DataFinUpdateOne) sqlSave(ctx context.Context) (_node *DataFin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   datafin.Table,
			Columns: datafin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: datafin.FieldID,
			},
		},
	}
	id, ok := dfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DataFin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, datafin.FieldID)
		for _, f := range fields {
			if !datafin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != datafin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dfuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldCreatedAt,
		})
	}
	if value, ok := dfuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldCreatedAt,
		})
	}
	if value, ok := dfuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldUpdatedAt,
		})
	}
	if value, ok := dfuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldUpdatedAt,
		})
	}
	if value, ok := dfuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldDeletedAt,
		})
	}
	if value, ok := dfuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldDeletedAt,
		})
	}
	if value, ok := dfuo.mutation.TopicID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldTopicID,
		})
	}
	if value, ok := dfuo.mutation.DataID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldDataID,
		})
	}
	if value, ok := dfuo.mutation.Datafin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldDatafin,
		})
	}
	if value, ok := dfuo.mutation.TxTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldTxTime,
		})
	}
	if value, ok := dfuo.mutation.AddedTxTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldTxTime,
		})
	}
	if dfuo.mutation.TxTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: datafin.FieldTxTime,
		})
	}
	if value, ok := dfuo.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldTxHash,
		})
	}
	if dfuo.mutation.TxHashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: datafin.FieldTxHash,
		})
	}
	if value, ok := dfuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldState,
		})
	}
	if value, ok := dfuo.mutation.Retries(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldRetries,
		})
	}
	if value, ok := dfuo.mutation.AddedRetries(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldRetries,
		})
	}
	if value, ok := dfuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldRemark,
		})
	}
	if dfuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: datafin.FieldRemark,
		})
	}
	_spec.Modifiers = dfuo.modifiers
	_node = &DataFin{config: dfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{datafin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}