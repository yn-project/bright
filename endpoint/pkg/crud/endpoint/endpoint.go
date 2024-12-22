package endpoint

import (
	"context"
	"errors"
	"fmt"
	"time"

	"yun.tea/block/bright/endpoint/pkg/db/ent/endpoint"

	"github.com/google/uuid"
	"yun.tea/block/bright/common/cruder"
	"yun.tea/block/bright/endpoint/pkg/db"
	"yun.tea/block/bright/endpoint/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/endpoint"
)

func Create(ctx context.Context, in *proto.EndpointReq) (*ent.Endpoint, error) {
	var info *ent.Endpoint
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Endpoint.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.EndpointCreate, in *proto.EndpointReq) *ent.EndpointCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Address != nil {
		c.SetAddress(in.GetAddress())
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.RPS != nil {
		c.SetRps(in.GetRPS())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func CreateBulk(ctx context.Context, in []*proto.EndpointReq) ([]*ent.Endpoint, error) {
	var err error
	rows := []*ent.Endpoint{}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.EndpointCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Endpoint.Create(), info)
		}
		rows, err = tx.Endpoint.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *proto.EndpointReq) (*ent.Endpoint, error) {
	var err error
	var info *ent.Endpoint

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.Endpoint.UpdateOneID(uuid.MustParse(in.GetID()))
		u = UpdateSet(u, in)
		info, err = u.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.EndpointUpdateOne, in *proto.EndpointReq) *ent.EndpointUpdateOne {
	if in.Name != nil {
		u.SetName(in.GetName())
	}
	if in.Address != nil {
		u.SetAddress(in.GetAddress())
	}
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	if in.RPS != nil {
		u.SetRps(in.GetRPS())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Endpoint, error) {
	var info *ent.Endpoint
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Endpoint.Query().Where(endpoint.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:gocyclo
func setQueryConds(conds *proto.Conds, cli *ent.Client) (*ent.EndpointQuery, error) {
	stm := cli.Endpoint.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.ID(id))
		default:
			return nil, fmt.Errorf("invalid endpoint id field")
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
			stm.Where(endpoint.IDIn(ids...))
		}
	}
	if conds.Address != nil {
		switch conds.GetAddress().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.Address(conds.GetAddress().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint address field")
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.Name(conds.GetName().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint name field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(endpoint.State(conds.GetState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid endpoint state field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *proto.Conds, offset, limit int) ([]*ent.Endpoint, int, error) {
	var err error
	rows := []*ent.Endpoint{}
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
			Offset(offset).
			Order(ent.Desc(endpoint.FieldCreatedAt)).
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Endpoint, error) {
	var info *ent.Endpoint
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Endpoint.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
