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
	"yun.tea/block/bright/contract/pkg/db/ent/contract"
)

// ContractCreate is the builder for creating a Contract entity.
type ContractCreate struct {
	config
	mutation *ContractMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *ContractCreate) SetCreatedAt(u uint32) *ContractCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableCreatedAt(u *uint32) *ContractCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ContractCreate) SetUpdatedAt(u uint32) *ContractCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableUpdatedAt(u *uint32) *ContractCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ContractCreate) SetDeletedAt(u uint32) *ContractCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableDeletedAt(u *uint32) *ContractCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ContractCreate) SetName(s string) *ContractCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetAddress sets the "address" field.
func (cc *ContractCreate) SetAddress(s string) *ContractCreate {
	cc.mutation.SetAddress(s)
	return cc
}

// SetRemark sets the "remark" field.
func (cc *ContractCreate) SetRemark(s string) *ContractCreate {
	cc.mutation.SetRemark(s)
	return cc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cc *ContractCreate) SetNillableRemark(s *string) *ContractCreate {
	if s != nil {
		cc.SetRemark(*s)
	}
	return cc
}

// SetVersion sets the "version" field.
func (cc *ContractCreate) SetVersion(s string) *ContractCreate {
	cc.mutation.SetVersion(s)
	return cc
}

// SetID sets the "id" field.
func (cc *ContractCreate) SetID(u uuid.UUID) *ContractCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ContractCreate) SetNillableID(u *uuid.UUID) *ContractCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the ContractMutation object of the builder.
func (cc *ContractCreate) Mutation() *ContractMutation {
	return cc.mutation
}

// Save creates the Contract in the database.
func (cc *ContractCreate) Save(ctx context.Context) (*Contract, error) {
	var (
		err  error
		node *Contract
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Contract)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ContractMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContractCreate) SaveX(ctx context.Context) *Contract {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContractCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContractCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContractCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if contract.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := contract.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if contract.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := contract.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		if contract.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := contract.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if contract.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized contract.DefaultID (forgotten import ent/runtime?)")
		}
		v := contract.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContractCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Contract.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Contract.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Contract.deleted_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Contract.name"`)}
	}
	if _, ok := cc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Contract.address"`)}
	}
	if _, ok := cc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Contract.version"`)}
	}
	return nil
}

func (cc *ContractCreate) sqlSave(ctx context.Context) (*Contract, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ContractCreate) createSpec() (*Contract, *sqlgraph.CreateSpec) {
	var (
		_node = &Contract{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: contract.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: contract.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := cc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := cc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldVersion,
		})
		_node.Version = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Contract.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContractUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cc *ContractCreate) OnConflict(opts ...sql.ConflictOption) *ContractUpsertOne {
	cc.conflict = opts
	return &ContractUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ContractCreate) OnConflictColumns(columns ...string) *ContractUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ContractUpsertOne{
		create: cc,
	}
}

type (
	// ContractUpsertOne is the builder for "upsert"-ing
	//  one Contract node.
	ContractUpsertOne struct {
		create *ContractCreate
	}

	// ContractUpsert is the "OnConflict" setter.
	ContractUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ContractUpsert) SetCreatedAt(v uint32) *ContractUpsert {
	u.Set(contract.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ContractUpsert) UpdateCreatedAt() *ContractUpsert {
	u.SetExcluded(contract.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ContractUpsert) AddCreatedAt(v uint32) *ContractUpsert {
	u.Add(contract.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContractUpsert) SetUpdatedAt(v uint32) *ContractUpsert {
	u.Set(contract.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContractUpsert) UpdateUpdatedAt() *ContractUpsert {
	u.SetExcluded(contract.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ContractUpsert) AddUpdatedAt(v uint32) *ContractUpsert {
	u.Add(contract.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ContractUpsert) SetDeletedAt(v uint32) *ContractUpsert {
	u.Set(contract.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ContractUpsert) UpdateDeletedAt() *ContractUpsert {
	u.SetExcluded(contract.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ContractUpsert) AddDeletedAt(v uint32) *ContractUpsert {
	u.Add(contract.FieldDeletedAt, v)
	return u
}

// SetName sets the "name" field.
func (u *ContractUpsert) SetName(v string) *ContractUpsert {
	u.Set(contract.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ContractUpsert) UpdateName() *ContractUpsert {
	u.SetExcluded(contract.FieldName)
	return u
}

// SetAddress sets the "address" field.
func (u *ContractUpsert) SetAddress(v string) *ContractUpsert {
	u.Set(contract.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *ContractUpsert) UpdateAddress() *ContractUpsert {
	u.SetExcluded(contract.FieldAddress)
	return u
}

// SetRemark sets the "remark" field.
func (u *ContractUpsert) SetRemark(v string) *ContractUpsert {
	u.Set(contract.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *ContractUpsert) UpdateRemark() *ContractUpsert {
	u.SetExcluded(contract.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *ContractUpsert) ClearRemark() *ContractUpsert {
	u.SetNull(contract.FieldRemark)
	return u
}

// SetVersion sets the "version" field.
func (u *ContractUpsert) SetVersion(v string) *ContractUpsert {
	u.Set(contract.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ContractUpsert) UpdateVersion() *ContractUpsert {
	u.SetExcluded(contract.FieldVersion)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contract.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContractUpsertOne) UpdateNewValues() *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(contract.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Contract.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ContractUpsertOne) Ignore() *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContractUpsertOne) DoNothing() *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContractCreate.OnConflict
// documentation for more info.
func (u *ContractUpsertOne) Update(set func(*ContractUpsert)) *ContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContractUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ContractUpsertOne) SetCreatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ContractUpsertOne) AddCreatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateCreatedAt() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContractUpsertOne) SetUpdatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ContractUpsertOne) AddUpdatedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateUpdatedAt() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ContractUpsertOne) SetDeletedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ContractUpsertOne) AddDeletedAt(v uint32) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateDeletedAt() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *ContractUpsertOne) SetName(v string) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateName() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateName()
	})
}

// SetAddress sets the "address" field.
func (u *ContractUpsertOne) SetAddress(v string) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateAddress() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateAddress()
	})
}

