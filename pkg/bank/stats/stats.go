package stats

import "bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
summ := 0
for _, pay := range payments {
	summ += int(pay.Amount)
}
avgAmount := summ / len(payments)
return types.Money(avgAmount)
}



