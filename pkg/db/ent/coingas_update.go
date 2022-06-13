// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/coingas"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinGasUpdate is the builder for updating CoinGas entities.
type CoinGasUpdate struct {
	config
	hooks    []Hook
	mutation *CoinGasMutation
}

// Where appends a list predicates to the CoinGasUpdate builder.
func (cgu *CoinGasUpdate) Where(ps ...predicate.CoinGas) *CoinGasUpdate {
	cgu.mutation.Where(ps...)
	return cgu
}

// SetCreatedAt sets the "created_at" field.
func (cgu *CoinGasUpdate) SetCreatedAt(u uint32) *CoinGasUpdate {
	cgu.mutation.ResetCreatedAt()
	cgu.mutation.SetCreatedAt(u)
	return cgu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cgu *CoinGasUpdate) SetNillableCreatedAt(u *uint32) *CoinGasUpdate {
	if u != nil {
		cgu.SetCreatedAt(*u)
	}
	return cgu
}

// AddCreatedAt adds u to the "created_at" field.
func (cgu *CoinGasUpdate) AddCreatedAt(u int32) *CoinGasUpdate {
	cgu.mutation.AddCreatedAt(u)
	return cgu
}

// SetUpdatedAt sets the "updated_at" field.
func (cgu *CoinGasUpdate) SetUpdatedAt(u uint32) *CoinGasUpdate {
	cgu.mutation.ResetUpdatedAt()
	cgu.mutation.SetUpdatedAt(u)
	return cgu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cgu *CoinGasUpdate) AddUpdatedAt(u int32) *CoinGasUpdate {
	cgu.mutation.AddUpdatedAt(u)
	return cgu
}

// SetDeletedAt sets the "deleted_at" field.
func (cgu *CoinGasUpdate) SetDeletedAt(u uint32) *CoinGasUpdate {
	cgu.mutation.ResetDeletedAt()
	cgu.mutation.SetDeletedAt(u)
	return cgu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cgu *CoinGasUpdate) SetNillableDeletedAt(u *uint32) *CoinGasUpdate {
	if u != nil {
		cgu.SetDeletedAt(*u)
	}
	return cgu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cgu *CoinGasUpdate) AddDeletedAt(u int32) *CoinGasUpdate {
	cgu.mutation.AddDeletedAt(u)
	return cgu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cgu *CoinGasUpdate) SetCoinTypeID(u uuid.UUID) *CoinGasUpdate {
	cgu.mutation.SetCoinTypeID(u)
	return cgu
}

// SetGasCoinTypeID sets the "gas_coin_type_id" field.
func (cgu *CoinGasUpdate) SetGasCoinTypeID(u uuid.UUID) *CoinGasUpdate {
	cgu.mutation.SetGasCoinTypeID(u)
	return cgu
}

// SetDepositThresholdLow sets the "deposit_threshold_low" field.
func (cgu *CoinGasUpdate) SetDepositThresholdLow(u uint64) *CoinGasUpdate {
	cgu.mutation.ResetDepositThresholdLow()
	cgu.mutation.SetDepositThresholdLow(u)
	return cgu
}

// AddDepositThresholdLow adds u to the "deposit_threshold_low" field.
func (cgu *CoinGasUpdate) AddDepositThresholdLow(u int64) *CoinGasUpdate {
	cgu.mutation.AddDepositThresholdLow(u)
	return cgu
}

// SetDepositAmount sets the "deposit_amount" field.
func (cgu *CoinGasUpdate) SetDepositAmount(u uint64) *CoinGasUpdate {
	cgu.mutation.ResetDepositAmount()
	cgu.mutation.SetDepositAmount(u)
	return cgu
}

// AddDepositAmount adds u to the "deposit_amount" field.
func (cgu *CoinGasUpdate) AddDepositAmount(u int64) *CoinGasUpdate {
	cgu.mutation.AddDepositAmount(u)
	return cgu
}

// SetOnlineScale sets the "online_scale" field.
func (cgu *CoinGasUpdate) SetOnlineScale(i int32) *CoinGasUpdate {
	cgu.mutation.ResetOnlineScale()
	cgu.mutation.SetOnlineScale(i)
	return cgu
}

// SetNillableOnlineScale sets the "online_scale" field if the given value is not nil.
func (cgu *CoinGasUpdate) SetNillableOnlineScale(i *int32) *CoinGasUpdate {
	if i != nil {
		cgu.SetOnlineScale(*i)
	}
	return cgu
}

// AddOnlineScale adds i to the "online_scale" field.
func (cgu *CoinGasUpdate) AddOnlineScale(i int32) *CoinGasUpdate {
	cgu.mutation.AddOnlineScale(i)
	return cgu
}

