// Code generated by ent, DO NOT EDIT.

package datafin

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"yun.tea/block/bright/datafin/pkg/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// TopicID applies equality check predicate on the "topic_id" field. It's identical to TopicIDEQ.
func TopicID(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopicID), v))
	})
}

// DataID applies equality check predicate on the "data_id" field. It's identical to DataIDEQ.
func DataID(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataID), v))
	})
}

// Datafin applies equality check predicate on the "datafin" field. It's identical to DatafinEQ.
func Datafin(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatafin), v))
	})
}

// TxTime applies equality check predicate on the "tx_time" field. It's identical to TxTimeEQ.
func TxTime(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxTime), v))
	})
}

// TxHash applies equality check predicate on the "tx_hash" field. It's identical to TxHashEQ.
func TxHash(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxHash), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// Retries applies equality check predicate on the "retries" field. It's identical to RetriesEQ.
func Retries(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetries), v))
	})
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// TopicIDEQ applies the EQ predicate on the "topic_id" field.
func TopicIDEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopicID), v))
	})
}

// TopicIDNEQ applies the NEQ predicate on the "topic_id" field.
func TopicIDNEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTopicID), v))
	})
}

// TopicIDIn applies the In predicate on the "topic_id" field.
func TopicIDIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTopicID), v...))
	})
}

// TopicIDNotIn applies the NotIn predicate on the "topic_id" field.
func TopicIDNotIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTopicID), v...))
	})
}

// TopicIDGT applies the GT predicate on the "topic_id" field.
func TopicIDGT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTopicID), v))
	})
}

// TopicIDGTE applies the GTE predicate on the "topic_id" field.
func TopicIDGTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTopicID), v))
	})
}

// TopicIDLT applies the LT predicate on the "topic_id" field.
func TopicIDLT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTopicID), v))
	})
}

// TopicIDLTE applies the LTE predicate on the "topic_id" field.
func TopicIDLTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTopicID), v))
	})
}

// TopicIDContains applies the Contains predicate on the "topic_id" field.
func TopicIDContains(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTopicID), v))
	})
}

// TopicIDHasPrefix applies the HasPrefix predicate on the "topic_id" field.
func TopicIDHasPrefix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTopicID), v))
	})
}

// TopicIDHasSuffix applies the HasSuffix predicate on the "topic_id" field.
func TopicIDHasSuffix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTopicID), v))
	})
}

// TopicIDEqualFold applies the EqualFold predicate on the "topic_id" field.
func TopicIDEqualFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTopicID), v))
	})
}

// TopicIDContainsFold applies the ContainsFold predicate on the "topic_id" field.
func TopicIDContainsFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTopicID), v))
	})
}

// DataIDEQ applies the EQ predicate on the "data_id" field.
func DataIDEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDataID), v))
	})
}

// DataIDNEQ applies the NEQ predicate on the "data_id" field.
func DataIDNEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDataID), v))
	})
}

// DataIDIn applies the In predicate on the "data_id" field.
func DataIDIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDataID), v...))
	})
}

// DataIDNotIn applies the NotIn predicate on the "data_id" field.
func DataIDNotIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDataID), v...))
	})
}

// DataIDGT applies the GT predicate on the "data_id" field.
func DataIDGT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDataID), v))
	})
}

// DataIDGTE applies the GTE predicate on the "data_id" field.
func DataIDGTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDataID), v))
	})
}

// DataIDLT applies the LT predicate on the "data_id" field.
func DataIDLT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDataID), v))
	})
}

// DataIDLTE applies the LTE predicate on the "data_id" field.
func DataIDLTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDataID), v))
	})
}

// DataIDContains applies the Contains predicate on the "data_id" field.
func DataIDContains(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDataID), v))
	})
}

// DataIDHasPrefix applies the HasPrefix predicate on the "data_id" field.
func DataIDHasPrefix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDataID), v))
	})
}

// DataIDHasSuffix applies the HasSuffix predicate on the "data_id" field.
func DataIDHasSuffix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDataID), v))
	})
}

// DataIDEqualFold applies the EqualFold predicate on the "data_id" field.
func DataIDEqualFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDataID), v))
	})
}

// DataIDContainsFold applies the ContainsFold predicate on the "data_id" field.
func DataIDContainsFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDataID), v))
	})
}

// DatafinEQ applies the EQ predicate on the "datafin" field.
func DatafinEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatafin), v))
	})
}

// DatafinNEQ applies the NEQ predicate on the "datafin" field.
func DatafinNEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDatafin), v))
	})
}

// DatafinIn applies the In predicate on the "datafin" field.
func DatafinIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDatafin), v...))
	})
}

// DatafinNotIn applies the NotIn predicate on the "datafin" field.
func DatafinNotIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDatafin), v...))
	})
}

// DatafinGT applies the GT predicate on the "datafin" field.
func DatafinGT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDatafin), v))
	})
}

// DatafinGTE applies the GTE predicate on the "datafin" field.
func DatafinGTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDatafin), v))
	})
}

// DatafinLT applies the LT predicate on the "datafin" field.
func DatafinLT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDatafin), v))
	})
}

// DatafinLTE applies the LTE predicate on the "datafin" field.
func DatafinLTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDatafin), v))
	})
}

