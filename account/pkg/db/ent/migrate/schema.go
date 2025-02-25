// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "address", Type: field.TypeString},
		{Name: "pri_key", Type: field.TypeString},
		{Name: "balance", Type: field.TypeString, Nullable: true},
		{Name: "nonce", Type: field.TypeUint64, Nullable: true, Default: 0},
		{Name: "state", Type: field.TypeString, Default: "AccountUnkonwn"},
		{Name: "is_root", Type: field.TypeBool, Default: false},
		{Name: "remark", Type: field.TypeString, Nullable: true},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "account_deleted_at_pri_key",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[3], AccountsColumns[5]},
			},
			{
				Name:    "account_created_at",
				Unique:  false,
				Columns: []*schema.Column{AccountsColumns[1]},
			},
		},
	}
	// BlockNumsColumns holds the columns for the "block_nums" table.
	BlockNumsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "time_at", Type: field.TypeUint32, Unique: true},
		{Name: "height", Type: field.TypeUint64},
	}
	// BlockNumsTable holds the schema information for the "block_nums" table.
	BlockNumsTable = &schema.Table{
		Name:       "block_nums",
		Columns:    BlockNumsColumns,
		PrimaryKey: []*schema.Column{BlockNumsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "blocknum_time_at",
				Unique:  false,
				Columns: []*schema.Column{BlockNumsColumns[4]},
			},
			{
				Name:    "blocknum_created_at",
				Unique:  false,
				Columns: []*schema.Column{BlockNumsColumns[1]},
			},
		},
	}
	// TxNumsColumns holds the columns for the "tx_nums" table.
	TxNumsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "time_at", Type: field.TypeUint32, Unique: true},
		{Name: "num", Type: field.TypeUint32},
	}
	// TxNumsTable holds the schema information for the "tx_nums" table.
	TxNumsTable = &schema.Table{
		Name:       "tx_nums",
		Columns:    TxNumsColumns,
		PrimaryKey: []*schema.Column{TxNumsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "txnum_time_at",
				Unique:  false,
				Columns: []*schema.Column{TxNumsColumns[4]},
			},
			{
				Name:    "txnum_created_at",
				Unique:  false,
				Columns: []*schema.Column{TxNumsColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		BlockNumsTable,
		TxNumsTable,
	}
)

func init() {
}
