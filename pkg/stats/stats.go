package stats

import (
	"github.com/nazurdinov95/bank/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Average(payments []types.Payment) types.Money {
	avg := types.Money(0)
	for _, payment := range payments {
		avg += payment.Amount
	}

	return avg / types.Money (len(payments))
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalAmountInCategory(payments []types.Payment, category types.Category) types.Money {
	totalCategory := types.Money(0)
	for _,payment := range payments {
		if payment.Category == category {

		totalCategory += payment.Amount
		}
	}
	return totalCategory
}