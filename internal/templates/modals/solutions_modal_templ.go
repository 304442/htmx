// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package modals

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SolutionsModal() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"modal hidden fixed inset-0 bg-gray-900 bg-opacity-50 flex items-center justify-center \" id=\"modal-2\"><div class=\"modal-content bg-white rounded-lg max-w-4xl w-full md:min-w-[720px] max-h-[80vh] flex flex-col\"><div class=\"flex justify-between items-center bg-gray-100 px-4 py-2 rounded-t-lg border-b border-gray-500 bg-slate-300\"><h2 class=\"text-xl font-bold\">Solutions</h2><span class=\"material-symbols-outlined close-modal cursor-pointer text-gray-500 hover:text-gray-800\">close</span></div><div class=\"p-4 overflow-y-auto flex-grow\"><div hx-get=\"/marketplace\" hx-trigger=\"load\" hx-target=\"#services\"><div id=\"services\"></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
