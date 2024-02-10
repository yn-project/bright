package mgr

import (
	"context"

	"github.com/Vigo-Tea/go-ethereum-ant/common"
	crud "yun.tea/block/bright/contract/pkg/crud/contract"
)

func GetContract(ctx context.Context) (*common.Address, error) {
	row, err := crud.Row(ctx)
	if err != nil {
		return nil, err
	}
	addr := common.HexToAddress(row.Address)
	return &addr, nil
}
