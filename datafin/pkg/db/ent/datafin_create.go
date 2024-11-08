// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/ent/datafin"
)

// DataFinCreate is the builder for creating a DataFin entity.
type DataFinCreate struct {
	config
	mutation *DataFinMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dfc *DataFinCreate) SetCreatedAt(u uint32) *DataFinCreate {
	dfc.mutation.SetCreatedAt(u)
	return dfc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dfc *DataFinCreate) SetNillableCreatedAt(u *uint32) *DataFinCreate {
	if u != nil {
		dfc.SetCreatedAt(*u)
	}
	return dfc
}

// SetUpdatedAt sets the "updated_at" field.
func (dfc *DataFinCreate) SetUpdatedAt(u uint32) *DataFinCreate {
	dfc.mutation.SetUpdatedAt(u)
	return dfc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dfc *DataFinCreate) SetNillableUpdatedAt(u *uint32) *DataFinCreate {
	if u != nil {
		dfc.SetUpdatedAt(*u)
	}
	return dfc
}

// SetDeletedAt sets the "deleted_at" field.
func (dfc *DataFinCreate) SetDeletedAt(u uint32) *DataFinCreate {
	dfc.mutation.SetDeletedAt(u)
	return dfc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dfc *DataFinCreate) SetNillableDeletedAt(u *uint32) *DataFinCreate {
	if u != nil {
		dfc.SetDeletedAt(*u)
	}
	return dfc
}

// SetTopicID sets the "topic_id" field.
func (dfc *DataFinCreate) SetTopicID(s string) *DataFinCreate {
	dfc.mutation.SetTopicID(s)
	return dfc
}

// SetDataID sets the "data_id" field.
func (dfc *DataFinCreate) SetDataID(s string) *DataFinCreate {
	dfc.mutation.SetDataID(s)
	return dfc
}

// SetDatafin sets the "datafin" field.
func (dfc *DataFinCreate) SetDatafin(s string) *DataFinCreate {
	dfc.mutation.SetDatafin(s)
	return dfc
}

// SetTxTime sets the "tx_time" field.
func (dfc *DataFinCreate) SetTxTime(u uint32) *DataFinCreate {
	dfc.mutation.SetTxTime(u)
	return dfc
}

// SetNillableTxTime sets the "tx_time" field if the given value is not nil.
func (dfc *DataFinCreate) SetNillableTxTime(u *uint32) *DataFinCreate {
	if u != nil {
		dfc.SetTxTime(*u)
	}
	return dfc
}

