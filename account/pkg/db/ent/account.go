// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"yun.tea/block/bright/account/pkg/db/ent/account"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Balance holds the value of the "balance" field.
	Balance string `json:"balance,omitempty"`
	// Enable holds the value of the "enable" field.
	Enable bool `json:"enable,omitempty"`
	// IsRoot holds the value of the "is_root" field.
	IsRoot bool `json:"is_root,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldEnable, account.FieldIsRoot:
			values[i] = new(sql.NullBool)
		case account.FieldCreatedAt, account.FieldUpdatedAt, account.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case account.FieldAddress, account.FieldBalance, account.FieldRemark:
			values[i] = new(sql.NullString)
		case account.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Account", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = uint32(value.Int64)
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = uint32(value.Int64)
			}
		case account.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = uint32(value.Int64)
			}
		case account.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				a.Address = value.String
			}
		case account.FieldBalance:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field balance", values[i])
			} else if value.Valid {
				a.Balance = value.String
			}
		case account.FieldEnable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable", values[i])
			} else if value.Valid {
				a.Enable = value.Bool
			}
		case account.FieldIsRoot:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_root", values[i])
			} else if value.Valid {
				a.IsRoot = value.Bool
			}
		case account.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				a.Remark = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return (&AccountClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(a.Address)
	builder.WriteString(", ")
	builder.WriteString("balance=")
	builder.WriteString(a.Balance)
	builder.WriteString(", ")
	builder.WriteString("enable=")
	builder.WriteString(fmt.Sprintf("%v", a.Enable))
	builder.WriteString(", ")
	builder.WriteString("is_root=")
	builder.WriteString(fmt.Sprintf("%v", a.IsRoot))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(a.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account

func (a Accounts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}