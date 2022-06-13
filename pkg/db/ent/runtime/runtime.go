// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/coingas"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/deposit"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/schema"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	coingasMixin := schema.CoinGas{}.Mixin()
	coingas.Policy = privacy.NewPolicies(coingasMixin[0], schema.CoinGas{})
	coingas.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := coingas.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	coingasMixinFields0 := coingasMixin[0].Fields()
	_ = coingasMixinFields0
	coingasFields := schema.CoinGas{}.Fields()
	_ = coingasFields
	// coingasDescCreatedAt is the schema descriptor for created_at field.
	coingasDescCreatedAt := coingasMixinFields0[0].Descriptor()
	// coingas.DefaultCreatedAt holds the default value on creation for the created_at field.
	coingas.DefaultCreatedAt = coingasDescCreatedAt.Default.(func() uint32)
	// coingasDescUpdatedAt is the schema descriptor for updated_at field.
	coingasDescUpdatedAt := coingasMixinFields0[1].Descriptor()
	// coingas.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coingas.DefaultUpdatedAt = coingasDescUpdatedAt.Default.(func() uint32)
	// coingas.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coingas.UpdateDefaultUpdatedAt = coingasDescUpdatedAt.UpdateDefault.(func() uint32)
	// coingasDescDeletedAt is the schema descriptor for deleted_at field.
	coingasDescDeletedAt := coingasMixinFields0[2].Descriptor()
	// coingas.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	coingas.DefaultDeletedAt = coingasDescDeletedAt.Default.(func() uint32)
	// coingasDescOnlineScale is the schema descriptor for online_scale field.
	coingasDescOnlineScale := coingasFields[5].Descriptor()
	// coingas.DefaultOnlineScale holds the default value on creation for the online_scale field.
	coingas.DefaultOnlineScale = coingasDescOnlineScale.Default.(int32)
	// coingasDescID is the schema descriptor for id field.
	coingasDescID := coingasFields[0].Descriptor()
	// coingas.DefaultID holds the default value on creation for the id field.
	coingas.DefaultID = coingasDescID.Default.(func() uuid.UUID)
	depositMixin := schema.Deposit{}.Mixin()
	deposit.Policy = privacy.NewPolicies(depositMixin[0], schema.Deposit{})
	deposit.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := deposit.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	depositMixinFields0 := depositMixin[0].Fields()
	_ = depositMixinFields0
	depositFields := schema.Deposit{}.Fields()
	_ = depositFields
	// depositDescCreatedAt is the schema descriptor for created_at field.
	depositDescCreatedAt := depositMixinFields0[0].Descriptor()
	// deposit.DefaultCreatedAt holds the default value on creation for the created_at field.
	deposit.DefaultCreatedAt = depositDescCreatedAt.Default.(func() uint32)
	// depositDescUpdatedAt is the schema descriptor for updated_at field.
	depositDescUpdatedAt := depositMixinFields0[1].Descriptor()
	// deposit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deposit.DefaultUpdatedAt = depositDescUpdatedAt.Default.(func() uint32)
	// deposit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deposit.UpdateDefaultUpdatedAt = depositDescUpdatedAt.UpdateDefault.(func() uint32)
	// depositDescDeletedAt is the schema descriptor for deleted_at field.
	depositDescDeletedAt := depositMixinFields0[2].Descriptor()
	// deposit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	deposit.DefaultDeletedAt = depositDescDeletedAt.Default.(func() uint32)
	// depositDescID is the schema descriptor for id field.
	depositDescID := depositFields[0].Descriptor()
	// deposit.DefaultID holds the default value on creation for the id field.
	deposit.DefaultID = depositDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)
