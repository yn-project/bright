package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"github.com/Vigo-Tea/go-ethereum-ant/common"
	"github.com/Vigo-Tea/go-ethereum-ant/crypto"
	"github.com/Vigo-Tea/go-ethereum-ant/ethclient"
	data_fin "yun.tea/block/bright/common/chains/eth/datafin"
	"yun.tea/block/bright/common/utils"
)

func main() {
	cli, err := ethclient.Dial("https://rest.baas.alipay.com/w3/api/a00e36c5/35N604248fA9u3IfW8BeR2RcQ4ZbMfXb")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// privateKeyStr := "9138747718925d94fb6f3ee732bb387dd779375119ce501e95c478c2ff0eeb2e"

	rec, err := cli.TransactionReceipt(context.Background(), common.HexToHash("0xf1f62e46fb3e3530cede697406b7baf88b8d1bc8085cbe56b7eb80535dac01da"))
	fmt.Println(err)
	fmt.Println(utils.PrettyStruct(rec))

	// addr, err := DeployContract2(cli, privateKeyStr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }
	// fmt.Println(addr)

	// GetContractAddr(cli)

	// fmt.Println(common.HexToAddress("2b55ecfbf6150b82c3b6889f426e277fc9f7f2cd").Hex())

	// GetAdminInfos(common.HexToAddress("0xE77E96548B2900767771403489eEe7EB8a9409d6"), cli)

	// AddAdmin(privateKeyStr, cli)
}

func DeployContract2(backend *ethclient.Client, priKey string) (common.Address, error) {
	var contractAddr common.Address

	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		return contractAddr, fmt.Errorf("parse key err: %v", err)
	}

	chainID, err := backend.ChainID(context.Background())
	if err != nil {
		return contractAddr, fmt.Errorf("get eth chainID err: %v", err)
	}
	fmt.Println("chainID:%v", chainID)

	keyedTransctor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	fmt.Println(err)
	_, tx, _, err := data_fin.DeployDataFin(keyedTransctor, backend)
	fmt.Println(err)
	fmt.Println(utils.PrettyStruct(tx))

	time.Sleep(time.Second * 20)

	recp, err := backend.TransactionReceipt(context.Background(), tx.Hash())
	fmt.Println(err)
	fmt.Println(utils.PrettyStruct(recp))
	return recp.ContractAddress, nil
}

func GetContractAddr(backend *ethclient.Client) {

	fmt.Println(common.HexToHash("0x2102f8637f5dbaff62bbf0382929f8134aeb9ccc0abdd584f928fce7f6ce3632"))
	rec, err := backend.TransactionReceipt(context.Background(), common.HexToHash("0x2102f8637f5dbaff62bbf0382929f8134aeb9ccc0abdd584f928fce7f6ce3632"))
	fmt.Println(utils.PrettyStruct(rec))
	fmt.Println(err)
}

func AddAdmin(priKey string, backend *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		fmt.Printf("parse key err: %v\n", err)
	}
	df, err := data_fin.NewDataFin(common.HexToAddress("0xE77E96548B2900767771403489eEe7EB8a9409d6"), backend)
	keyTrans, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(16))
	tx, err := df.AddAdmin(keyTrans, common.HexToAddress("0x22AC1F7bC57B30F385b7d9898Fe4219F0e8B03fB"), "ssssss")
	fmt.Println(err)
	fmt.Println(utils.PrettyStruct(tx))
}

func GetAdminInfos(dfContract common.Address, backend *ethclient.Client) {
	df, err := data_fin.NewDataFin(dfContract, backend)
	fmt.Println(err)
	rets1, rets2, err := df.GetAdminInfos(&bind.CallOpts{
		From: common.HexToAddress("0x7243176257d634A65Ce737c2cd5FAb6B3f65154B"),
	})
	fmt.Println(rets1)
	fmt.Println(rets2)
	fmt.Println(err)
}
