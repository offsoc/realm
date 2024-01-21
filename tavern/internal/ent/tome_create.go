// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"realm.pub/tavern/internal/ent/file"
	"realm.pub/tavern/internal/ent/tome"
)

// TomeCreate is the builder for creating a Tome entity.
type TomeCreate struct {
	config
	mutation *TomeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tc *TomeCreate) SetCreatedAt(t time.Time) *TomeCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TomeCreate) SetNillableCreatedAt(t *time.Time) *TomeCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (tc *TomeCreate) SetLastModifiedAt(t time.Time) *TomeCreate {
	tc.mutation.SetLastModifiedAt(t)
	return tc
}

// SetNillableLastModifiedAt sets the "last_modified_at" field if the given value is not nil.
func (tc *TomeCreate) SetNillableLastModifiedAt(t *time.Time) *TomeCreate {
	if t != nil {
		tc.SetLastModifiedAt(*t)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TomeCreate) SetName(s string) *TomeCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDescription sets the "description" field.
func (tc *TomeCreate) SetDescription(s string) *TomeCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetParamDefs sets the "param_defs" field.
func (tc *TomeCreate) SetParamDefs(s string) *TomeCreate {
	tc.mutation.SetParamDefs(s)
	return tc
}

// SetNillableParamDefs sets the "param_defs" field if the given value is not nil.
func (tc *TomeCreate) SetNillableParamDefs(s *string) *TomeCreate {
	if s != nil {
		tc.SetParamDefs(*s)
	}
	return tc
}

// SetHash sets the "hash" field.
func (tc *TomeCreate) SetHash(s string) *TomeCreate {
	tc.mutation.SetHash(s)
	return tc
}

// SetEldritch sets the "eldritch" field.
func (tc *TomeCreate) SetEldritch(s string) *TomeCreate {
	tc.mutation.SetEldritch(s)
	return tc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (tc *TomeCreate) AddFileIDs(ids ...int) *TomeCreate {
	tc.mutation.AddFileIDs(ids...)
	return tc
}

// AddFiles adds the "files" edges to the File entity.
func (tc *TomeCreate) AddFiles(f ...*File) *TomeCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tc.AddFileIDs(ids...)
}

// Mutation returns the TomeMutation object of the builder.
func (tc *TomeCreate) Mutation() *TomeMutation {
	return tc.mutation
}

// Save creates the Tome in the database.
func (tc *TomeCreate) Save(ctx context.Context) (*Tome, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TomeCreate) SaveX(ctx context.Context) *Tome {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TomeCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TomeCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TomeCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if tome.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized tome.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := tome.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.LastModifiedAt(); !ok {
		if tome.DefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized tome.DefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := tome.DefaultLastModifiedAt()
		tc.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TomeCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Tome.created_at"`)}
	}
	if _, ok := tc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "last_modified_at", err: errors.New(`ent: missing required field "Tome.last_modified_at"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Tome.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := tome.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tome.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Tome.description"`)}
	}
	if _, ok := tc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Tome.hash"`)}
	}
	if v, ok := tc.mutation.Hash(); ok {
		if err := tome.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "Tome.hash": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Eldritch(); !ok {
		return &ValidationError{Name: "eldritch", err: errors.New(`ent: missing required field "Tome.eldritch"`)}
	}
	return nil
}

