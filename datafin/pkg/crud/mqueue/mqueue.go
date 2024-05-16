package mqueue

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/datafin/pkg/db"
	"yun.tea/block/bright/datafin/pkg/db/ent"
	"yun.tea/block/bright/datafin/pkg/db/ent/mqueue"
	proto "yun.tea/block/bright/proto/bright/mqueue"
)

func Create(ctx context.Context, in *proto.MqueueReq) (*ent.Mqueue, error) {
	var info *ent.Mqueue
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		c := CreateSet(tx.Mqueue.Create(), in)
		info, err = c.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.MqueueCreate, in *proto.MqueueReq) *ent.MqueueCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.Name != nil {
		c.SetName(*in.Name)
	}
	if in.Remark != nil {
		c.SetRemark(*in.Remark)
	}
	if in.TopicName != nil {
		c.SetTopicName(*in.TopicName)
	}
	return c
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Mqueue, error) {
	var info *ent.Mqueue
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Mqueue.Query().Where(mqueue.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:gocyclo
func setQueryConds(conds *proto.Conds, cli *ent.Client) (*ent.MqueueQuery, error) {
	stm := cli.Mqueue.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(mqueue.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.IDs != nil {
		if conds.GetIDs().GetOp() == cruder.IN {
			var ids []uuid.UUID
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(mqueue.IDIn(ids...))
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.EQ:
			stm.Where(mqueue.Name(conds.GetName().GetValue()))
		default:
			return nil, fmt.Errorf("invalid name field")
		}
	}
	if conds.Remark != nil {
		switch conds.GetRemark().GetOp() {
		case cruder.EQ:
			stm.Where(mqueue.Remark(conds.GetRemark().GetValue()))
		default:
			return nil, fmt.Errorf("invalid description field")
		}
	}
	if conds.TopicName != nil {
		switch conds.GetTopicName().GetOp() {
		case cruder.EQ:
			stm.Where(mqueue.TopicName(conds.GetTopicName().GetValue()))
		default:
			return nil, fmt.Errorf("invalid topic name field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *proto.Conds, offset, limit int) ([]*ent.Mqueue, int, error) {
	var err error
	rows := []*ent.Mqueue{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := setQueryConds(conds, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		rows, err = stm.
			Order(ent.Desc(mqueue.FieldCreatedAt)).
			Offset(offset).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func AllRows(ctx context.Context, desc bool, offset, limit int) ([]*ent.Mqueue, int, error) {
	var err error
	rows := []*ent.Mqueue{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		dtm := cli.Mqueue.Query()
		if desc {
			dtm.Order(ent.Desc(mqueue.FieldCreatedAt))
		} else {
			dtm.Order(ent.Asc(mqueue.FieldCreatedAt))
		}
		rows, err = dtm.
			Offset(offset).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Mqueue, error) {
	var info *ent.Mqueue
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Mqueue.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
