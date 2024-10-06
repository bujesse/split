// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Stats --><div hx-get=\"/api/expenses/stats\" hx-trigger=\"load, reloadExpenses from:body\"></div><div class=\"mt-6 flex w-full flex-col items-center justify-center\" x-data=\"{ offset: 0, isLastOffset: true }\"><!-- Expense Table --><div class=\"mt-6 w-full\" id=\"expenses-table\" hx-get=\"/api/expenses\" hx-trigger=\"load, reloadExpenses from:body\" @htmx:after-request=\"offset = 1; isLastOffset = JSON.parse([...document.querySelectorAll(&#39;#is-last-offset&#39;)].pop().innerText)\"></div><button class=\"btn btn-primary mt-6 w-full\" hx-trigger=\"click\" hx-target=\"#expenses-table\" hx-get=\"/api/expenses\" :hx-vals=\"JSON.stringify({ offset: offset })\" hx-swap=\"beforeend\" x-cloak @htmx:after-request=\"offset += 1; isLastOffset = JSON.parse([...document.querySelectorAll(&#39;#is-last-offset&#39;)].pop().innerText)\" :class=\"isLastOffset &amp;&amp; &#39;hidden&#39;\">Load More</button></div><!-- Add Expense FAB --><button class=\"btn btn-circle btn-success btn-lg fixed bottom-6 right-6 md:hidden\" onclick=\"baseModal.showModal()\" hx-get=\"/partials/expenses/new\" hx-trigger=\"click\" hx-target=\"#modal-container\"><svg class=\"h-6 w-6\" fill=\"#000000\" version=\"1.1\" id=\"Capa_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"800px\" height=\"800px\" viewBox=\"0 0 45.402 45.402\" xml:space=\"preserve\"><g><path d=\"M41.267,18.557H26.832V4.134C26.832,1.851,24.99,0,22.707,0c-2.283,0-4.124,1.851-4.124,4.135v14.432H4.141\n\t\tc-2.283,0-4.139,1.851-4.138,4.135c-0.001,1.141,0.46,2.187,1.207,2.934c0.748,0.749,1.78,1.222,2.92,1.222h14.453V41.27\n\t\tc0,1.142,0.453,2.176,1.201,2.922c0.748,0.748,1.777,1.211,2.919,1.211c2.282,0,4.129-1.851,4.129-4.133V26.857h14.435\n\t\tc2.283,0,4.134-1.867,4.133-4.15C45.399,20.425,43.548,18.557,41.267,18.557z\"></path></g></svg></button><script>\n\t\t\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