// DatafinContains applies the Contains predicate on the "datafin" field.
func DatafinContains(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDatafin), v))
	})
}

// DatafinHasPrefix applies the HasPrefix predicate on the "datafin" field.
func DatafinHasPrefix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDatafin), v))
	})
}

// DatafinHasSuffix applies the HasSuffix predicate on the "datafin" field.
func DatafinHasSuffix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDatafin), v))
	})
}

// DatafinEqualFold applies the EqualFold predicate on the "datafin" field.
func DatafinEqualFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDatafin), v))
	})
}

// DatafinContainsFold applies the ContainsFold predicate on the "datafin" field.
func DatafinContainsFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDatafin), v))
	})
}

// TxTimeEQ applies the EQ predicate on the "tx_time" field.
func TxTimeEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxTime), v))
	})
}

// TxTimeNEQ applies the NEQ predicate on the "tx_time" field.
func TxTimeNEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTxTime), v))
	})
}

// TxTimeIn applies the In predicate on the "tx_time" field.
func TxTimeIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTxTime), v...))
	})
}

// TxTimeNotIn applies the NotIn predicate on the "tx_time" field.
func TxTimeNotIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTxTime), v...))
	})
}

// TxTimeGT applies the GT predicate on the "tx_time" field.
func TxTimeGT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTxTime), v))
	})
}

// TxTimeGTE applies the GTE predicate on the "tx_time" field.
func TxTimeGTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTxTime), v))
	})
}

// TxTimeLT applies the LT predicate on the "tx_time" field.
func TxTimeLT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTxTime), v))
	})
}

// TxTimeLTE applies the LTE predicate on the "tx_time" field.
func TxTimeLTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTxTime), v))
	})
}

// TxTimeIsNil applies the IsNil predicate on the "tx_time" field.
func TxTimeIsNil() predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTxTime)))
	})
}

// TxTimeNotNil applies the NotNil predicate on the "tx_time" field.
func TxTimeNotNil() predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTxTime)))
	})
}

// TxHashEQ applies the EQ predicate on the "tx_hash" field.
func TxHashEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxHash), v))
	})
}

// TxHashNEQ applies the NEQ predicate on the "tx_hash" field.
func TxHashNEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTxHash), v))
	})
}

// TxHashIn applies the In predicate on the "tx_hash" field.
func TxHashIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTxHash), v...))
	})
}

// TxHashNotIn applies the NotIn predicate on the "tx_hash" field.
func TxHashNotIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTxHash), v...))
	})
}

// TxHashGT applies the GT predicate on the "tx_hash" field.
func TxHashGT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTxHash), v))
	})
}

// TxHashGTE applies the GTE predicate on the "tx_hash" field.
func TxHashGTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTxHash), v))
	})
}

// TxHashLT applies the LT predicate on the "tx_hash" field.
func TxHashLT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTxHash), v))
	})
}

// TxHashLTE applies the LTE predicate on the "tx_hash" field.
func TxHashLTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTxHash), v))
	})
}

// TxHashContains applies the Contains predicate on the "tx_hash" field.
func TxHashContains(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTxHash), v))
	})
}

// TxHashHasPrefix applies the HasPrefix predicate on the "tx_hash" field.
func TxHashHasPrefix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTxHash), v))
	})
}

// TxHashHasSuffix applies the HasSuffix predicate on the "tx_hash" field.
func TxHashHasSuffix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTxHash), v))
	})
}

// TxHashIsNil applies the IsNil predicate on the "tx_hash" field.
func TxHashIsNil() predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTxHash)))
	})
}

// TxHashNotNil applies the NotNil predicate on the "tx_hash" field.
func TxHashNotNil() predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTxHash)))
	})
}

// TxHashEqualFold applies the EqualFold predicate on the "tx_hash" field.
func TxHashEqualFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTxHash), v))
	})
}

// TxHashContainsFold applies the ContainsFold predicate on the "tx_hash" field.
func TxHashContainsFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTxHash), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// RetriesEQ applies the EQ predicate on the "retries" field.
func RetriesEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetries), v))
	})
}

// RetriesNEQ applies the NEQ predicate on the "retries" field.
func RetriesNEQ(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRetries), v))
	})
}

// RetriesIn applies the In predicate on the "retries" field.
func RetriesIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRetries), v...))
	})
}

// RetriesNotIn applies the NotIn predicate on the "retries" field.
func RetriesNotIn(vs ...uint32) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRetries), v...))
	})
}

// RetriesGT applies the GT predicate on the "retries" field.
func RetriesGT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRetries), v))
	})
}

// RetriesGTE applies the GTE predicate on the "retries" field.
func RetriesGTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRetries), v))
	})
}

// RetriesLT applies the LT predicate on the "retries" field.
func RetriesLT(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRetries), v))
	})
}

// RetriesLTE applies the LTE predicate on the "retries" field.
func RetriesLTE(v uint32) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRetries), v))
	})
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemark), v...))
	})
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.DataFin {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemark), v...))
	})
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRemark)))
	})
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRemark)))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DataFin) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DataFin) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DataFin) predicate.DataFin {
	return predicate.DataFin(func(s *sql.Selector) {
		p(s.Not())
	})
}