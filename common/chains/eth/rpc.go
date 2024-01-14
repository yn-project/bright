package eth

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/Vigo-Tea/go-ethereum-ant"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/core/types"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
)

func (ethCli ethClients) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]*types.Log, error) {
	_logs := []types.Log{}

	var err error
	err = ethCli.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		_logs, err = c.FilterLogs(ctx, query)
		if err != nil {
			return false, err
		}
		return false, nil
	})
	logs := make([]*types.Log, len(_logs))
	for i := range _logs {
		logs[i] = &_logs[i]
	}
	return logs, err
}

func (ethCli ethClients) BlockByNumber(ctx context.Context, blockNum *big.Int) (*types.Block, error) {
	var block *types.Block
	var err error
	err = ethCli.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		block, err = c.BlockByNumber(ctx, blockNum)
		return false, err
	})
	return block, err
}

type ContractCreator struct {
	From        common.Address
	BlockNumber uint64
	TxHash      common.Hash
	TxTime      uint64
}

type EthCurrencyMetadata struct {
	Name     string
	Symbol   string
	Decimals uint32
}

func (ethCli ethClients) GetCurrencyMetadata(ctx context.Context, contractAddr string) (*EthCurrencyMetadata, error) {
	return &EthCurrencyMetadata{
		Name:     "Ethereum",
		Symbol:   "ETH",
		Decimals: 18,
	}, nil
}

// ReplaceID replaces the token's ID with the given ID
func (ethCli ethClients) ReplaceID(tokenURI, id string) string {
	_id := fmt.Sprintf("%064s", id)
	return strings.TrimSpace(strings.ReplaceAll(tokenURI, "{id}", _id))
}
