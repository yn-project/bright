// Code generated by ent, DO NOT EDIT.

package ent

import (
	"yun.tea/block/bright/account/pkg/db/ent/account"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 1)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   account.Table,
			Columns: account.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: account.FieldID,
			},
		},
		Type: "Account",
		Fields: map[string]*sqlgraph.FieldSpec{
			account.FieldCreatedAt: {Type: field.TypeUint32, Column: account.FieldCreatedAt},
			account.FieldUpdatedAt: {Type: field.TypeUint32, Column: account.FieldUpdatedAt},
			account.FieldDeletedAt: {Type: field.TypeUint32, Column: account.FieldDeletedAt},
			account.FieldAddress:   {Type: field.TypeString, Column: account.FieldAddress},
			account.FieldBalance:   {Type: field.TypeString, Column: account.FieldBalance},
			account.FieldEnable:    {Type: field.TypeBool, Column: account.FieldEnable},
			account.FieldIsRoot:    {Type: field.TypeBool, Column: account.FieldIsRoot},
			account.FieldRemark:    {Type: field.TypeString, Column: account.FieldRemark},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *AccountQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AccountQuery builder.
func (aq *AccountQuery) Filter() *AccountFilter {
	return &AccountFilter{config: aq.config, predicateAdder: aq}
}

// addPredicate implements the predicateAdder interface.
func (m *AccountMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AccountMutation builder.
func (m *AccountMutation) Filter() *AccountFilter {
	return &AccountFilter{config: m.config, predicateAdder: m}
}

// AccountFilter provides a generic filtering capability at runtime for AccountQuery.
type AccountFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AccountFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AccountFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(account.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AccountFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(account.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AccountFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(account.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AccountFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(account.FieldDeletedAt))
}

// WhereAddress applies the entql string predicate on the address field.
func (f *AccountFilter) WhereAddress(p entql.StringP) {
	f.Where(p.Field(account.FieldAddress))
}

// WhereBalance applies the entql string predicate on the balance field.
func (f *AccountFilter) WhereBalance(p entql.StringP) {
	f.Where(p.Field(account.FieldBalance))
}

// WhereEnable applies the entql bool predicate on the enable field.
func (f *AccountFilter) WhereEnable(p entql.BoolP) {
	f.Where(p.Field(account.FieldEnable))
}

// WhereIsRoot applies the entql bool predicate on the is_root field.
func (f *AccountFilter) WhereIsRoot(p entql.BoolP) {
	f.Where(p.Field(account.FieldIsRoot))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *AccountFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(account.FieldRemark))
}