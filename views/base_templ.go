// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"split/views/components"
)

func Base() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"root\" x-data><div class=\"flex h-screen\"><main class=\"min-h-screen w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Nav().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"main-content\" class=\"container mx-auto px-4 pb-24 sm:w-1/3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Modal().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"toast toast-top\" x-show=\"$store.global.toasts.length &gt; 0\" x-transition:enter.delay.50ms x-cloak><template x-for=\"toast in $store.global.toasts\"><div id=\"error-toast\" class=\"alert max-w-xs overflow-auto p-4\" :class=\"{\n\t\t\t\t\t\t\t\t&#39;alert-error&#39;: toast.toastType === &#39;error&#39;,\n\t\t\t\t\t\t\t\t&#39;alert-success&#39;: toast.toastType === &#39;success&#39;\n\t\t\t\t\t\t\t}\" x-text=\"toast.toastText\"></div></template></div></div></body><script>\n\t\t\tconst MAX_DIGITS = 5\n\t\t\tconst ERROR_TOAST_TIMEOUT = 5000 // milliseconds\n\t\t\tconst TOAST_TYPES = {\n\t\t\t\tERROR: 'error',\n\t\t\t\tSUCCESS: 'success',\n\t\t\t}\n\n\t\t\tfunction FormatAsCurrency(currency, amount, truncate = true) {\n\t\t\t\tif (!currency) {\n\t\t\t\t\tcurrency = 'USD';\n\t\t\t\t}\n\n\t\t\t\tconst totalDigitsWithCents = amount.toFixed(2).replace('.', '').length\n\n\t\t\t\tlet formatter = new Intl.NumberFormat('en-US', {\n\t\t\t\t\tstyle: 'currency',\n\t\t\t\t\tcurrency: currency,\n\t\t\t\t\tmaximumFractionDigits: truncate && totalDigitsWithCents > MAX_DIGITS ? 0 : 2,\n\t\t\t\t});\n\t\t\t\treturn formatter.format(amount);\n\t\t\t}\n\n\t\t\tfunction LocalizeDate(isoString) {\n\t\t\t\treturn new Date(isoString).toLocaleDateString(undefined, {\n\t\t\t\t\tmonth: 'short',\n\t\t\t\t\tday: '2-digit'\n\t\t\t\t});\n\t\t\t}\n\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.store('global', {\n\t\t\t\t\ttoasts: [],\n\n\t\t\t\t\tshowErrorToast(message) {\n\t\t\t\t\t\tnewToast = {\n\t\t\t\t\t\t\ttoastText: message,\n\t\t\t\t\t\t\ttoastType: TOAST_TYPES.ERROR,\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.pushToast(newToast)\n\t\t\t\t\t},\n\n\t\t\t\t\tshowSuccessToast(message) {\n\t\t\t\t\t\tnewToast = {\n\t\t\t\t\t\t\ttoastText: message,\n\t\t\t\t\t\t\ttoastType: TOAST_TYPES.SUCCESS,\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.pushToast(newToast)\n\t\t\t\t\t},\n\n\t\t\t\t\tpushToast(toast) {\n\t\t\t\t\t\tthis.toasts.push(toast)\n\t\t\t\t\t\tsetTimeout(() => {\n\t\t\t\t\t\t\tthis.toasts.shift()\n\t\t\t\t\t\t}, ERROR_TOAST_TIMEOUT)\n\t\t\t\t\t},\n\t\t\t\t})\n\t\t\t})\n\n\t\t\tdocument.body.addEventListener('htmx:responseError', function(event) {\n\t\t\t\tif (!event.detail.xhr.status.toString().startsWith('2')) {\n\t\t\t\t\tAlpine.store('global').showErrorToast(event.detail.xhr.responseText)\n\t\t\t\t}\n\t\t\t});\n\n\t\t\tdocument.body.addEventListener('htmx:afterRequest', function(event) {\n\t\t\t\tif (event.detail.xhr.status === 200 && event.detail.requestConfig.verb !== 'get') {\n\t\t\t\t\tswitch (event.detail.requestConfig.verb) {\n\t\t\t\t\t\tcase 'post':\n\t\t\t\t\t\t\tAlpine.store('global').showSuccessToast('Successful Update')\n\t\t\t\t\t\t\tbreak\n\t\t\t\t\t\tcase 'delete':\n\t\t\t\t\t\t\tAlpine.store('global').showSuccessToast('Successful Delete')\n\t\t\t\t\t\t\tbreak\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t});\n\t\t</script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
