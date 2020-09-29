package stats

import (
	"github.com/nazurdinov95/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	avg := types.Money(0)
	var i = 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			avg += payment.Amount
			i++
		}
	}

	return avg /types.Money(i)
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	totalCategory := types.Money(0)
	for _,payment := range payments {
		if payment.Category == category && payment.Status != types.StatusFail {

		totalCategory += payment.Amount
		}
	}
	return totalCategory
}

//FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

// CategoriesAvg рассчитывает среднюю сумму платежа по категорию. 
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	mp := make(map[types.Category]types.Money)
	for _, v := range payments {
		if _, er := mp[v.Category]; er {
			continue
		}
		filtered := FilterByCategory(payments, v.Category)
		mp[v.Category]=Avg(filtered)
	}
	return mp
	}