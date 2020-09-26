package stats

import (
	"github.com/nazurdinov95/bank/pkg/types"
	"fmt"
)


func ExampleAvg()  {
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
	result := Avg(payments)
	fmt.Println(result)
	// Output: 150
}

func ExampleTotaleInCategory()  {
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
		result := TotaleInCategory(payments, "авто")
		fmt.Println(result)
		//Output: 100
}
