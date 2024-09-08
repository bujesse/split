// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

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
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"typeaheadInit($el, $data)\" class=\"relative\" items=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(templ.JSONString(items))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/typeahead.templ`, Line: 7, Col: 33}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" model=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(model)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/typeahead.templ`, Line: 8, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" @keydown.down.prevent=\"focusNextItem()\" @keydown.up.prevent=\"focusPrevItem()\" @keydown.enter.prevent=\"selectItem(activeItem)\" @keydown.tab.stop=\"open = false; activeItem = null\"><label class=\"input input-bordered flex items-center gap-2\"><input type=\"text\" class=\"grow\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(placeholder)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/typeahead.templ`, Line: 18, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-model=\"search\" @focus=\"open = true\" @keydown=\"open = true\" @keydown.escape=\"open = false; activeItem = null\" @click.outside=\"open = false; activeItem = null\" @keydown.escape=\"open = false; activeItem = null\"> <input type=\"hidden\" x-model=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(model)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/typeahead.templ`, Line: 26, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/components/typeahead.templ`, Line: 26, Col: 53}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <span x-show=\"search !== &#39;&#39;\" @click=\"search = &#39;&#39;; open = false; activeItem = null\" class=\"badge badge-ghost cursor-pointer\">✕</span></label><div x-show=\"open\" class=\"absolute w-full\"><ul tabindex=\"0\" class=\"menu bg-base-200 rounded-box w-full block max-h-44 overflow-y-scroll\"><template x-for=\"item in filteredItems\" :key=\"item.ID\"><li><a @click=\"selectItem(item)\" x-text=\"getDisplayName(item)\" :class=\"activeItem === item &amp;&amp; &#39;active&#39;\"></a></li></template></ul></div></div><script>\n\t\tfunction typeaheadInit($el, $data) {\n\t\t\tconst items = JSON.parse($el.getAttribute('items'))\n\t\t\tconst model = $el.getAttribute('model')\n\t\t\tconst defaultDisplay = $el.getAttribute('defaultDisplay')\n\t\t\treturn {\n\t\t\t\tsearch: $data.defaultSearchDisplay || '',\n\t\t\t\topen: false,\n\t\t\t\titems: items,\n\t\t\t\tactiveItem: null,\n\t\t\t\tgetDisplayName(item) {\n\t\t\t\t\treturn item.Type + ' > ' + item.Name\n\t\t\t\t},\n\t\t\t\tfilteredItems() {\n\t\t\t\t\treturn this.items.filter(item => {\n\t\t\t\t\t\treturn this.getDisplayName(item).toLowerCase().includes(this.search.toLowerCase())\n\t\t\t\t\t})\n\t\t\t\t},\n\t\t\t\tfocusItem(item) {\n\t\t\t\t\tthis.activeItem = item\n\t\t\t\t},\n\t\t\t\tfocusNextItem() {\n\t\t\t\t\tif (!this.open) this.open = true\n\t\t\t\t\tconst items = this.filteredItems()\n\t\t\t\t\tconst index = items.indexOf(this.activeItem)\n\t\t\t\t\tif (index === -1) {\n\t\t\t\t\t\tthis.activeItem = items[0]\n\t\t\t\t\t} else if (index < items.length - 1) {\n\t\t\t\t\t\tthis.activeItem = items[index + 1]\n\t\t\t\t\t}\n\t\t\t\t},\n\t\t\t\tfocusPrevItem(item) {\n\t\t\t\t\tif (!this.open) this.open = true\n\t\t\t\t\tconst items = this.filteredItems()\n\t\t\t\t\tconst index = items.indexOf(this.activeItem)\n\t\t\t\t\tif (index === -1) {\n\t\t\t\t\t\tthis.activeItem = items[items.length - 1]\n\t\t\t\t\t} else if (index > 0) {\n\t\t\t\t\t\tthis.activeItem = items[index - 1]\n\t\t\t\t\t}\n\t\t\t\t},\n\t\t\t\tselectItem(item) {\n\t\t\t\t\tif (!item) return\n\t\t\t\t\tconst value = item.ID\n\t\t\t\t\tthis[model] = value\n\t\t\t\t\tthis.search = this.getDisplayName(item)\n\t\t\t\t\tthis.open = false\n\t\t\t\t\tthis.activeItem = null\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
