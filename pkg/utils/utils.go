package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func MustExecTemplate(tpl *template.Template, data any) string {
	buf := new(bytes.Buffer)

	if err := tpl.Execute(buf, data); err != nil {
		fmt.Println(err)

		return ""
	}

	return buf.String()
}