// SetTxHash sets the "tx_hash" field.
func (dfc *DataFinCreate) SetTxHash(s string) *DataFinCreate {
	dfc.mutation.SetTxHash(s)
	return dfc
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (dfc *DataFinCreate) SetNillableTxHash(s *string) *DataFinCreate {
	if s != nil {
		dfc.SetTxHash(*s)
	}
	return dfc
}

// SetState sets the "state" field.
func (dfc *DataFinCreate) SetState(s string) *DataFinCreate {
	dfc.mutation.SetState(s)
	return dfc
}

// SetRetries sets the "retries" field.
func (dfc *DataFinCreate) SetRetries(u uint32) *DataFinCreate {
	dfc.mutation.SetRetries(u)
	return dfc
}

// SetRemark sets the "remark" field.
func (dfc *DataFinCreate) SetRemark(s string) *DataFinCreate {
	dfc.mutation.SetRemark(s)
	return dfc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dfc *DataFinCreate) SetNillableRemark(s *string) *DataFinCreate {
	if s != nil {
		dfc.SetRemark(*s)
	}
	return dfc
}

// SetID sets the "id" field.
func (dfc *DataFinCreate) SetID(u uuid.UUID) *DataFinCreate {
	dfc.mutation.SetID(u)
	return dfc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dfc *DataFinCreate) SetNillableID(u *uuid.UUID) *DataFinCreate {
	if u != nil {
		dfc.SetID(*u)
	}
	return dfc
}

// Mutation returns the DataFinMutation object of the builder.
func (dfc *DataFinCreate) Mutation() *DataFinMutation {
	return dfc.mutation
}

// Save creates the DataFin in the database.
func (dfc *DataFinCreate) Save(ctx context.Context) (*DataFin, error) {
	var (
		err  error
		node *DataFin
	)
	if err := dfc.defaults(); err != nil {
		return nil, err
	}
	if len(dfc.hooks) == 0 {
		if err = dfc.check(); err != nil {
			return nil, err
		}
		node, err = dfc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DataFinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dfc.check(); err != nil {
				return nil, err
			}
			dfc.mutation = mutation
			if node, err = dfc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dfc.hooks) - 1; i >= 0; i-- {
			if dfc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dfc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dfc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (dfc *DataFinCreate) SaveX(ctx context.Context) *DataFin {
	v, err := dfc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dfc *DataFinCreate) Exec(ctx context.Context) error {
	_, err := dfc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dfc *DataFinCreate) ExecX(ctx context.Context) {
	if err := dfc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dfc *DataFinCreate) defaults() error {
	if _, ok := dfc.mutation.CreatedAt(); !ok {
		if datafin.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized datafin.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := datafin.DefaultCreatedAt()
		dfc.mutation.SetCreatedAt(v)
	}
	if _, ok := dfc.mutation.UpdatedAt(); !ok {
		if datafin.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized datafin.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := datafin.DefaultUpdatedAt()
		dfc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dfc.mutation.DeletedAt(); !ok {
		if datafin.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized datafin.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := datafin.DefaultDeletedAt()
		dfc.mutation.SetDeletedAt(v)
	}
	if _, ok := dfc.mutation.ID(); !ok {
		if datafin.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized datafin.DefaultID (forgotten import ent/runtime?)")
		}
		v := datafin.DefaultID()
		dfc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dfc *DataFinCreate) check() error {
	if _, ok := dfc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DataFin.created_at"`)}
	}
	if _, ok := dfc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DataFin.updated_at"`)}
	}
	if _, ok := dfc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "DataFin.deleted_at"`)}
	}
	if _, ok := dfc.mutation.TopicID(); !ok {
		return &ValidationError{Name: "topic_id", err: errors.New(`ent: missing required field "DataFin.topic_id"`)}
	}
	if _, ok := dfc.mutation.DataID(); !ok {
		return &ValidationError{Name: "data_id", err: errors.New(`ent: missing required field "DataFin.data_id"`)}
	}
	if _, ok := dfc.mutation.Datafin(); !ok {
		return &ValidationError{Name: "datafin", err: errors.New(`ent: missing required field "DataFin.datafin"`)}
	}
	if _, ok := dfc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "DataFin.state"`)}
	}
	if _, ok := dfc.mutation.Retries(); !ok {
		return &ValidationError{Name: "retries", err: errors.New(`ent: missing required field "DataFin.retries"`)}
	}
	return nil
}