// Mutation returns the CoinGasMutation object of the builder.
func (cgu *CoinGasUpdate) Mutation() *CoinGasMutation {
	return cgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cgu *CoinGasUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cgu.defaults(); err != nil {
		return 0, err
	}
	if len(cgu.hooks) == 0 {
		affected, err = cgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinGasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cgu.mutation = mutation
			affected, err = cgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cgu.hooks) - 1; i >= 0; i-- {
			if cgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cgu *CoinGasUpdate) SaveX(ctx context.Context) int {
	affected, err := cgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cgu *CoinGasUpdate) Exec(ctx context.Context) error {
	_, err := cgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cgu *CoinGasUpdate) ExecX(ctx context.Context) {
	if err := cgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cgu *CoinGasUpdate) defaults() error {
	if _, ok := cgu.mutation.UpdatedAt(); !ok {
		if coingas.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coingas.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coingas.UpdateDefaultUpdatedAt()
		cgu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cgu *CoinGasUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coingas.Table,
			Columns: coingas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coingas.FieldID,
			},
		},
	}
	if ps := cgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cgu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldCreatedAt,
		})
	}
	if value, ok := cgu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldCreatedAt,
		})
	}
	if value, ok := cgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldUpdatedAt,
		})
	}
	if value, ok := cgu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldUpdatedAt,
		})
	}
	if value, ok := cgu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldDeletedAt,
		})
	}
	if value, ok := cgu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldDeletedAt,
		})
	}
	if value, ok := cgu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coingas.FieldCoinTypeID,
		})
	}
	if value, ok := cgu.mutation.GasCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coingas.FieldGasCoinTypeID,
		})
	}
	if value, ok := cgu.mutation.DepositThresholdLow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositThresholdLow,
		})
	}
	if value, ok := cgu.mutation.AddedDepositThresholdLow(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositThresholdLow,
		})
	}
	if value, ok := cgu.mutation.DepositAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositAmount,
		})
	}
	if value, ok := cgu.mutation.AddedDepositAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositAmount,
		})
	}
	if value, ok := cgu.mutation.OnlineScale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coingas.FieldOnlineScale,
		})
	}
	if value, ok := cgu.mutation.AddedOnlineScale(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coingas.FieldOnlineScale,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coingas.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CoinGasUpdateOne is the builder for updating a single CoinGas entity.
type CoinGasUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CoinGasMutation
}

// SetCreatedAt sets the "created_at" field.
func (cguo *CoinGasUpdateOne) SetCreatedAt(u uint32) *CoinGasUpdateOne {
	cguo.mutation.ResetCreatedAt()
	cguo.mutation.SetCreatedAt(u)
	return cguo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cguo *CoinGasUpdateOne) SetNillableCreatedAt(u *uint32) *CoinGasUpdateOne {
	if u != nil {
		cguo.SetCreatedAt(*u)
	}
	return cguo
}

// AddCreatedAt adds u to the "created_at" field.
func (cguo *CoinGasUpdateOne) AddCreatedAt(u int32) *CoinGasUpdateOne {
	cguo.mutation.AddCreatedAt(u)
	return cguo
}

// SetUpdatedAt sets the "updated_at" field.
func (cguo *CoinGasUpdateOne) SetUpdatedAt(u uint32) *CoinGasUpdateOne {
	cguo.mutation.ResetUpdatedAt()
	cguo.mutation.SetUpdatedAt(u)
	return cguo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cguo *CoinGasUpdateOne) AddUpdatedAt(u int32) *CoinGasUpdateOne {
	cguo.mutation.AddUpdatedAt(u)
	return cguo
}

