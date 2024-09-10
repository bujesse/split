package helpers

import "split/models"

// Function to calculate the person who owes the most and how much they owe
func CalculateOwedDetails(
	expenses []models.Expense,
	settlements []models.Settlement,
) map[string]interface{} {
	result := make(map[string]interface{})
	userOwes := make(map[uint]float64)
	userNames := make(map[uint]string)
	totalSpent := 0.0

	for _, expense := range expenses {
		totalExpense := expense.Amount
		totalSpent += totalExpense

		// TODO: implement FX rates here
		for _, split := range expense.ExpenseSplits {
			userID := split.UserID
			userNames[userID] = split.User.Username

			switch split.SplitType {
			case models.Amt:
				userOwes[userID] += split.SplitValue
			case models.Pct:
				percentageOwed := (split.SplitValue / 100) * totalExpense
				userOwes[userID] += percentageOwed
			}
		}
	}

	// Subtract settlement amounts from each user's owed amount
	for _, settlement := range settlements {
		userOwes[settlement.SettledByID] -= settlement.Amount
	}

	totalOwed := 0.0
	for _, amount := range userOwes {
		totalOwed += amount
	}

	// Net amount is what the user owes minus their share of the total owed by others
	netUserOwes := make(map[uint]float64)
	for userID, amount := range userOwes {
		netUserOwes[userID] = amount - (totalOwed - amount)
	}

	var maxAmountOwed float64
	var userIDWithMax uint

	for userID, amount := range netUserOwes {
		if amount > maxAmountOwed {
			maxAmountOwed = amount
			userIDWithMax = userID
		}
	}

	result["whoOwesMostUserID"] = userIDWithMax
	result["whoOwesMostUsername"] = userNames[userIDWithMax]
	result["maxAmountOwed"] = maxAmountOwed
	result["totalSpent"] = totalSpent

	return result
}
