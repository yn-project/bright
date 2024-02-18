package mgr

import (
	"context"

	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"yun.tea/block/bright/common/logger"
	crud "yun.tea/block/bright/contract/pkg/crud/contract"
	"yun.tea/block/bright/contract/pkg/db"
)

func init() {
	err := db.Init()
	if err != nil {
		logger.Sugar().Error(err)
	}
}

func GetContract(ctx context.Context) (*common.Address, error) {
	row, err := crud.Row(ctx)
	if err != nil {
		return nil, err
	}
	addr := common.HexToAddress(row.Address)
	return &addr, nil
}
