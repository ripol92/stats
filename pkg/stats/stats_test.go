package stats

import (
	"github.com/ripol92/bank/v2/pkg/types"
	"reflect"
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

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto" : 10,
		"food" : 20,
	}

	second := map[types.Category]types.Money{
		"auto" : 5,
		"food" : 3,
	}

	expected := map[types.Category]types.Money{
		"auto" : -5,
		"food" : -17,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Invalid result")
	}
}

func TestPeriodsDynamic_noCategoryCase(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto" : 10,
		"food" : 20,
	}

	second := map[types.Category]types.Money{
		"food" : 20,
	}

	expected := map[types.Category]types.Money{
		"auto" : -10,
		"food" : 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Invalid result")
	}
}

func TestPeriodsDynamic_noCategoryInFirstCase(t *testing.T) {
	first := map[types.Category]types.Money{
		"food" : 20,
		"auto" : 10,
	}

	second := map[types.Category]types.Money{
		"food" : 25,
		"auto" : 10,
		"mobile" : 5,
	}

	expected := map[types.Category]types.Money{
		"auto" : 0,
		"food" : 5,
		"mobile" : 5,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Error("Invalid result")
	}
}
