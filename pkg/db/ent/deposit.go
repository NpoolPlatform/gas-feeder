// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/deposit"
	"github.com/google/uuid"
)

// Deposit is the model entity for the Deposit schema.
type Deposit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	TransactionID uuid.UUID `json:"transaction_id,omitempty"`
	// DepositAmount holds the value of the "deposit_amount" field.
	DepositAmount uint64 `json:"deposit_amount,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deposit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case deposit.FieldCreatedAt, deposit.FieldUpdatedAt, deposit.FieldDeletedAt, deposit.FieldDepositAmount:
			values[i] = new(sql.NullInt64)
		case deposit.FieldID, deposit.FieldAccountID, deposit.FieldTransactionID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Deposit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deposit fields.
func (d *Deposit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deposit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case deposit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = uint32(value.Int64)
			}
		case deposit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = uint32(value.Int64)
			}
		case deposit.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = uint32(value.Int64)
			}
		case deposit.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				d.AccountID = *value
			}
		case deposit.FieldTransactionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value != nil {
				d.TransactionID = *value
			}
		case deposit.FieldDepositAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deposit_amount", values[i])
			} else if value.Valid {
				d.DepositAmount = uint64(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Deposit.
// Note that you need to call Deposit.Unwrap() before calling this method if this Deposit
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deposit) Update() *DepositUpdateOne {
	return (&DepositClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Deposit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deposit) Unwrap() *Deposit {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deposit is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deposit) String() string {
	var builder strings.Builder
	builder.WriteString("Deposit(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatedAt))
	builder.WriteString(", updated_at=")
	builder.WriteString(fmt.Sprintf("%v", d.UpdatedAt))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", d.DeletedAt))
	builder.WriteString(", account_id=")
	builder.WriteString(fmt.Sprintf("%v", d.AccountID))
	builder.WriteString(", transaction_id=")
	builder.WriteString(fmt.Sprintf("%v", d.TransactionID))
	builder.WriteString(", deposit_amount=")
	builder.WriteString(fmt.Sprintf("%v", d.DepositAmount))
	builder.WriteByte(')')
	return builder.String()
}

// Deposits is a parsable slice of Deposit.
type Deposits []*Deposit

func (d Deposits) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
