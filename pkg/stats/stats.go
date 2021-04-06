package stats

import "github.com/ripol92/bank/pkg/types"

func Avg(payments []types.Payment) types.Money {
	count := len(payments)

	sum := types.Money(0)

	for _, payment := range payments {
		sum += payment.Amount
	}

	return sum / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}

	return sum
}