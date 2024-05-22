package blocknum

import (
	"context"
	"time"

	"yun.tea/block/bright/account/pkg/db"
	"yun.tea/block/bright/account/pkg/db/ent"
	blocknument "yun.tea/block/bright/account/pkg/db/ent/blocknum"
)

func UpsertLatestHeight(ctx context.Context, height uint64) (*ent.BlockNum, error) {
	info := &ent.BlockNum{
		TimeAt: uint32(time.Now().Unix()),
		Height: height,
	}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return tx.BlockNum.Create().
			SetTimeAt(info.TimeAt).
			SetHeight(info.Height).
			OnConflict().
			UpdateHeight().Exec(ctx)
	})
	return info, err
}

func Rows(ctx context.Context, latest int) ([]*ent.BlockNum, error) {
	var err error
	rows := []*ent.BlockNum{}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		rows, err = cli.BlockNum.Query().Order(ent.Desc(blocknument.FieldTimeAt)).
			Offset(0).
			Limit(latest).
			All(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
