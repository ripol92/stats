package stats

import "github.com/ripol92/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	count := 0

	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status != 'StatusFail' {
			sum += payment.Amount
			count++
		}
	}

	return sum / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category && payment.Status != 'StatusFail' {
			sum += payment.Amount
		}
	}

	return sum
}