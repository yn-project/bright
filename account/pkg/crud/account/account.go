package account

import (
	"context"
	"errors"
	"fmt"
	"time"

	"yun.tea/block/bright/account/pkg/db/ent/account"
	"yun.tea/block/bright/account/pkg/sign"

	"github.com/google/uuid"
	"yun.tea/block/bright/account/pkg/db"
	"yun.tea/block/bright/account/pkg/db/ent"
	"yun.tea/block/bright/common/cruder"
	proto "yun.tea/block/bright/proto/bright/account"
)

func Create(ctx context.Context, in *proto.AccountReq) (*ent.Account, error) {
	var info *ent.Account
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Account.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.AccountCreate, in *proto.AccountReq) *ent.AccountCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.Address != nil {
		c.SetAddress(in.GetAddress())
	}
	if in.PriKey != nil {
		fuzzedKey := sign.DefaultFuzzStr(in.GetPriKey())
		c.SetPriKey(fuzzedKey)
	}
	if in.Balance != nil {
		c.SetBalance(in.GetBalance())
	}
	if in.Nonce != nil {
		c.SetNonce(in.GetNonce())
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.IsRoot != nil {
		c.SetIsRoot(in.GetIsRoot())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func Update(ctx context.Context, in *proto.AccountReq) (*ent.Account, error) {
	var err error
	var info *ent.Account

	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, err
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := cli.Account.UpdateOneID(uuid.MustParse(in.GetID()))
		u = UpdateSet(u, in)
		info, err = u.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.AccountUpdateOne, in *proto.AccountReq) *ent.AccountUpdateOne {
	if in.Balance != nil {
		u.SetBalance(in.GetBalance())
	}
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	if in.IsRoot != nil {
		u.SetIsRoot(in.GetIsRoot())
	}
	if in.Balance != nil {
		u.SetBalance(in.GetBalance())
	}
	if in.Nonce != nil {
		u.SetNonce(in.GetNonce())
	}
	if in.Remark != nil {
		u.SetRemark(in.GetRemark())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	var info *ent.Account
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Account.Query().Where(account.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	info.PriKey, err = sign.DefaultDefuzzStr(info.PriKey)
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:gocyclo
func setQueryConds(conds *proto.Conds, cli *ent.Client) (*ent.AccountQuery, error) {
	stm := cli.Account.Query()
	if conds == nil {
		return stm, nil
	}
	if _, err := uuid.Parse(conds.GetID().GetValue()); err == nil {
		id := uuid.MustParse(conds.GetID().GetValue())
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(account.ID(id))
		default:
			return nil, fmt.Errorf("invalid account id field")
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
			stm.Where(account.IDIn(ids...))
		}
	}
	if conds.Address != nil {
		switch conds.GetAddress().GetOp() {
		case cruder.EQ:
			stm.Where(account.Address(conds.GetAddress().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account address field")
		}
	}
	if conds.Balance != nil {
		switch conds.GetBalance().GetOp() {
		case cruder.EQ:
			stm.Where(account.Balance(conds.GetBalance().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account balance field")
		}
	}
	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(account.State(conds.GetState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account state field")
		}
	}
	if conds.IsRoot != nil {
		switch conds.GetIsRoot().GetOp() {
		case cruder.EQ:
			stm.Where(account.IsRoot(conds.GetIsRoot().GetValue()))
		default:
			return nil, fmt.Errorf("invalid account isroot field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *proto.Conds, offset, limit int) ([]*ent.Account, int, error) {
	var err error
	rows := []*ent.Account{}
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
			Order(ent.Desc(account.FieldCreatedAt)).
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

	for _, v := range rows {
		v.PriKey, err = sign.DefaultDefuzzStr(v.PriKey)
		if err != nil {
			v.PriKey = ""
		}
	}
	return rows, total, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Account, error) {
	var info *ent.Account
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Account.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
