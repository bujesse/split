package helpers

import (
	"math"
	"split/models"
	"split/repositories"
)

// Calculate total spent in base currency and by currency
func calculateTotalSpent(expenses []repositories.ExpenseWithFxRate) (float64, map[string]float64) {
	totalSpentBaseCcy := 0.0
	totalSpentByCurrency := make(map[string]float64)

	for _, expense := range expenses {
		expenseAmountBaseCcy := expense.Amount
		if expense.FxRate != 1 && expense.FxRate != 0 {
			expenseAmountBaseCcy = expenseAmountBaseCcy / expense.FxRate
		}
		totalSpentBaseCcy += expenseAmountBaseCcy
		totalSpentByCurrency[expense.CurrencyCode] += expense.Amount
	}

	return totalSpentBaseCcy, totalSpentByCurrency
}

// Calculate how much each user owes by currency and in base currency, including settlements
func calculateUserOwes(
	expenses []repositories.ExpenseWithFxRate,
	settlements []models.Settlement,
) (map[uint]float64, map[uint]float64, map[uint]map[string]float64, map[uint]string, map[uint]float64) {

	userOwesBaseCcy := make(map[uint]float64)
	userOwesByCurrency := make(map[uint]map[string]float64)
	userNames := make(map[uint]string)

	for _, expense := range expenses {
		for _, split := range expense.ExpenseSplits {
			userID := split.UserID
			userNames[userID] = split.User.Username

			var amountOwed float64
			switch split.SplitType {
			case models.Amt:
				amountOwed = split.SplitValue
			case models.Pct:
				amountOwed = (split.SplitValue / 100) * expense.Amount
			}

			if userOwesByCurrency[split.UserID] == nil {
				userOwesByCurrency[split.UserID] = make(map[string]float64)
			}
			userOwesByCurrency[split.UserID][expense.CurrencyCode] += amountOwed

			amountOwedBaseCcy := amountOwed
			if expense.FxRate != 1 && expense.FxRate != 0 {
				amountOwedBaseCcy = amountOwed / expense.FxRate
			}

			userOwesBaseCcy[userID] += amountOwedBaseCcy
		}
	}

	userSettlements := make(map[uint]float64)
	userOwesBCWithSettlements := DeepCopyMap(userOwesBaseCcy)
	for _, settlement := range settlements {
		userOwesBCWithSettlements[settlement.SettledByID] -= settlement.Amount
		userSettlements[settlement.SettledByID] += settlement.Amount
	}

	return userOwesBCWithSettlements, userOwesBaseCcy, userOwesByCurrency, userNames, userSettlements
}

// Calculate the user who owes the most and their net amount owed
// Net amount is the amount owed by the user minus the amount owed to the user
func calculateMaxOwed(userOwesBaseCcy map[uint]float64) (uint, float64) {
	totalOwed := 0.0
	for _, amount := range userOwesBaseCcy {
		totalOwed += amount
	}

	netUserOwes := make(map[uint]float64)
	for userID, amount := range userOwesBaseCcy {
		netUserOwes[userID] = amount - (totalOwed - amount)
	}

	var maxAmountOwedBaseCcy float64
	var userIDWithMax uint

	for userID, amount := range netUserOwes {
		if amount > maxAmountOwedBaseCcy {
			maxAmountOwedBaseCcy = amount
			userIDWithMax = userID
		}
	}

	maxAmountOwedBaseCcy = math.Round(maxAmountOwedBaseCcy*100) / 100

	return userIDWithMax, maxAmountOwedBaseCcy
}

// Return a map of details about total spendature and owed amounts
func CalculateOwedDetails(
	expenses []repositories.ExpenseWithFxRate,
	settlements []models.Settlement,
) map[string]interface{} {
	result := make(map[string]interface{})

	totalSpentBaseCcy, totalSpentByCurrency := calculateTotalSpent(expenses)

	userOwesBaseCcyWithSettlements, userOwesBaseCcyTotal, userOwesByCurrency, userNames, userSettlements := calculateUserOwes(
		expenses,
		settlements,
	)

	userIDWithMax, maxAmountOwedBaseCcy := calculateMaxOwed(userOwesBaseCcyWithSettlements)

	// Get total amount owed, not including settlements
	// This is used for editing settlements
	_, maxAmountOwedBaseCcyTotal := calculateMaxOwed(
		userOwesBaseCcyTotal,
	)

	if maxAmountOwedBaseCcy == 0 {
		result["whoOwesMostUserID"] = nil
		result["whoOwesMostUsername"] = nil
		result["maxAmountOwed"] = 0
		result["pctOwed"] = 0
	} else {
		result["whoOwesMostUserID"] = userIDWithMax
		result["whoOwesMostUsername"] = userNames[userIDWithMax]
		result["maxAmountOwed"] = maxAmountOwedBaseCcy
		result["pctOwed"] = (maxAmountOwedBaseCcy / totalSpentBaseCcy) * 100
	}

	result["totalSpent"] = totalSpentBaseCcy
	result["totalSpentByCurrency"] = totalSpentByCurrency
	result["userOwesByCurrency"] = userOwesByCurrency

	result["maxAmountOwedBaseCcyTotal"] = maxAmountOwedBaseCcyTotal
	result["userSettlements"] = userSettlements

	return result
}
