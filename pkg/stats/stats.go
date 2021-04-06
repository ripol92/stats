package stats

import "github.com/ripol92/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	count := 0

	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += payment.Amount
			count++
		}
	}

	return sum / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {
			sum += payment.Amount
		}
	}

	return sum
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money  {
	sumMap := map[types.Category]types.Money{}
	countMap := map[types.Category]types.Money{}
	avgMap := map[types.Category]types.Money{}
	for _, payment := range payments {
		sumMap[payment.Category] += payment.Amount
		countMap[payment.Category] += 1
	}

	for category, money := range sumMap {
		avgMap[category] = money / countMap[category]
	}

	return avgMap
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}

	for category, money := range first {
		result[category] = 0 - money
		for c, m := range second {
			if category == c {
				 result[category] += m
			}
		}
	}

	for category, money := range second {
		inCollection := false
		for c := range first {
			if category == c {
				inCollection = true
				break
			}
		}
		if !inCollection {
			result[category] = money
		}
	}
	
	return result
}