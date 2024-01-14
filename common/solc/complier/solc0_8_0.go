package complier

import (
	_ "embed"

	"yun.tea/block/bright/common/solc/types"
)

//go:embed soljson0_8_0.js
var solc0_8_0_bin string

func Solc0_8_0() types.Solc {
	solc, err := types.New(solc0_8_0_bin)
	if err != nil {
		panic(err)
	}
	return solc
}
