package mgr

import (
	"context"
	"math/big"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/logger"
	"yun.tea/block/bright/common/utils"
	contractmgr "yun.tea/block/bright/contract/pkg/mgr"
	endpointmgr "yun.tea/block/bright/endpoint/pkg/mgr"
	"yun.tea/block/bright/proto/bright/basetype"
)

const (
	// TODO:确定大数据量时最小的耗费
	MinAccountBalance = 50000
	TestContract      = "0x462A090B319ACE4A12a1FF2218aB5416C5e3445E"
)

type AccountReport struct {
	Balance string
	Nonce   uint64
	IsRoot  bool
	State   basetype.AccountState
	Remark  string
}

// 检测账户状态，最终账户非管理员会设置成管理员角色
func GetAccountReport(ctx context.Context, address string) (acc AccountReport, err error) {
	acc = AccountReport{
		Balance: "0",
		State:   basetype.AccountState_AccountAvailable,
		IsRoot:  false,
		Remark:  "可用",
	}
	fromAcc := common.HexToAddress(address)

	defer func() {
		if err != nil && acc.State == basetype.AccountState_AccountAvailable {
			acc.State = basetype.AccountState_AccountUnkonwn
			acc.Remark = "状态未知"
		}
	}()

	_ = endpointmgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
		df, err := data_fin.NewDataFin(common.HexToAddress(TestContract), cli)
		if err != nil {
			return err
		}
		_, err = df.IsAdminEnable(&bind.CallOpts{From: fromAcc}, common.HexToAddress(fromAcc.Hex()))
		if err != nil {
			acc.State = basetype.AccountState_AccountError
			acc.Remark = "账户不可用"
			return err
		}
		return nil
	})

	if err != nil || acc.State != basetype.AccountState_AccountAvailable {
		return
	}

	err = endpointmgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
		_balance, err := cli.BalanceAt(ctx, fromAcc, nil)
		if err != nil {
			return err
		}

		acc.Balance = _balance.String()
		if _balance.Cmp(big.NewInt(MinAccountBalance)) < 0 {
			acc.State = basetype.AccountState_AccountLow
			acc.Remark = "账户余额不足"
		}

		acc.Nonce, err = cli.NonceAt(ctx, fromAcc, nil)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil || acc.State != basetype.AccountState_AccountAvailable {
		return
	}

	contractAddr, err := contractmgr.GetContract(ctx)
	logger.Sugar().Info(utils.PrettyStruct(contractAddr))
	logger.Sugar().Info(utils.PrettyStruct(err))
	if err != nil || len(contractAddr) == 0 {
		acc.State = basetype.AccountState_AccountError
		acc.Remark = "合约不可用"
		return
	}

	rootAcc, treeAccs, err := GetAllEnableAdmin(ctx, *contractAddr, fromAcc)
	if err != nil {
		return
	}

	if rootAcc == address {
		acc.IsRoot = true
	}

	if _, ok := treeAccs[address]; ok {
		acc.State = basetype.AccountState_AccountAvailable
	} else {
		acc.State = basetype.AccountState_AccountError
		acc.Remark = "非合约管理员"
		go func() {
			err = WithWriteContract(ctx, true, func(ctx context.Context, txOpts *bind.TransactOpts, contract *data_fin.DataFin, cli *ethclient.Client) error {
				_, err = contract.AddAdmin(txOpts, common.HexToAddress(address), "auto set to admin")
				return err
			})
		}()
	}
	return
}
