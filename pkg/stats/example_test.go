package stats

import (
	"reflect"
	"testing"
	"github.com/nazurdinov95/bank/v2/pkg/types"
	"fmt"
)


func ExampleAvg()  {
	payments := []types.Payment{
	{
		ID: 1,
		Amount: 100,
		Category: "авто",
		Status: types.StatusOk,
	},
	{
		ID: 1,
		Amount: 200,
		Category: "аптеки",
		Status: types.StatusFail,
	},
}
	result := Avg(payments)
	fmt.Println(result)
	// Output: 150
}

func ExamplTotaleInCategory()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 100,
			Category: "авто",
		},
		{
			ID: 1,
			Amount: 200,
			Category: "аптеки",
		},
	}
		result := TotalInCategory(payments, "авто")
		fmt.Println(result)
		//Output: 100
}

func TestPeriodDynamic_positive (t *testing. T)  {
	first := map[types.Category]types.Money{
		"auto": 25,
		 "food": 35,
		}
	second :=map[types.Category]types.Money{
		"auto": 45,
		"food": 55,
		}

	
	expected :=map[types.Category] types.Money {
		"auto": 20,
		"food": 20,
	}
	difference :=PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, difference){
		t.Errorf("difference > %v \n expected > %v", difference,expected)
	}
}
