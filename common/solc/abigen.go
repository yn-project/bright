package solc

import (
	"fmt"

	"github.com/Vigo-Tea/go-ethereum-ant/accounts/abi/bind"
	"yun.tea/block/bright/common/solc/complier"
	"yun.tea/block/bright/common/solc/types"
	"yun.tea/block/bright/common/utils"
)

func GenAPICode(abi, bin, pkg string) (string, error) {
	// Generate the contract binding
	return bind.Bind(
		[]string{pkg},
		[]string{abi},
		[]string{bin},
		[]map[string]string{},
		pkg,
		bind.LangGo,
		make(map[string]string),
		make(map[string]string),
	)
}

func GettypesDefaultParams() *types.Input {
	return &types.Input{
		Language: "Solidity",
		Sources: map[string]types.SourceIn{
			"Default.sol": {Content: "pragma solidity ^0.8.0; contract One { function one() public pure returns (uint) { return 1; } }"},
		},
		Settings: types.Settings{
			Optimizer: types.Optimizer{
				Enabled: true,
				Runs:    200,
			},
			EVMVersion: "byzantium",
			OutputSelection: map[string]map[string][]string{
				"*": {
					"*": []string{
						"abi",
						"evm.bytecode.object",
						"evm.bytecode.sourceMap",
						"evm.deployedBytecode.object",
						"evm.deployedBytecode.sourceMap",
						"evm.methodIdentifiers",
					},
					"": []string{
						"ast",
					},
				},
			},
		},
	}

}

func GetABIAndBIN(fileNmae string, code string, contractName string) (string, string, error) {
	compiler := complier.Solc0_8_0()

	input := GettypesDefaultParams()
	input.Sources = map[string]types.SourceIn{
		fileNmae: {Content: code},
	}

	output, err := compiler.Compile(input)
	if err != nil {
		return "", "", fmt.Errorf("failed to complie code, err %v", fileNmae)
	}

	if _, ok := output.Contracts[fileNmae]; !ok {
		return "", "", fmt.Errorf("cannot find filename %v", fileNmae)
	}

	if _, ok := output.Contracts[fileNmae][contractName]; !ok {
		return "", "", fmt.Errorf("cannot find contractName %v", contractName)
	}

	return utils.PrettyStruct(output.Contracts[fileNmae][contractName].ABI), output.Contracts[fileNmae][contractName].EVM.Bytecode.Object, err
}