// SetRemark sets the "remark" field.
func (u *ContractUpsertOne) SetRemark(v string) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateRemark() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *ContractUpsertOne) ClearRemark() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.ClearRemark()
	})
}

// SetVersion sets the "version" field.
func (u *ContractUpsertOne) SetVersion(v string) *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ContractUpsertOne) UpdateVersion() *ContractUpsertOne {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateVersion()
	})
}

// Exec executes the query.
func (u *ContractUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContractCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContractUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ContractUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ContractUpsertOne.ID is not supported by MySQL driver. Use ContractUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ContractUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ContractCreateBulk is the builder for creating many Contract entities in bulk.
type ContractCreateBulk struct {
	config
	builders []*ContractCreate
	conflict []sql.ConflictOption
}

// Save creates the Contract entities in the database.
func (ccb *ContractCreateBulk) Save(ctx context.Context) ([]*Contract, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Contract, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContractMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContractCreateBulk) SaveX(ctx context.Context) []*Contract {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContractCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContractCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Contract.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContractUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ccb *ContractCreateBulk) OnConflict(opts ...sql.ConflictOption) *ContractUpsertBulk {
	ccb.conflict = opts
	return &ContractUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ContractCreateBulk) OnConflictColumns(columns ...string) *ContractUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ContractUpsertBulk{
		create: ccb,
	}
}

// ContractUpsertBulk is the builder for "upsert"-ing
// a bulk of Contract nodes.
type ContractUpsertBulk struct {
	create *ContractCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contract.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContractUpsertBulk) UpdateNewValues() *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(contract.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Contract.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ContractUpsertBulk) Ignore() *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContractUpsertBulk) DoNothing() *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContractCreateBulk.OnConflict
// documentation for more info.
func (u *ContractUpsertBulk) Update(set func(*ContractUpsert)) *ContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContractUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ContractUpsertBulk) SetCreatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ContractUpsertBulk) AddCreatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateCreatedAt() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ContractUpsertBulk) SetUpdatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ContractUpsertBulk) AddUpdatedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateUpdatedAt() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ContractUpsertBulk) SetDeletedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ContractUpsertBulk) AddDeletedAt(v uint32) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateDeletedAt() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *ContractUpsertBulk) SetName(v string) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateName() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateName()
	})
}

// SetAddress sets the "address" field.
func (u *ContractUpsertBulk) SetAddress(v string) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateAddress() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateAddress()
	})
}

// SetRemark sets the "remark" field.
func (u *ContractUpsertBulk) SetRemark(v string) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateRemark() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *ContractUpsertBulk) ClearRemark() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.ClearRemark()
	})
}

// SetVersion sets the "version" field.
func (u *ContractUpsertBulk) SetVersion(v string) *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *ContractUpsertBulk) UpdateVersion() *ContractUpsertBulk {
	return u.Update(func(s *ContractUpsert) {
		s.UpdateVersion()
	})
}

// Exec executes the query.
func (u *ContractUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ContractCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContractCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContractUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
