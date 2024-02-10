package mgr

import (
	"context"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	crud "yun.tea/block/bright/account/pkg/crud/account"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/ctredis"
	"yun.tea/block/bright/common/logger"
	contractmgr "yun.tea/block/bright/contract/pkg/mgr"
	endpointmgr "yun.tea/block/bright/endpoint/pkg/mgr"
	"yun.tea/block/bright/proto/bright/account"
	"yun.tea/block/bright/proto/bright/basetype"
)

const (
	RefreshTime             = time.Minute * 10
	CheckAllAccountTaskTime = time.Minute * 5
	MaxUseAccount           = 100
	MinBalance              = 100000
	CheckAllAccountTaskLock = "check_all_acc_lock"
)

func Maintain(ctx context.Context) {
	for {
		locked, _ := ctredis.TryPubLock(CheckAllAccountTaskLock, CheckAllAccountTaskTime)
		if locked {
			CheckAllAccountState(ctx)
		}
		select {
		case <-time.NewTicker(RefreshTime).C:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func CheckAllAccountState(ctx context.Context) {
	rows, total, err := crud.Rows(ctx, nil, 0, MaxUseAccount)
	if err != nil {
		logger.Sugar().Errorw("CheckAllAccountState", "Msg", "failed to check state of accounts", "Err", err)
		return
	}

	if total == 0 || len(rows) == 0 {
		return
	}

	var availableRootAcc *AccountKey
	availableTreeAccs := []*AccountKey{}
	for _, v := range rows {
		acc, err := GetAccountReport(ctx, v.Address)
		if err != nil {
			logger.Sugar().Errorf("CheckAllAccountState", "Address", v.Address, "Err", err)
		}

		if acc.IsRoot {
			availableRootAcc = &AccountKey{
				Pri: v.PriKey,
				Pub: v.Address,
			}
		}
		if acc.State == basetype.AccountState_AccountAvailable {
			availableTreeAccs = append(availableTreeAccs, &AccountKey{
				Pri: v.PriKey,
				Pub: v.Address,
			})
		}

		id := v.ID.String()
		_, err = crud.Update(ctx, &account.AccountReq{
			ID:      &id,
			Balance: &acc.Balance,
			Nonce:   &acc.Nonce,
			State:   &acc.State,
			IsRoot:  &acc.IsRoot,
			Remark:  &acc.Remark,
		})
		if err != nil {
			logger.Sugar().Errorf("CheckAllAccountState", "Address", v.Address, "Err", err)
		}
	}

	err = GetAccountMGR().SetRootAccount(availableRootAcc)
	if err != nil {
		logger.Sugar().Errorf("CheckAllAccountState", "Err", err)
	}

	err = GetAccountMGR().SetTreeAccounts(availableTreeAccs)
	if err != nil {
		logger.Sugar().Errorf("CheckAllAccountState", "Err", err)
	}

	logger.Sugar().Infow("CheckAllAccountState", "root account", availableRootAcc.Pub)
	treeAccList := []string{}
	for _, v := range availableTreeAccs {
		treeAccList = append(treeAccList, v.Pub)
	}
	logger.Sugar().Infow("CheckAllAccountState", "tree accounts", treeAccList)
}

func GetAllEnableAdmin(ctx context.Context, contractAddr, from common.Address) (string, map[string]bool, error) {
	rootAccount := ""
	treeAccounts := make(map[string]bool)
	err := endpointmgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
		df, err := data_fin.NewDataFin(contractAddr, cli)
		if err != nil {
			return err
		}

		owner, err := df.GetOwner(&bind.CallOpts{
			Pending: true,
			From:    from,
			Context: ctx,
		})
		if err != nil {
			return err
		}
		rootAccount = owner.Hex()

		addrs, infos, err := df.GetAdminInfos(&bind.CallOpts{
			Pending: true,
			From:    from,
			Context: ctx,
		})
		if err != nil {
			return err
		}

		for i := 0; i < len(addrs); i++ {
			if infos[i].Enable {
				treeAccounts[addrs[i].Hex()] = infos[i].Enable
			}
		}
		return nil
	})

	return rootAccount, treeAccounts, err
}

func WithWriteContract(ctx context.Context, needRoot bool, handle func(ctx context.Context, acc *AccountKey, contract *data_fin.DataFin) error) error {
	contractAddr, err := contractmgr.GetContract(ctx)
	if err != nil {
		return err
	}

	amgr := GetAccountMGR()
	var unlock func()
	var acc *AccountKey
	if needRoot {
		acc, unlock, err = amgr.GetRootAccount(ctx)
	} else {
		acc, unlock, err = amgr.GetTreeAccount(ctx)
	}

	if err != nil {
		return err
	}
	defer unlock()

	return endpointmgr.WithClient(ctx, func(ctx context.Context, cli *ethclient.Client) error {
		contract, err := data_fin.NewDataFin(*contractAddr, cli)
		if err != nil {
			return err
		}
		return handle(ctx, acc, contract)
	})
}