// SetDeletedAt sets the "deleted_at" field.
func (cguo *CoinGasUpdateOne) SetDeletedAt(u uint32) *CoinGasUpdateOne {
	cguo.mutation.ResetDeletedAt()
	cguo.mutation.SetDeletedAt(u)
	return cguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cguo *CoinGasUpdateOne) SetNillableDeletedAt(u *uint32) *CoinGasUpdateOne {
	if u != nil {
		cguo.SetDeletedAt(*u)
	}
	return cguo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cguo *CoinGasUpdateOne) AddDeletedAt(u int32) *CoinGasUpdateOne {
	cguo.mutation.AddDeletedAt(u)
	return cguo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cguo *CoinGasUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinGasUpdateOne {
	cguo.mutation.SetCoinTypeID(u)
	return cguo
}

// SetGasCoinTypeID sets the "gas_coin_type_id" field.
func (cguo *CoinGasUpdateOne) SetGasCoinTypeID(u uuid.UUID) *CoinGasUpdateOne {
	cguo.mutation.SetGasCoinTypeID(u)
	return cguo
}

// SetDepositThresholdLow sets the "deposit_threshold_low" field.
func (cguo *CoinGasUpdateOne) SetDepositThresholdLow(u uint64) *CoinGasUpdateOne {
	cguo.mutation.ResetDepositThresholdLow()
	cguo.mutation.SetDepositThresholdLow(u)
	return cguo
}

// AddDepositThresholdLow adds u to the "deposit_threshold_low" field.
func (cguo *CoinGasUpdateOne) AddDepositThresholdLow(u int64) *CoinGasUpdateOne {
	cguo.mutation.AddDepositThresholdLow(u)
	return cguo
}

// SetDepositAmount sets the "deposit_amount" field.
func (cguo *CoinGasUpdateOne) SetDepositAmount(u uint64) *CoinGasUpdateOne {
	cguo.mutation.ResetDepositAmount()
	cguo.mutation.SetDepositAmount(u)
	return cguo
}

// AddDepositAmount adds u to the "deposit_amount" field.
func (cguo *CoinGasUpdateOne) AddDepositAmount(u int64) *CoinGasUpdateOne {
	cguo.mutation.AddDepositAmount(u)
	return cguo
}

// SetOnlineScale sets the "online_scale" field.
func (cguo *CoinGasUpdateOne) SetOnlineScale(i int32) *CoinGasUpdateOne {
	cguo.mutation.ResetOnlineScale()
	cguo.mutation.SetOnlineScale(i)
	return cguo
}

// SetNillableOnlineScale sets the "online_scale" field if the given value is not nil.
func (cguo *CoinGasUpdateOne) SetNillableOnlineScale(i *int32) *CoinGasUpdateOne {
	if i != nil {
		cguo.SetOnlineScale(*i)
	}
	return cguo
}

// AddOnlineScale adds i to the "online_scale" field.
func (cguo *CoinGasUpdateOne) AddOnlineScale(i int32) *CoinGasUpdateOne {
	cguo.mutation.AddOnlineScale(i)
	return cguo
}

// Mutation returns the CoinGasMutation object of the builder.
func (cguo *CoinGasUpdateOne) Mutation() *CoinGasMutation {
	return cguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cguo *CoinGasUpdateOne) Select(field string, fields ...string) *CoinGasUpdateOne {
	cguo.fields = append([]string{field}, fields...)
	return cguo
}

// Save executes the query and returns the updated CoinGas entity.
func (cguo *CoinGasUpdateOne) Save(ctx context.Context) (*CoinGas, error) {
	var (
		err  error
		node *CoinGas
	)
	if err := cguo.defaults(); err != nil {
		return nil, err
	}
	if len(cguo.hooks) == 0 {
		node, err = cguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinGasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cguo.mutation = mutation
			node, err = cguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cguo.hooks) - 1; i >= 0; i-- {
			if cguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cguo *CoinGasUpdateOne) SaveX(ctx context.Context) *CoinGas {
	node, err := cguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cguo *CoinGasUpdateOne) Exec(ctx context.Context) error {
	_, err := cguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cguo *CoinGasUpdateOne) ExecX(ctx context.Context) {
	if err := cguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cguo *CoinGasUpdateOne) defaults() error {
	if _, ok := cguo.mutation.UpdatedAt(); !ok {
		if coingas.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coingas.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coingas.UpdateDefaultUpdatedAt()
		cguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (cguo *CoinGasUpdateOne) sqlSave(ctx context.Context) (_node *CoinGas, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coingas.Table,
			Columns: coingas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coingas.FieldID,
			},
		},
	}
	id, ok := cguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinGas.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coingas.FieldID)
		for _, f := range fields {
			if !coingas.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coingas.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldCreatedAt,
		})
	}
	if value, ok := cguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldCreatedAt,
		})
	}
	if value, ok := cguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldUpdatedAt,
		})
	}
	if value, ok := cguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldUpdatedAt,
		})
	}
	if value, ok := cguo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldDeletedAt,
		})
	}
	if value, ok := cguo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coingas.FieldDeletedAt,
		})
	}
	if value, ok := cguo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coingas.FieldCoinTypeID,
		})
	}
	if value, ok := cguo.mutation.GasCoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coingas.FieldGasCoinTypeID,
		})
	}
	if value, ok := cguo.mutation.DepositThresholdLow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositThresholdLow,
		})
	}
	if value, ok := cguo.mutation.AddedDepositThresholdLow(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositThresholdLow,
		})
	}
	if value, ok := cguo.mutation.DepositAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositAmount,
		})
	}
	if value, ok := cguo.mutation.AddedDepositAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: coingas.FieldDepositAmount,
		})
	}
	if value, ok := cguo.mutation.OnlineScale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coingas.FieldOnlineScale,
		})
	}
	if value, ok := cguo.mutation.AddedOnlineScale(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coingas.FieldOnlineScale,
		})
	}
	_node = &CoinGas{config: cguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coingas.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