func (dfc *DataFinCreate) sqlSave(ctx context.Context) (*DataFin, error) {
	_node, _spec := dfc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dfc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (dfc *DataFinCreate) createSpec() (*DataFin, *sqlgraph.CreateSpec) {
	var (
		_node = &DataFin{config: dfc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: datafin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: datafin.FieldID,
			},
		}
	)
	_spec.OnConflict = dfc.conflict
	if id, ok := dfc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dfc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dfc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dfc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := dfc.mutation.TopicID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldTopicID,
		})
		_node.TopicID = value
	}
	if value, ok := dfc.mutation.DataID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldDataID,
		})
		_node.DataID = value
	}
	if value, ok := dfc.mutation.Datafin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldDatafin,
		})
		_node.Datafin = value
	}
	if value, ok := dfc.mutation.TxTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldTxTime,
		})
		_node.TxTime = value
	}
	if value, ok := dfc.mutation.TxHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldTxHash,
		})
		_node.TxHash = value
	}
	if value, ok := dfc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldState,
		})
		_node.State = value
	}
	if value, ok := dfc.mutation.Retries(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: datafin.FieldRetries,
		})
		_node.Retries = value
	}
	if value, ok := dfc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datafin.FieldRemark,
		})
		_node.Remark = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DataFin.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DataFinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dfc *DataFinCreate) OnConflict(opts ...sql.ConflictOption) *DataFinUpsertOne {
	dfc.conflict = opts
	return &DataFinUpsertOne{
		create: dfc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DataFin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dfc *DataFinCreate) OnConflictColumns(columns ...string) *DataFinUpsertOne {
	dfc.conflict = append(dfc.conflict, sql.ConflictColumns(columns...))
	return &DataFinUpsertOne{
		create: dfc,
	}
}

type (
	// DataFinUpsertOne is the builder for "upsert"-ing
	//  one DataFin node.
	DataFinUpsertOne struct {
		create *DataFinCreate
	}

	// DataFinUpsert is the "OnConflict" setter.
	DataFinUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *DataFinUpsert) SetCreatedAt(v uint32) *DataFinUpsert {
	u.Set(datafin.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateCreatedAt() *DataFinUpsert {
	u.SetExcluded(datafin.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DataFinUpsert) AddCreatedAt(v uint32) *DataFinUpsert {
	u.Add(datafin.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DataFinUpsert) SetUpdatedAt(v uint32) *DataFinUpsert {
	u.Set(datafin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateUpdatedAt() *DataFinUpsert {
	u.SetExcluded(datafin.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DataFinUpsert) AddUpdatedAt(v uint32) *DataFinUpsert {
	u.Add(datafin.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DataFinUpsert) SetDeletedAt(v uint32) *DataFinUpsert {
	u.Set(datafin.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateDeletedAt() *DataFinUpsert {
	u.SetExcluded(datafin.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DataFinUpsert) AddDeletedAt(v uint32) *DataFinUpsert {
	u.Add(datafin.FieldDeletedAt, v)
	return u
}

// SetTopicID sets the "topic_id" field.
func (u *DataFinUpsert) SetTopicID(v string) *DataFinUpsert {
	u.Set(datafin.FieldTopicID, v)
	return u
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateTopicID() *DataFinUpsert {
	u.SetExcluded(datafin.FieldTopicID)
	return u
}

// SetDataID sets the "data_id" field.
func (u *DataFinUpsert) SetDataID(v string) *DataFinUpsert {
	u.Set(datafin.FieldDataID, v)
	return u
}

// UpdateDataID sets the "data_id" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateDataID() *DataFinUpsert {
	u.SetExcluded(datafin.FieldDataID)
	return u
}

// SetDatafin sets the "datafin" field.
func (u *DataFinUpsert) SetDatafin(v string) *DataFinUpsert {
	u.Set(datafin.FieldDatafin, v)
	return u
}

// UpdateDatafin sets the "datafin" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateDatafin() *DataFinUpsert {
	u.SetExcluded(datafin.FieldDatafin)
	return u
}

// SetTxTime sets the "tx_time" field.
func (u *DataFinUpsert) SetTxTime(v uint32) *DataFinUpsert {
	u.Set(datafin.FieldTxTime, v)
	return u
}

// UpdateTxTime sets the "tx_time" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateTxTime() *DataFinUpsert {
	u.SetExcluded(datafin.FieldTxTime)
	return u
}

// AddTxTime adds v to the "tx_time" field.
func (u *DataFinUpsert) AddTxTime(v uint32) *DataFinUpsert {
	u.Add(datafin.FieldTxTime, v)
	return u
}

// ClearTxTime clears the value of the "tx_time" field.
func (u *DataFinUpsert) ClearTxTime() *DataFinUpsert {
	u.SetNull(datafin.FieldTxTime)
	return u
}

// SetTxHash sets the "tx_hash" field.
func (u *DataFinUpsert) SetTxHash(v string) *DataFinUpsert {
	u.Set(datafin.FieldTxHash, v)
	return u
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateTxHash() *DataFinUpsert {
	u.SetExcluded(datafin.FieldTxHash)
	return u
}

// ClearTxHash clears the value of the "tx_hash" field.
func (u *DataFinUpsert) ClearTxHash() *DataFinUpsert {
	u.SetNull(datafin.FieldTxHash)
	return u
}

// SetState sets the "state" field.
func (u *DataFinUpsert) SetState(v string) *DataFinUpsert {
	u.Set(datafin.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateState() *DataFinUpsert {
	u.SetExcluded(datafin.FieldState)
	return u
}

// SetRetries sets the "retries" field.
func (u *DataFinUpsert) SetRetries(v uint32) *DataFinUpsert {
	u.Set(datafin.FieldRetries, v)
	return u
}

// UpdateRetries sets the "retries" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateRetries() *DataFinUpsert {
	u.SetExcluded(datafin.FieldRetries)
	return u
}

// AddRetries adds v to the "retries" field.
func (u *DataFinUpsert) AddRetries(v uint32) *DataFinUpsert {
	u.Add(datafin.FieldRetries, v)
	return u
}

// SetRemark sets the "remark" field.
func (u *DataFinUpsert) SetRemark(v string) *DataFinUpsert {
	u.Set(datafin.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DataFinUpsert) UpdateRemark() *DataFinUpsert {
	u.SetExcluded(datafin.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *DataFinUpsert) ClearRemark() *DataFinUpsert {
	u.SetNull(datafin.FieldRemark)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DataFin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(datafin.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DataFinUpsertOne) UpdateNewValues() *DataFinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(datafin.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DataFin.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DataFinUpsertOne) Ignore() *DataFinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DataFinUpsertOne) DoNothing() *DataFinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DataFinCreate.OnConflict
// documentation for more info.
func (u *DataFinUpsertOne) Update(set func(*DataFinUpsert)) *DataFinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DataFinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DataFinUpsertOne) SetCreatedAt(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DataFinUpsertOne) AddCreatedAt(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateCreatedAt() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DataFinUpsertOne) SetUpdatedAt(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DataFinUpsertOne) AddUpdatedAt(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateUpdatedAt() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DataFinUpsertOne) SetDeletedAt(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DataFinUpsertOne) AddDeletedAt(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateDeletedAt() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTopicID sets the "topic_id" field.
func (u *DataFinUpsertOne) SetTopicID(v string) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetTopicID(v)
	})
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateTopicID() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateTopicID()
	})
}

// SetDataID sets the "data_id" field.
func (u *DataFinUpsertOne) SetDataID(v string) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetDataID(v)
	})
}

// UpdateDataID sets the "data_id" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateDataID() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateDataID()
	})
}

// SetDatafin sets the "datafin" field.
func (u *DataFinUpsertOne) SetDatafin(v string) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetDatafin(v)
	})
}

