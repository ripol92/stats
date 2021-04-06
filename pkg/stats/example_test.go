package stats

import (
	"github.com/ripol92/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{ID: 1, Amount: types.Money(10)},
		{ID: 2, Amount: types.Money(20)},
		{ID: 3, Amount: types.Money(30)},
	}

	avg := Avg(payments)
	fmt.Println(avg)

	// Output: 20
}

func ExampleTotalInCategory() {
	category := types.Category("card")

	payments := []types.Payment {
		{ID: 1, Amount: types.Money(10), Category: types.Category("overdraft")},
		{ID: 2, Amount: types.Money(20), Category: types.Category("card")},
		{ID: 3, Amount: types.Money(30), Category: types.Category("card")},
	}

	total := TotalInCategory(payments, category)
	fmt.Println(total)

	// Output: 50
}
