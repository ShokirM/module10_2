package main

import (
	"bank/pkg/bank/stats"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	payments := []types.Payment{
		{
			ID:     256,
			Amount: 100_00,
			Category: "Auto",
		},
		{
			ID:     128,
			Amount: 200_00,
			Category: "Drinks",
		},
		{
			ID:     128,
			Amount: 320_00,
			Category: "Clothes",
		},
		{
			ID:     128,
			Amount: 70_00,
		},
	}
	fmt.Println(stats.Avg(payments))
}