// UpdateDatafin sets the "datafin" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateDatafin() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateDatafin()
	})
}

// SetTxTime sets the "tx_time" field.
func (u *DataFinUpsertOne) SetTxTime(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetTxTime(v)
	})
}

// AddTxTime adds v to the "tx_time" field.
func (u *DataFinUpsertOne) AddTxTime(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.AddTxTime(v)
	})
}

// UpdateTxTime sets the "tx_time" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateTxTime() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateTxTime()
	})
}

// ClearTxTime clears the value of the "tx_time" field.
func (u *DataFinUpsertOne) ClearTxTime() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.ClearTxTime()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *DataFinUpsertOne) SetTxHash(v string) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateTxHash() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateTxHash()
	})
}

// ClearTxHash clears the value of the "tx_hash" field.
func (u *DataFinUpsertOne) ClearTxHash() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.ClearTxHash()
	})
}

// SetState sets the "state" field.
func (u *DataFinUpsertOne) SetState(v string) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateState() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateState()
	})
}

// SetRetries sets the "retries" field.
func (u *DataFinUpsertOne) SetRetries(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetRetries(v)
	})
}

// AddRetries adds v to the "retries" field.
func (u *DataFinUpsertOne) AddRetries(v uint32) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.AddRetries(v)
	})
}

// UpdateRetries sets the "retries" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateRetries() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateRetries()
	})
}

