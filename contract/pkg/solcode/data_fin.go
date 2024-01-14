package solcode

import (
	_ "embed"
)

//go:embed data_fin.sol
var SOL_CODE string

const SOL_FILENAME = "data_fin.sol"
const SOL_CONTRACT = "DataFin"
const PKG = "data_fin"
const VERSION = "DataFin-v1.2.0-with-Admin-Owner"
const DESCRIPTION = "DataFin v1.2.0 for yn.tea"
