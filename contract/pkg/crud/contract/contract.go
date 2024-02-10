package contract

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"yun.tea/block/bright/contract/pkg/db"
	"yun.tea/block/bright/contract/pkg/db/ent"
	proto "yun.tea/block/bright/proto/bright/contract"
)

func Create(ctx context.Context, in *proto.ContractReq) (*ent.Contract, error) {
	var info *ent.Contract
	var err error
	if in == nil {
		return nil, errors.New("input is nil")
	}
	db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		rows, err := tx.Contract.Query().All(ctx)
		if err != nil {
			return err
		}
		for _, row := range rows {
			_, err := tx.Contract.UpdateOneID(row.ID).SetDeletedAt(uint32(time.Now().Unix())).Save(ctx)
			if err != nil {
				return err
			}
		}
		c := CreateSet(tx.Contract.Create(), in)
		info, err = c.Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateSet(c *ent.ContractCreate, in *proto.ContractReq) *ent.ContractCreate {
	if in.ID != nil {
		c.SetID(uuid.New())
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Address != nil {
		c.SetAddress(in.GetAddress())
	}
	if in.Version != nil {
		c.SetVersion(in.GetVersion())
	}
	if in.Remark != nil {
		c.SetRemark(in.GetRemark())
	}
	return c
}

func Row(ctx context.Context) (*ent.Contract, error) {
	var info *ent.Contract
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Contract.Query().Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Contract, error) {
	var info *ent.Contract
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Contract.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
