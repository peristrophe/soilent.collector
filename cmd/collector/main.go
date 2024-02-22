package main

import (
	"fmt"

	"github.com/peristrophe/soilent.api/pkg/api.bitflyer.com/v1/board"
	"github.com/peristrophe/soilent.api/pkg/queryparams"
)

func main() {
	board, err := board.Request(queryparams.NewQueryParameters(map[string]any{"product_code": "XRP_JPY"}))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", board)
}
