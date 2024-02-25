package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(hex.DecodeString("123abc"))
}
