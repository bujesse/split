// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "split/models"
import "split/helpers"
import "split/repositories"

func Stats(expenses []repositories.ExpenseWithFxRate, settlements []models.Settlement) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"stats mx-auto mt-4 flex w-full overflow-x-hidden bg-primary text-primary-content\" x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(helpers.CalculateOwedDetails(expenses, settlements)))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/stats.templ`, Line: 10, Col: 80}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"stat px-3\"><div class=\"stat-title text-primary-content\">Total spent</div><div class=\"stat-value\" x-text=\"FormatAsCurrency(&#39;USD&#39;, totalSpent)\"></div><div class=\"stat-desc mt-2 text-secondary-content\"><span x-show=\"Object.keys(totalSpentByCurrency).length &gt; 1\"><template x-for=\"(amt, ccy) in totalSpentByCurrency\"><span class=\"badge badge-accent badge-sm\" x-text=\"FormatAsCurrency(ccy, amt)\"></span></template></span></div><div class=\"stat-actions\"><button class=\"btn btn-success btn-sm\" onclick=\"baseModal.showModal()\" hx-get=\"/partials/expenses/new\" hx-trigger=\"click\" hx-target=\"#modal-container\">New Expense</button></div></div><div class=\"stat px-2 text-right\"><div class=\"stat-title text-primary-content\" x-text=\"whoOwesMostUsername ? whoOwesMostUsername + &#39; owes&#39; : &#39;Settled Up! 🎉&#39;\"></div><div class=\"stat-value\" x-text=\"FormatAsCurrency(&#39;USD&#39;, maxAmountOwed)\"></div><div class=\"stat-desc mt-2 flex justify-end space-x-1 text-secondary-content\"><span class=\"badge badge-neutral badge-sm\" x-show=\"Math.round(pctOwed) !== 0\" x-text=\"Math.round(pctOwed) + &#39;%&#39;\"></span><!--\n\t\t\t\t<span x-show=\"Object.keys(userOwesByCurrency[whoOwesMostUserID]).length > 1\">\n\t\t\t\t\t<template x-for=\"(amt, ccy) in userOwesByCurrency[whoOwesMostUserID]\">\n\t\t\t\t\t\t<span class=\"badge badge-accent badge-sm\" x-text=\"FormatAsCurrency(ccy, amt)\"></span>\n\t\t\t\t\t</template>\n\t\t\t\t</span>\n\t\t\t\t--></div><div class=\"stat-actions\"><button class=\"btn btn-sm\" x-show=\"whoOwesMostUsername !== null\" onclick=\"baseModal.showModal()\" hx-get=\"/partials/settlements/new\" hx-trigger=\"click\" hx-target=\"#modal-container\">Settle Up</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
