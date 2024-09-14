// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "split/models"
import "fmt"
import "context"

func getExpensePostTarget(expense *models.Expense) string {
	if expense == nil {
		return "/api/expenses"
	}
	return "/api/expenses/" + fmt.Sprintf("%d", expense.ID)
}

func GetContextUserID(ctx context.Context) string {
	if username, ok := ctx.Value("currentUserID").(string); ok {
		return username
	}
	return ""
}

func countryCodeToFlag(code string) string {
	// Generate the flag emoji
	flag := ""
	for _, char := range code {
		flag += string(rune(127397 + char))
	}

	return flag
}

func ExpenseForm(expense *models.Expense, categories []models.Category, currencies []models.Currency, users []models.User) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full max-w-2xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.JSONScript("expense", expense).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span id=\"current-user-id\" class=\"hidden\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(GetContextUserID(ctx))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 34, Col: 67}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><h2 class=\"mb-6 text-2xl font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if expense == nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Add New Expense")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Edit Expense: \"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(expense.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 39, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><form hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(getExpensePostTarget(expense))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 43, Col: 42}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" @htmx:after-request=\"baseModal.close()\" hx-swap=\"none\" x-data=\"init\"><div class=\"mb-4\"><label class=\"input input-bordered flex items-center gap-2\">Name <input x-model=\"Title\" autofocus type=\"text\" id=\"title\" name=\"title\" class=\"grow\" required></label></div><div class=\"mb-4\"><div class=\"join w-full\"><label class=\"input input-bordered flex items-center gap-2 join-item w-7/12 sm:w-full\">Amount <input x-model=\"Amount\" type=\"number\" id=\"amount\" name=\"amount\" class=\"grow\" required></label> <select x-model=\"currencyCode\" name=\"currencyCode\" class=\"select select-bordered join-item w-5/12 sm:w-auto\" @change=\"fxRateUSD = $el.options[$el.selectedIndex].getAttribute(&#39;data-fx-rate-usd&#39;)\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, currency := range currencies {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(currency.Code)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 68, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-fx-rate-usd=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(currency.LatestFxRateUSD))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 69, Col: 69}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(countryCodeToFlag(currency.TwoCharCountryCode))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 70, Col: 56}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(currency.Code)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 70, Col: 74}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></div><div class=\"mt-2 flex\"><span x-show=\"currencyCode !== &#39;USD&#39;\" class=\"text-sm\">Equivalent to: <span x-text=\"FormatAsCurrency(&#39;USD&#39;, Amount / parseFloat(fxRateUSD))\"></span></span></div></div><div class=\"mb-4\"><div class=\"join w-full\"><label class=\"input input-bordered flex items-center gap-2 join-item w-8/12 sm:w-full\"><span>Split Amt</span> <input x-model=\"SplitValue\" type=\"number\" id=\"splitValue\" name=\"SplitValue\" class=\"grow\" required></label> <input x-model=\"SplitType\" value=\"pct\" class=\"join-item btn w-2/12 sm:w-auto\" type=\"radio\" name=\"SplitType\" aria-label=\"%\"> <input x-model=\"SplitType\" value=\"amt\" class=\"join-item btn w-2/12 sm:w-auto\" type=\"radio\" name=\"SplitType\" aria-label=\"$\"></div><div class=\"mt-2 flex\"><span class=\"text-sm\">Equivalent to: <strong x-text=\"FormatAsCurrency(currencyCode, calculateEquivalentAmount($data))\"></strong> <span x-show=\"currencyCode !== &#39;USD&#39;\" class=\"text-sm\">(<span x-text=\"FormatAsCurrency(&#39;USD&#39;, calculateEquivalentAmount($data) / parseFloat(fxRateUSD))\"></span>)</span></span></div></div><div class=\"mb-4 flex space-x-4\"><div class=\"w-1/2\"><label class=\"form-control w-full max-w-xs\"><div class=\"label\"><span class=\"label-text\">Paid By</span></div><select x-model=\"paidByID\" id=\"paidByID\" name=\"paidByID\" @change=\"splitByID = selectNextOption(document.getElementById(&#39;splitByID&#39;), paidByID)\" class=\"select select-bordered\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, user := range users {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", user.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 112, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(user.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 112, Col: 68}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label></div><div class=\"w-1/2\"><label class=\"form-control w-full max-w-xs\"><div class=\"label\"><span class=\"label-text\">Split By</span></div><select x-model=\"splitByID\" id=\"splitByID\" name=\"splitByID\" @change=\"paidByID = selectNextOption(document.getElementById(&#39;paidByID&#39;), splitByID)\" class=\"select select-bordered\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, user := range users {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", user.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 130, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(user.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 130, Col: 68}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label></div></div><div class=\"mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Typeahead("categoryID", "CategoryID", "Search Categories...", categories).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"mb-4\"><textarea x-model=\"Notes\" id=\"notes\" name=\"notes\" class=\"textarea textarea-bordered w-full\" placeholder=\"Notes\"></textarea></div><div class=\"flex justify-between items-center\"><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if expense != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"btn btn-error\" type=\"button\" hx-delete=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs("/api/expenses/" + fmt.Sprintf("%d", expense.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 145, Col: 110}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Delete</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"flex space-x-2\"><button class=\"btn btn-neutral\" type=\"button\" onclick=\"baseModal.close()\">Cancel</button> <button type=\"submit\" class=\"btn btn-primary\">Submit</button></div></div></form><script>\n\t\t\tfunction selectNextOption(selectElement, targetValue) {\n\t\t\t\tconst options = selectElement.options;\n\t\t\t\tlet foundTarget = false;\n\n\t\t\t\tfor (let i = 0; i < options.length; i++) {\n\t\t\t\t\tif (foundTarget) {\n\t\t\t\t\t\tselectElement.selectedIndex = i;\n\t\t\t\t\t\treturn options[i].value;\n\t\t\t\t\t}\n\n\t\t\t\t\tif (options[i].value === targetValue) {\n\t\t\t\t\t\tfoundTarget = true;\n\t\t\t\t\t}\n\t\t\t\t}\n\n\t\t\t\treturn null\n\t\t\t}\n\t\t\tAlpine.data('init', () => {\n\t\t\t\tconst data = JSON.parse(document.getElementById('expense').textContent)\n\t\t\t\tconst currentUserID = document.getElementById('current-user-id').innerText\n\t\t\t\tconst defaultPaidByID = data?.PaidByID || currentUserID\n\t\t\t\tconst otherUserID = selectNextOption(document.getElementById('paidByID'), currentUserID)\n\t\t\t\tconst expenseSplit = data?.ExpenseSplits.length ? data.ExpenseSplits[0] : {\n\t\t\t\t\tUserID: otherUserID,\n\t\t\t\t\tSplitType: 'pct',\n\t\t\t\t\tSplitValue: 50,\n\t\t\t\t}\n\t\t\t\tconst defaultCategoryDisplay = data?.Category ? `${data.Category.Type} > ${data.Category.Name}` : null\n\t\t\t\tconst defaultCurrency = data?.Currency.Code || 'USD'\n\t\t\t\tconst fxRateUSD = data?.Currency.LatestFxRateUSD || null\n\t\t\t\treturn {\n\t\t\t\t\tTitle: null,\n\t\t\t\t\tAmount: null,\n\t\t\t\t\tNotes: null,\n\t\t\t\t\tcurrencyCode: defaultCurrency,\n\t\t\t\t\tfxRateUSD: fxRateUSD,\n\t\t\t\t\tdefaultSearchDisplay: defaultCategoryDisplay,\n\t\t\t\t\tCategoryID: null,\n\t\t\t\t\tpaidByID: defaultPaidByID,\n\t\t\t\t\tsplitByID: expenseSplit.UserID,\n\t\t\t\t\tSplitType: expenseSplit.SplitType,\n\t\t\t\t\tSplitValue: expenseSplit.SplitValue,\n\t\t\t\t\tSplitTypeChecked: expenseSplit.SplitType === 'pct',\n\t\t\t\t\tcalculateEquivalentAmount($data) {\n\t\t\t\t\t\tif ($data.SplitType === 'pct') {\n\t\t\t\t\t\t\treturn ($data.SplitValue / 100) * $data.Amount\n\t\t\t\t\t\t}\n\t\t\t\t\t\treturn $data.SplitValue\n\t\t\t\t\t},\n\t\t\t\t\t...data,\n\t\t\t\t}\n\t\t\t})\n\t\t</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
