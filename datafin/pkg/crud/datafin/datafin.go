package datafin

import (
	"context"
	"errors"
	"fmt"
	"time"

	"yun.tea/block/bright/datafin/pkg/db/ent/datafin"

	"github.com/google/uuid"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/datafin/pkg/db"
	"yun.tea/block/bright/datafin/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/datafin"
)

func Create(ctx context.Context, in *proto.DataFinReq) (*ent.DataFin, error) {
	var info *ent.DataFin
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.DataFin.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.DataFinCreate, in *proto.DataFinReq) *ent.DataFinCreate {
	c.SetID(uuid.New())
	if in.DataID != nil {
		c.SetDataID(in.GetDataID())
	}
	if in.DataID != nil {
		c.SetDataID(in.GetDataID())
	}
	if in.TopicID != nil {
		c.SetTopicID(in.GetTopicID())
	}
	if in.DataFin != nil {
		c.SetDatafin(in.GetDataFin())
	}
	if in.TxTime != nil {
		c.SetTxTime(in.GetTxTime())
	}
	if in.TxHash != nil {
		c.SetTxHash(in.GetTxHash())
	}
	if in.Retries != nil {
		c.SetRetries(in.GetRetries())
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*proto.DataFinReq) ([]*ent.DataFin, error) {
	var err error
	rows := []*ent.DataFin{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.DataFinCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.DataFin.Create(), info)
		}
		rows, err = tx.DataFin.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *proto.DataFinReq) (*ent.DataFin, error) {
	var err error
	var info *ent.DataFin

	if _, err := uuid.Parse(in.GetDataFinID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.DataFin.UpdateOneID(uuid.MustParse(in.GetDataFinID()))
		u = UpdateSet(u, in)
		info, err = u.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.DataFinUpdateOne, in *proto.DataFinReq) *ent.DataFinUpdateOne {
	if in.DataID != nil {
		u.SetDataID(in.GetDataID())
	}
	if in.TopicID != nil {
		u.SetTopicID(in.GetTopicID())
	}
	if in.DataFin != nil {
		u.SetDatafin(in.GetDataFin())
	}
	if in.TxTime != nil {
		u.SetTxTime(in.GetTxTime())
	}
	if in.TxHash != nil {
		u.SetTxHash(in.GetTxHash())
	}
	if in.Retries != nil {
		u.SetRetries(in.GetRetries())
	}
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.DataFin, error) {
	var info *ent.DataFin
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.DataFin.Query().Where(datafin.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:gocyclo
func setQueryConds(conds *proto.Conds, cli *ent.Client) (*ent.DataFinQuery, error) {
	stm := cli.DataFin.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetDataFinID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetDataFinID().GetValue())
		switch conds.GetDataFinID().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.ID(id))
		default:
			return nil, fmt.Errorf("invalid datafin id field")
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
			stm.Where(datafin.IDIn(ids...))
		}
	}
	if conds.DataID != nil {
		switch conds.GetDataID().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.DataID(conds.GetDataID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid datafin dataid field")
		}
	}
	if conds.TopicID != nil {
		switch conds.GetTopicID().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.TopicID(conds.GetTopicID().GetValue()))
		default:
			return nil, fmt.Errorf("invalid datafin topicid field")
		}
	}
	if conds.DataFin != nil {
		switch conds.GetDataFin().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.Datafin(conds.GetDataFin().GetValue()))
		default:
			return nil, fmt.Errorf("invalid datafin datafin field")
		}
	}
	if conds.TxTime != nil {
		switch conds.GetTxTime().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.TxTime(conds.GetTxTime().GetValue()))
		default:
			return nil, fmt.Errorf("invalid datafin txtime field")
		}
	}
	if conds.TxHash != nil {
		switch conds.GetTxHash().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.TxHash(conds.GetTxHash().GetValue()))
		default:
			return nil, fmt.Errorf("invalid datafin txhash field")
		}
	}
	if conds.Retries != nil {
		switch conds.GetRetries().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.Retries(conds.GetRetries().GetValue()))
		default:
			return nil, fmt.Errorf("invalid datafin retries field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(datafin.State(conds.GetState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid datafin state field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *proto.Conds, offset, limit int) ([]*ent.DataFin, int, error) {
	var err error
	rows := []*ent.DataFin{}
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
			Order(ent.Desc(datafin.FieldCreatedAt)).
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

func AllRows(ctx context.Context, desc bool, offset, limit int) ([]*ent.DataFin, int, error) {
	var err error
	rows := []*ent.DataFin{}
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		dtm := cli.DataFin.Query()
		if desc {
			dtm.Order(ent.Desc(datafin.FieldCreatedAt))
		} else {
			dtm.Order(ent.Asc(datafin.FieldCreatedAt))
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.DataFin, error) {
	var info *ent.DataFin
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.DataFin.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
