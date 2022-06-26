package stats

import "github.com/ShokirM/module10_2"

func Avg(payments []types.Payment) types.Money {
	summ := 0
	for _, pay := range payments {
		summ += int(pay.Amount)
	}
	avgAmount := summ / len(payments)
	return types.Money(avgAmount)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	summ := 0
	for _, pay := range payments {
		if pay.Category == category {
			summ += int(pay.Amount)
		}
	}
	return types.Money(summ)
}