func (tc *TomeCreate) sqlSave(ctx context.Context) (*Tome, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TomeCreate) createSpec() (*Tome, *sqlgraph.CreateSpec) {
	var (
		_node = &Tome{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tome.Table, sqlgraph.NewFieldSpec(tome.FieldID, field.TypeInt))
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tome.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.LastModifiedAt(); ok {
		_spec.SetField(tome.FieldLastModifiedAt, field.TypeTime, value)
		_node.LastModifiedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(tome.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(tome.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.ParamDefs(); ok {
		_spec.SetField(tome.FieldParamDefs, field.TypeString, value)
		_node.ParamDefs = value
	}
	if value, ok := tc.mutation.Hash(); ok {
		_spec.SetField(tome.FieldHash, field.TypeString, value)
		_node.Hash = value
	}
	if value, ok := tc.mutation.Eldritch(); ok {
		_spec.SetField(tome.FieldEldritch, field.TypeString, value)
		_node.Eldritch = value
	}
	if nodes := tc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tome.FilesTable,
			Columns: tome.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tome.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TomeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tc *TomeCreate) OnConflict(opts ...sql.ConflictOption) *TomeUpsertOne {
	tc.conflict = opts
	return &TomeUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tome.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TomeCreate) OnConflictColumns(columns ...string) *TomeUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TomeUpsertOne{
		create: tc,
	}
}

type (
	// TomeUpsertOne is the builder for "upsert"-ing
	//  one Tome node.
	TomeUpsertOne struct {
		create *TomeCreate
	}

	// TomeUpsert is the "OnConflict" setter.
	TomeUpsert struct {
		*sql.UpdateSet
	}
)

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *TomeUpsert) SetLastModifiedAt(v time.Time) *TomeUpsert {
	u.Set(tome.FieldLastModifiedAt, v)
	return u
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *TomeUpsert) UpdateLastModifiedAt() *TomeUpsert {
	u.SetExcluded(tome.FieldLastModifiedAt)
	return u
}

// SetName sets the "name" field.
func (u *TomeUpsert) SetName(v string) *TomeUpsert {
	u.Set(tome.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TomeUpsert) UpdateName() *TomeUpsert {
	u.SetExcluded(tome.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *TomeUpsert) SetDescription(v string) *TomeUpsert {
	u.Set(tome.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TomeUpsert) UpdateDescription() *TomeUpsert {
	u.SetExcluded(tome.FieldDescription)
	return u
}

// SetParamDefs sets the "param_defs" field.
func (u *TomeUpsert) SetParamDefs(v string) *TomeUpsert {
	u.Set(tome.FieldParamDefs, v)
	return u
}

// UpdateParamDefs sets the "param_defs" field to the value that was provided on create.
func (u *TomeUpsert) UpdateParamDefs() *TomeUpsert {
	u.SetExcluded(tome.FieldParamDefs)
	return u
}

// ClearParamDefs clears the value of the "param_defs" field.
func (u *TomeUpsert) ClearParamDefs() *TomeUpsert {
	u.SetNull(tome.FieldParamDefs)
	return u
}

// SetHash sets the "hash" field.
func (u *TomeUpsert) SetHash(v string) *TomeUpsert {
	u.Set(tome.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TomeUpsert) UpdateHash() *TomeUpsert {
	u.SetExcluded(tome.FieldHash)
	return u
}

// SetEldritch sets the "eldritch" field.
func (u *TomeUpsert) SetEldritch(v string) *TomeUpsert {
	u.Set(tome.FieldEldritch, v)
	return u
}

// UpdateEldritch sets the "eldritch" field to the value that was provided on create.
func (u *TomeUpsert) UpdateEldritch() *TomeUpsert {
	u.SetExcluded(tome.FieldEldritch)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Tome.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TomeUpsertOne) UpdateNewValues() *TomeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(tome.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tome.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TomeUpsertOne) Ignore() *TomeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TomeUpsertOne) DoNothing() *TomeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TomeCreate.OnConflict
// documentation for more info.
func (u *TomeUpsertOne) Update(set func(*TomeUpsert)) *TomeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TomeUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *TomeUpsertOne) SetLastModifiedAt(v time.Time) *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *TomeUpsertOne) UpdateLastModifiedAt() *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetName sets the "name" field.
func (u *TomeUpsertOne) SetName(v string) *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TomeUpsertOne) UpdateName() *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *TomeUpsertOne) SetDescription(v string) *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TomeUpsertOne) UpdateDescription() *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateDescription()
	})
}

// SetParamDefs sets the "param_defs" field.
func (u *TomeUpsertOne) SetParamDefs(v string) *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.SetParamDefs(v)
	})
}

// UpdateParamDefs sets the "param_defs" field to the value that was provided on create.
func (u *TomeUpsertOne) UpdateParamDefs() *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateParamDefs()
	})
}

