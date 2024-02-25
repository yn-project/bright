package main

import (
	"encoding/json"
	"fmt"
)

type Ss struct {
	Name string
	Type string
	Vv   string
}

func main() {
	ss, err := json.Marshal(Ss{Name: "sss", Type: "sdfa", Vv: "sdfadf"})
	fmt.Println(err)
	fmt.Println(string(ss))
}