// SetRemark sets the "remark" field.
func (u *DataFinUpsertOne) SetRemark(v string) *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DataFinUpsertOne) UpdateRemark() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *DataFinUpsertOne) ClearRemark() *DataFinUpsertOne {
	return u.Update(func(s *DataFinUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *DataFinUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DataFinCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DataFinUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DataFinUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DataFinUpsertOne.ID is not supported by MySQL driver. Use DataFinUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DataFinUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DataFinCreateBulk is the builder for creating many DataFin entities in bulk.
type DataFinCreateBulk struct {
	config
	builders []*DataFinCreate
	conflict []sql.ConflictOption
}

// Save creates the DataFin entities in the database.
func (dfcb *DataFinCreateBulk) Save(ctx context.Context) ([]*DataFin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dfcb.builders))
	nodes := make([]*DataFin, len(dfcb.builders))
	mutators := make([]Mutator, len(dfcb.builders))
	for i := range dfcb.builders {
		func(i int, root context.Context) {
			builder := dfcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DataFinMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dfcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dfcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dfcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dfcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dfcb *DataFinCreateBulk) SaveX(ctx context.Context) []*DataFin {
	v, err := dfcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dfcb *DataFinCreateBulk) Exec(ctx context.Context) error {
	_, err := dfcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dfcb *DataFinCreateBulk) ExecX(ctx context.Context) {
	if err := dfcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DataFin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DataFinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dfcb *DataFinCreateBulk) OnConflict(opts ...sql.ConflictOption) *DataFinUpsertBulk {
	dfcb.conflict = opts
	return &DataFinUpsertBulk{
		create: dfcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DataFin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dfcb *DataFinCreateBulk) OnConflictColumns(columns ...string) *DataFinUpsertBulk {
	dfcb.conflict = append(dfcb.conflict, sql.ConflictColumns(columns...))
	return &DataFinUpsertBulk{
		create: dfcb,
	}
}

// DataFinUpsertBulk is the builder for "upsert"-ing
// a bulk of DataFin nodes.
type DataFinUpsertBulk struct {
	create *DataFinCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DataFin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(datafin.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DataFinUpsertBulk) UpdateNewValues() *DataFinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(datafin.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DataFin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DataFinUpsertBulk) Ignore() *DataFinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DataFinUpsertBulk) DoNothing() *DataFinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DataFinCreateBulk.OnConflict
// documentation for more info.
func (u *DataFinUpsertBulk) Update(set func(*DataFinUpsert)) *DataFinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DataFinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DataFinUpsertBulk) SetCreatedAt(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DataFinUpsertBulk) AddCreatedAt(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateCreatedAt() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DataFinUpsertBulk) SetUpdatedAt(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DataFinUpsertBulk) AddUpdatedAt(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateUpdatedAt() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DataFinUpsertBulk) SetDeletedAt(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DataFinUpsertBulk) AddDeletedAt(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateDeletedAt() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTopicID sets the "topic_id" field.
func (u *DataFinUpsertBulk) SetTopicID(v string) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetTopicID(v)
	})
}

// UpdateTopicID sets the "topic_id" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateTopicID() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateTopicID()
	})
}

// SetDataID sets the "data_id" field.
func (u *DataFinUpsertBulk) SetDataID(v string) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetDataID(v)
	})
}

// UpdateDataID sets the "data_id" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateDataID() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateDataID()
	})
}

// SetDatafin sets the "datafin" field.
func (u *DataFinUpsertBulk) SetDatafin(v string) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetDatafin(v)
	})
}

// UpdateDatafin sets the "datafin" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateDatafin() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateDatafin()
	})
}

// SetTxTime sets the "tx_time" field.
func (u *DataFinUpsertBulk) SetTxTime(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetTxTime(v)
	})
}

// AddTxTime adds v to the "tx_time" field.
func (u *DataFinUpsertBulk) AddTxTime(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.AddTxTime(v)
	})
}

// UpdateTxTime sets the "tx_time" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateTxTime() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateTxTime()
	})
}

// ClearTxTime clears the value of the "tx_time" field.
func (u *DataFinUpsertBulk) ClearTxTime() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.ClearTxTime()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *DataFinUpsertBulk) SetTxHash(v string) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateTxHash() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateTxHash()
	})
}

// ClearTxHash clears the value of the "tx_hash" field.
func (u *DataFinUpsertBulk) ClearTxHash() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.ClearTxHash()
	})
}

// SetState sets the "state" field.
func (u *DataFinUpsertBulk) SetState(v string) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateState() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateState()
	})
}

// SetRetries sets the "retries" field.
func (u *DataFinUpsertBulk) SetRetries(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetRetries(v)
	})
}

// AddRetries adds v to the "retries" field.
func (u *DataFinUpsertBulk) AddRetries(v uint32) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.AddRetries(v)
	})
}

// UpdateRetries sets the "retries" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateRetries() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateRetries()
	})
}

// SetRemark sets the "remark" field.
func (u *DataFinUpsertBulk) SetRemark(v string) *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *DataFinUpsertBulk) UpdateRemark() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *DataFinUpsertBulk) ClearRemark() *DataFinUpsertBulk {
	return u.Update(func(s *DataFinUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *DataFinUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DataFinCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DataFinCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DataFinUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