// ClearParamDefs clears the value of the "param_defs" field.
func (u *TomeUpsertOne) ClearParamDefs() *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.ClearParamDefs()
	})
}

// SetHash sets the "hash" field.
func (u *TomeUpsertOne) SetHash(v string) *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TomeUpsertOne) UpdateHash() *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateHash()
	})
}

// SetEldritch sets the "eldritch" field.
func (u *TomeUpsertOne) SetEldritch(v string) *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.SetEldritch(v)
	})
}

// UpdateEldritch sets the "eldritch" field to the value that was provided on create.
func (u *TomeUpsertOne) UpdateEldritch() *TomeUpsertOne {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateEldritch()
	})
}

// Exec executes the query.
func (u *TomeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TomeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TomeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TomeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TomeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TomeCreateBulk is the builder for creating many Tome entities in bulk.
type TomeCreateBulk struct {
	config
	err      error
	builders []*TomeCreate
	conflict []sql.ConflictOption
}

// Save creates the Tome entities in the database.
func (tcb *TomeCreateBulk) Save(ctx context.Context) ([]*Tome, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tome, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TomeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TomeCreateBulk) SaveX(ctx context.Context) []*Tome {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TomeCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TomeCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Tome.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TomeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tcb *TomeCreateBulk) OnConflict(opts ...sql.ConflictOption) *TomeUpsertBulk {
	tcb.conflict = opts
	return &TomeUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Tome.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TomeCreateBulk) OnConflictColumns(columns ...string) *TomeUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TomeUpsertBulk{
		create: tcb,
	}
}

// TomeUpsertBulk is the builder for "upsert"-ing
// a bulk of Tome nodes.
type TomeUpsertBulk struct {
	create *TomeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Tome.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TomeUpsertBulk) UpdateNewValues() *TomeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(tome.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Tome.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TomeUpsertBulk) Ignore() *TomeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TomeUpsertBulk) DoNothing() *TomeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TomeCreateBulk.OnConflict
// documentation for more info.
func (u *TomeUpsertBulk) Update(set func(*TomeUpsert)) *TomeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TomeUpsert{UpdateSet: update})
	}))
	return u
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (u *TomeUpsertBulk) SetLastModifiedAt(v time.Time) *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.SetLastModifiedAt(v)
	})
}

// UpdateLastModifiedAt sets the "last_modified_at" field to the value that was provided on create.
func (u *TomeUpsertBulk) UpdateLastModifiedAt() *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateLastModifiedAt()
	})
}

// SetName sets the "name" field.
func (u *TomeUpsertBulk) SetName(v string) *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TomeUpsertBulk) UpdateName() *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *TomeUpsertBulk) SetDescription(v string) *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *TomeUpsertBulk) UpdateDescription() *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateDescription()
	})
}

// SetParamDefs sets the "param_defs" field.
func (u *TomeUpsertBulk) SetParamDefs(v string) *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.SetParamDefs(v)
	})
}

// UpdateParamDefs sets the "param_defs" field to the value that was provided on create.
func (u *TomeUpsertBulk) UpdateParamDefs() *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateParamDefs()
	})
}

// ClearParamDefs clears the value of the "param_defs" field.
func (u *TomeUpsertBulk) ClearParamDefs() *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.ClearParamDefs()
	})
}

// SetHash sets the "hash" field.
func (u *TomeUpsertBulk) SetHash(v string) *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *TomeUpsertBulk) UpdateHash() *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateHash()
	})
}

// SetEldritch sets the "eldritch" field.
func (u *TomeUpsertBulk) SetEldritch(v string) *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.SetEldritch(v)
	})
}

// UpdateEldritch sets the "eldritch" field to the value that was provided on create.
func (u *TomeUpsertBulk) UpdateEldritch() *TomeUpsertBulk {
	return u.Update(func(s *TomeUpsert) {
		s.UpdateEldritch()
	})
}

// Exec executes the query.
func (u *TomeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TomeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TomeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TomeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
