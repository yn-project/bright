// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/ent/datafin"
)

// DataFin is the model entity for the DataFin schema.
type DataFin struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// TopicID holds the value of the "topic_id" field.
	TopicID string `json:"topic_id,omitempty"`
	// DataID holds the value of the "data_id" field.
	DataID string `json:"data_id,omitempty"`
	// Datafin holds the value of the "datafin" field.
	Datafin string `json:"datafin,omitempty"`
	// TxTime holds the value of the "tx_time" field.
	TxTime uint32 `json:"tx_time,omitempty"`
	// TxHash holds the value of the "tx_hash" field.
	TxHash string `json:"tx_hash,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Retries holds the value of the "retries" field.
	Retries uint32 `json:"retries,omitempty"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DataFin) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case datafin.FieldCreatedAt, datafin.FieldUpdatedAt, datafin.FieldDeletedAt, datafin.FieldTxTime, datafin.FieldRetries:
			values[i] = new(sql.NullInt64)
		case datafin.FieldTopicID, datafin.FieldDataID, datafin.FieldDatafin, datafin.FieldTxHash, datafin.FieldState, datafin.FieldRemark:
			values[i] = new(sql.NullString)
		case datafin.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DataFin", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DataFin fields.
func (df *DataFin) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case datafin.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				df.ID = *value
			}
		case datafin.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				df.CreatedAt = uint32(value.Int64)
			}
		case datafin.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				df.UpdatedAt = uint32(value.Int64)
			}
		case datafin.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				df.DeletedAt = uint32(value.Int64)
			}
		case datafin.FieldTopicID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field topic_id", values[i])
			} else if value.Valid {
				df.TopicID = value.String
			}
		case datafin.FieldDataID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data_id", values[i])
			} else if value.Valid {
				df.DataID = value.String
			}
		case datafin.FieldDatafin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field datafin", values[i])
			} else if value.Valid {
				df.Datafin = value.String
			}
		case datafin.FieldTxTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tx_time", values[i])
			} else if value.Valid {
				df.TxTime = uint32(value.Int64)
			}
		case datafin.FieldTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tx_hash", values[i])
			} else if value.Valid {
				df.TxHash = value.String
			}
		case datafin.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				df.State = value.String
			}
		case datafin.FieldRetries:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field retries", values[i])
			} else if value.Valid {
				df.Retries = uint32(value.Int64)
			}
		case datafin.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				df.Remark = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DataFin.
// Note that you need to call DataFin.Unwrap() before calling this method if this DataFin
// was returned from a transaction, and the transaction was committed or rolled back.
func (df *DataFin) Update() *DataFinUpdateOne {
	return (&DataFinClient{config: df.config}).UpdateOne(df)
}

// Unwrap unwraps the DataFin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (df *DataFin) Unwrap() *DataFin {
	_tx, ok := df.config.driver.(*txDriver)
	if !ok {
		panic("ent: DataFin is not a transactional entity")
	}
	df.config.driver = _tx.drv
	return df
}

// String implements the fmt.Stringer.
func (df *DataFin) String() string {
	var builder strings.Builder
	builder.WriteString("DataFin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", df.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", df.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", df.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", df.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("topic_id=")
	builder.WriteString(df.TopicID)
	builder.WriteString(", ")
	builder.WriteString("data_id=")
	builder.WriteString(df.DataID)
	builder.WriteString(", ")
	builder.WriteString("datafin=")
	builder.WriteString(df.Datafin)
	builder.WriteString(", ")
	builder.WriteString("tx_time=")
	builder.WriteString(fmt.Sprintf("%v", df.TxTime))
	builder.WriteString(", ")
	builder.WriteString("tx_hash=")
	builder.WriteString(df.TxHash)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(df.State)
	builder.WriteString(", ")
	builder.WriteString("retries=")
	builder.WriteString(fmt.Sprintf("%v", df.Retries))
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(df.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// DataFins is a parsable slice of DataFin.
type DataFins []*DataFin

func (df DataFins) config(cfg config) {
	for _i := range df {
		df[_i].config = cfg
	}
}
