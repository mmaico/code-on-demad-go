package support

import (
	"os"
	v8 "rogchap.com/v8go"
)

func MaskData(value string) string {
	content, _ := os.ReadFile("/Users/peo_mjesus/mmaico/code-on-demand/src/main/resources/js-code/json-mask.js")

	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)
	ctx := v8.NewContext(global, iso)

	ctx.RunScript(string(content), "json-our.js")

	json := "{\"value\": \"" + value + "\" }"
	setting := "[{\"type\": \"" + "MOBILE" + "\", \"path\": \"$.value\"}]"

	val, _ := ctx.RunScript("EntryPoint.default("+json+","+setting+")", "/Users/peo_mjesus/mmaico/code-on-demand/src/main/resources/js-code/json-mask.js")

	return val.String()
}
