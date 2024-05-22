package txnum

import (
	"context"
	"time"

	"yun.tea/block/bright/account/pkg/db"
	"yun.tea/block/bright/account/pkg/db/ent"
	txnument "yun.tea/block/bright/account/pkg/db/ent/txnum"
)

func UpsertAddNum(ctx context.Context, num uint32) error {
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return tx.TxNum.Create().
			SetTimeAt(today00Timesamp()).
			SetNum(num).
			OnConflict().
			AddNum(num).Exec(ctx)
	})
}

func today00Timesamp() uint32 {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return uint32(addTime.Unix())
}

func Rows(ctx context.Context, latest int) ([]*ent.TxNum, error) {
	var err error
	rows := []*ent.TxNum{}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		rows, err = cli.TxNum.Query().Order(ent.Desc(txnument.FieldTimeAt)).
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
