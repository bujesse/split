// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "split/models"
import "fmt"

func getTotalSpent(expenses []models.Expense) string {
	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}
	return fmt.Sprintf("%.2f", total)
}

// Function to calculate the person who owes the most and how much they owe
func calculateOwedDetails(expenses []models.Expense) map[string]interface{} {
	result := make(map[string]interface{})
	userOwes := make(map[string]float64)
	totalSpent := 0.0

	for _, expense := range expenses {
		totalExpense := expense.Amount
		totalSpent += totalExpense

		// TODO: implement FX rates here
		for _, split := range expense.ExpenseSplits {
			switch split.SplitType {
			case models.Amt:
				userOwes[split.User.Username] += split.SplitValue
			case models.Pct:
				percentageOwed := (split.SplitValue / 100) * totalExpense
				userOwes[split.User.Username] += percentageOwed
			}
		}
	}

	totalOwed := 0.0
	for _, amount := range userOwes {
		totalOwed += amount
	}

	// Net amount is what the user owes minus their share of the total owed by others
	netUserOwes := make(map[string]float64)
	for username, amount := range userOwes {
		netUserOwes[username] = amount - (totalOwed - amount)
	}

	var maxAmountOwed float64
	var usernameWithMax string

	for username, amount := range netUserOwes {
		if amount > maxAmountOwed {
			maxAmountOwed = amount
			usernameWithMax = username
		}
	}

	result["whoOwesMost"] = usernameWithMax
	result["maxAmountOwed"] = maxAmountOwed
	result["totalSpent"] = totalSpent

	return result
}

func Stats(expenses []models.Expense) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex mt-4 stats bg-primary text-primary-content w-full sm:w-1/4 mx-auto\" x-data=\"{ owedDetails: JSON.parse($el.getAttribute(&#39;data-owed-details&#39;)) }\" data-owed-details=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(calculateOwedDetails(expenses)))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/partials/stats.templ`, Line: 68, Col: 70}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"stat px-3\"><div class=\"stat-title text-primary-content\">Total spent</div><div class=\"stat-value\" x-text=\"FormatAsCurrency(&#39;USD&#39;, owedDetails.totalSpent)\"></div><div class=\"stat-actions\"><button class=\"btn btn-sm btn-success\" onclick=\"baseModal.showModal()\" hx-get=\"/partials/expenses/new\" hx-trigger=\"click\" hx-target=\"#modal-container\">New Expense</button></div></div><div class=\"stat px-2 text-right\"><div class=\"stat-title text-primary-content\" x-text=\"owedDetails.whoOwesMost ? owedDetails.whoOwesMost + &#39; owes&#39; : &#39;Settled Up! 🎉&#39;\"></div><div class=\"stat-value\" x-text=\"FormatAsCurrency(&#39;USD&#39;, owedDetails.maxAmountOwed)\"></div><div class=\"stat-actions\"><button class=\"btn btn-sm\">Settle Up</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
