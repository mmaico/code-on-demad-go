package support

import (
	"io"
	"net/http"
	v8 "rogchap.com/v8go"
)

func MaskData(value string) string {
	baseUrl := "http://localhost:3000/json-mask.js"

	client := http.Client{}

	resp, _ := client.Get(baseUrl)
	defer resp.Body.Close()
	bodyStr, _ := io.ReadAll(resp.Body)

	iso := v8.NewIsolate()
	global := v8.NewObjectTemplate(iso)
	ctx := v8.NewContext(global, iso)

	ctx.RunScript(string(bodyStr), "json-our.js")

	json := "{\"value\": \"" + value + "\" }"
	setting := "[{\"type\": \"" + "MOBILE" + "\", \"path\": \"$.value\"}]"

	val, _ := ctx.RunScript("EntryPoint.default("+json+","+setting+")", "json-mask.js")

	return val.String()
}
