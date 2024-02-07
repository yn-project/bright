package mgr

import (
	"context"
	"fmt"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	crud "yun.tea/block/bright/account/pkg/crud/account"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/logger"
	contractmgr "yun.tea/block/bright/contract/pkg/mgr"
	endpointmgr "yun.tea/block/bright/endpoint/pkg/mgr"
	"yun.tea/block/bright/proto/bright/account"
	"yun.tea/block/bright/proto/bright/basetype"
)

const (
	RefreshTime   = time.Minute * 10
	MaxUseAccount = 100
	MinBalance    = 100000
)

func Maintain(ctx context.Context) {
	CheckAllAccountState(ctx)
	for {
		select {
		case <-time.NewTicker(RefreshTime).C:
			CheckAllAccountState(ctx)
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

	contractAddr, err := contractmgr.GetContract()
	if err != nil {
		logger.Sugar().Errorw("CheckAllAccountState", "Msg", "failed to check state of accounts", "Err", err)
		return
	}

	from, err := getFromAccount(ctx)
	if err != nil {
		logger.Sugar().Errorw("CheckAllAccountState", "Msg", "failed to check state of accounts", "Err", err)
		return
	}

	fmt.Println(contractAddr, from)
	rootAccount, treeAccounts, err := GetAllEnableAdmin(ctx, contractAddr, from)
	if err != nil {
		logger.Sugar().Errorw("CheckAllAccountState", "Msg", "failed to check state of accounts", "Err", err)
		return
	}

	var availableRootAcc *AccountKey
	availableTreeAccs := []*AccountKey{}
	for _, v := range rows {
		state := basetype.AccountState_AccountUnkonwn
		if _, ok := treeAccounts[v.Address]; ok && v.Address != rootAccount {
			availableTreeAccs = append(availableTreeAccs, &AccountKey{Pub: v.Address, Pri: v.PriKey})
			state = basetype.AccountState_AccountAvaliable
		} else {
			state = basetype.AccountState_AccountError
		}

		if v.Address == rootAccount {
			availableRootAcc = &AccountKey{
				Pub: v.Address,
				Pri: v.PriKey,
			}
			v.IsRoot = true
			state = basetype.AccountState_AccountAvaliable
		} else {
			v.IsRoot = false
		}
		id := v.ID.String()
		_, err = crud.Update(ctx, &account.AccountReq{
			ID:     &id,
			IsRoot: &v.IsRoot,
			State:  &state,
		})
		if err != nil {
			logger.Sugar().Warnw("CheckAllAccountState", "Err", err)
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

	logger.Sugar().Infow("CheckAllAccountState", "contract", contractAddr.Hex(), "root account", rootAccount)
	treeAccList := []string{}
	for _, v := range availableTreeAccs {
		treeAccList = append(treeAccList, v.Pub)
	}
	logger.Sugar().Infow("CheckAllAccountState", "contract", contractAddr.Hex(), "tree accounts", treeAccList)
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

func getFromAccount(ctx context.Context) (common.Address, error) {
	amgr := GetAccountMGR()
	pubkey, err := amgr.GetRootAccountPub(ctx)
	if err == nil {
		return common.HexToAddress(pubkey), err
	}
	pubkeys, err := amgr.GetTreeAccountPub(ctx)
	if err == nil && len(pubkeys) > 0 {
		return common.HexToAddress(pubkeys[0]), err
	}

	rows, _, err := crud.Rows(ctx, &account.Conds{}, 0, 1)
	if err != nil {
		return common.HexToAddress(""), err
	}

	if len(rows) == 0 {
		return common.HexToAddress(""), fmt.Errorf("have no available account")
	}

	return common.HexToAddress(rows[0].Address), nil
}

func WithWriteContract(ctx context.Context, needRoot bool, handle func(ctx context.Context, acc *AccountKey, contract *data_fin.DataFin) error) error {
	contractAddr, err := contractmgr.GetContract()
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
		contract, err := data_fin.NewDataFin(contractAddr, cli)
		if err != nil {
			return err
		}
		return handle(ctx, acc, contract)
	})
}
