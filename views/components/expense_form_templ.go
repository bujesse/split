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

func getContextUserID(ctx context.Context) string {
	if username, ok := ctx.Value("currentUserID").(string); ok {
		return username
	}
	return ""
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
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(getContextUserID(ctx))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 24, Col: 67}
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
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 29, Col: 34}
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
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 32, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" @htmx:after-request=\"$el.reset(); baseModal.close()\" hx-swap=\"none\" x-data=\"init\"><div class=\"mb-4\"><label class=\"input input-bordered flex items-center gap-2\">Name <input x-model=\"Title\" autofocus type=\"text\" id=\"title\" name=\"title\" class=\"grow\" required></label></div><div class=\"mb-4\"><label class=\"input input-bordered flex items-center gap-2\">Amount <input x-model=\"Amount\" type=\"number\" id=\"amount\" name=\"amount\" class=\"grow\" required> <button type=\"button\" x-text=\"currencyCode\" class=\"badge badge-neutral\"></button></label></div><div class=\"mb-4 flex space-x-4\"><div class=\"w-1/2\"><label class=\"form-control w-full max-w-xs\"><div class=\"label\"><span class=\"label-text\">Paid By</span></div><select x-model=\"paidByID\" id=\"paidByID\" name=\"paidByID\" @change=\"splitByID = selectNextOption(document.getElementById(&#39;splitByID&#39;), paidByID)\" class=\"select select-bordered\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, user := range users {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", user.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 60, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(user.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 60, Col: 68}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label></div><div class=\"w-1/2\"><label class=\"form-control w-full max-w-xs\"><div class=\"label\"><span class=\"label-text\">Paid By</span></div><select x-model=\"splitByID\" id=\"splitByID\" name=\"splitByID\" @change=\"paidByID = selectNextOption(document.getElementById(&#39;paidByID&#39;), splitByID)\" class=\"select select-bordered\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, user := range users {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", user.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 78, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(user.Username)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 78, Col: 68}
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></label></div></div><div class=\"mb-4\"><label class=\"input input-bordered flex items-center gap-2\">Split Value <input x-model=\"SplitValue\" type=\"number\" id=\"splitValue\" name=\"SplitValue\" class=\"grow\" required><div class=\"flex rounded-md shadow-sm\"><label class=\"swap swap-rotate bg-neutral text-neutral-content text-center px-4 rounded-full\"><!-- this hidden checkbox controls the state --><input type=\"checkbox\" x-model=\"SplitTypeChecked\" x-effect=\"SplitType = SplitTypeChecked ? &#39;pct&#39; : &#39;amt&#39;\"><div class=\"swap-on\" type=\"button\">%</div><div class=\"swap-off\" type=\"button\">$</div></label> <input x-model=\"SplitType\" type=\"hidden\" name=\"SplitType\"></div></label><div class=\"mt-2 flex\"><span class=\"text-sm\">Equivalent to: $<span x-text=\"calculateEquivalentAmount($data)\"></span></span></div></div><div class=\"mb-4\">")
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"btn btn-warning\" type=\"button\" hx-delete=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs("/api/expenses/" + fmt.Sprintf("%d", expense.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 121, Col: 112}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Delete</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"flex space-x-2\"><button class=\"btn btn-neutral\" type=\"button\" onclick=\"baseModal.close()\">Cancel</button> <button type=\"submit\" class=\"btn btn-primary\">Submit</button></div></div></form><script>\n\t\t\tfunction selectNextOption(selectElement, targetValue) {\n\t\t\t\tconst options = selectElement.options;\n\t\t\t\tlet foundTarget = false;\n\n\t\t\t\tfor (let i = 0; i < options.length; i++) {\n\t\t\t\t\tif (foundTarget) {\n\t\t\t\t\t\tselectElement.selectedIndex = i;\n\t\t\t\t\t\treturn options[i].value;\n\t\t\t\t\t}\n\n\t\t\t\t\tif (options[i].value === targetValue) {\n\t\t\t\t\t\tfoundTarget = true;\n\t\t\t\t\t}\n\t\t\t\t}\n\n\t\t\t\treturn null; // if the target value is the last option or not found\n\t\t\t}\n\t\t\tAlpine.data('init', () => {\n\t\t\t\tconst data = JSON.parse(document.getElementById('expense').textContent)\n\t\t\t\tconst currentUserID = document.getElementById('current-user-id').innerText\n\t\t\t\tconst defaultPaidByID = data?.PaidByID || currentUserID\n\t\t\t\tconst otherUserID = selectNextOption(document.getElementById('paidByID'), currentUserID)\n\t\t\t\tconst expenseSplit = data?.ExpenseSplits.length ? data.ExpenseSplits[0] : {\n\t\t\t\t\tUserID: otherUserID,\n\t\t\t\t\tSplitType: 'pct',\n\t\t\t\t\tSplitValue: 50,\n\t\t\t\t}\n\t\t\t\tconst defaultCurrency = data?.Currency.Code || 'USD'\n\t\t\t\treturn {\n\t\t\t\t\tTitle: null,\n\t\t\t\t\tAmount: null,\n\t\t\t\t\tNotes: null,\n\t\t\t\t\tcurrencyCode: defaultCurrency,\n\t\t\t\t\tCategoryID: null,\n\t\t\t\t\tpaidByID: defaultPaidByID,\n\t\t\t\t\tsplitByID: expenseSplit.UserID,\n\t\t\t\t\tSplitType: expenseSplit.SplitType,\n\t\t\t\t\tSplitValue: expenseSplit.SplitValue,\n\t\t\t\t\tSplitTypeChecked: expenseSplit.SplitType === 'pct',\n\t\t\t\t\tcalculateEquivalentAmount($data) {\n\t\t\t\t\t\tif ($data.SplitType === 'pct') {\n\t\t\t\t\t\t\treturn (($data.SplitValue / 100) * $data.Amount).toFixed(2)\n\t\t\t\t\t\t}\n\t\t\t\t\t\treturn Number($data.SplitValue).toFixed(2)\n\t\t\t\t\t},\n\t\t\t\t\t...data,\n\t\t\t\t}\n\t\t\t})\n\t\t</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Typeahead[T any](name string, model string, placeholder string, items []T) templ.Component {
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
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"typeaheadInit($el)\" class=\"relative\" items=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(items))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 193, Col: 33}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" model=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(model)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 194, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" @keydown.down=\"focusNextItem()\" @keydown.up=\"focusPrevItem()\" @keydown.enter.prevent=\"selectItem(activeItem)\"><label class=\"input input-bordered flex items-center gap-2\"><input type=\"text\" class=\"grow\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(placeholder)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 203, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-model=\"search\" @keydown.escape=\"open = false; activeItem = null\" @focus=\"open = true\" @blur=\"open = false; activeItem = null\"> <input type=\"hidden\" x-model=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(model)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 209, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/expense_form.templ`, Line: 209, Col: 53}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <span x-show=\"search !== &#39;&#39;\" @click=\"search = &#39;&#39;; open = false; activeItem = null\" class=\"badge badge-ghost cursor-pointer\">✕</span></label><!-- <label class=\"form-control w-full\"> --><!-- \t<div class=\"label\"> --><!-- \t\t<span class=\"label-text\">What is your name?</span> --><!-- \t</div> --><!-- \t<input --><!-- \t\ttype=\"text\" --><!-- \t\tclass=\"input input-bordered w-full\" --><!-- \t\tplaceholder={ placeholder } --><!-- \t\tx-model=\"search\" --><!-- \t\t@keydown.escape=\"open = false; activeItem = null\" --><!-- \t\t@focus=\"open = true\" --><!-- \t\t@blur=\"open = false; activeItem = null\" --><!-- \t/> --><!-- \t<input type=\"hidden\" x-model={ model } name={ name }/> --><!-- </label> --><div x-show=\"open\" class=\"absolute w-full\"><ul tabindex=\"0\" class=\"menu bg-base-200 rounded-box w-full block max-h-44 overflow-y-scroll\"><template x-for=\"item in filteredItems\" :key=\"item.ID\"><li><a @click=\"selectItem(item)\" x-text=\"getDisplayName(item)\" :class=\"activeItem === item &amp;&amp; &#39;active&#39;\"></a></li></template></ul></div></div><script>\n\t\tfunction typeaheadInit($el) {\n\t\t\tconst items = JSON.parse($el.getAttribute('items'))\n\t\t\tconst model = $el.getAttribute('model')\n\t\t\treturn {\n\t\t\t\tsearch: '',\n\t\t\t\topen: false,\n\t\t\t\titems: items,\n\t\t\t\tactiveItem: null,\n\t\t\t\tgetDisplayName(item) {\n\t\t\t\t\treturn item.Type + ' > ' + item.Name\n\t\t\t\t},\n\t\t\t\tfilteredItems() {\n\t\t\t\t\treturn this.items.filter(item => {\n\t\t\t\t\t\treturn this.getDisplayName(item).toLowerCase().includes(this.search.toLowerCase())\n\t\t\t\t\t})\n\t\t\t\t},\n\t\t\t\tfocusItem(item) {\n\t\t\t\t\tthis.activeItem = item\n\t\t\t\t},\n\t\t\t\tfocusNextItem() {\n\t\t\t\t\tif (!this.open) this.open = true\n\t\t\t\t\tconst items = this.filteredItems()\n\t\t\t\t\tconst index = items.indexOf(this.activeItem)\n\t\t\t\t\tif (index === -1) {\n\t\t\t\t\t\tthis.activeItem = items[0]\n\t\t\t\t\t} else if (index < items.length - 1) {\n\t\t\t\t\t\tthis.activeItem = items[index + 1]\n\t\t\t\t\t}\n\t\t\t\t},\n\t\t\t\tfocusPrevItem(item) {\n\t\t\t\t\tif (!this.open) this.open = true\n\t\t\t\t\tconst items = this.filteredItems()\n\t\t\t\t\tconst index = items.indexOf(this.activeItem)\n\t\t\t\t\tif (index === -1) {\n\t\t\t\t\t\tthis.activeItem = items[items.length - 1]\n\t\t\t\t\t} else if (index > 0) {\n\t\t\t\t\t\tthis.activeItem = items[index - 1]\n\t\t\t\t\t}\n\t\t\t\t},\n\t\t\t\tselectItem(item) {\n\t\t\t\t\tconst value = item.ID\n\t\t\t\t\tthis[model] = value\n\t\t\t\t\tthis.search = this.getDisplayName(item)\n\t\t\t\t\tthis.open = false\n\t\t\t\t\tthis.activeItem = null\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
