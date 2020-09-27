package stats

import (
	"github.com/nazurdinov95/bank/v2/pkg/types"
	"fmt"
)


func ExampleAvg()  {
	payments := []types.Payment{
	{
		ID: 1,
		Amount: 800,
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
	// Output: 800
}

func ExamplTotaleInCategory()  {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 600,
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
		result := TotalInCategory(payments, "авто")
		fmt.Println(result)
		//Output: 600
}
