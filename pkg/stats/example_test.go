package stats

import (
	"github.com/nazurdinov95/bank/pkg/bank/types"
	"fmt"
)


func ExampleAverage()  {
	payments := []types.Payment{
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
	result := Average(payments)
	fmt.Println(result)
	// Output: 150
}

func ExampleTotalAmountInCategory()  {
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
		result := TotalAmountInCategory(payments, "авто")
		fmt.Println(result)
		//Output: 100
}
