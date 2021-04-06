package stats

import (
	"github.com/ripol92/bank/v2/pkg/types"
	"testing"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: types.Money(10), Category: "mobile"},
		{ID: 2, Amount: types.Money(10), Category: "mobile"},
		{ID: 3, Amount: types.Money(10), Category: "mobile"},
		{ID: 4, Amount: types.Money(10), Category: "mobile"},
		{ID: 5, Amount: types.Money(10), Category: "computer"},
		{ID: 6, Amount: types.Money(22), Category: "computer"},
	}

	result := CategoriesAvg(payments)

	if result["mobile"] != 10 {
		t.Errorf("Wrong result! Expected: %v, actual: %v", 10, result["mobile"])
	}

	if result["computer"] != 16 {
		t.Errorf("Wrong result! Expected: %v, actual: %v", 16, result["computer"])
	}

}

func TestCategoriesAvg_emptyPayments(t *testing.T) {
	payments := []types.Payment{}

	result := CategoriesAvg(payments)

	if len(result) > 0 {
		t.Errorf("Wrong result! Expected: %v, actual: %v", 0, len(result))
	}
}
