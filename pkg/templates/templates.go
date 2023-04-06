package templates

import "text/template"

var (
	StructField = template.Must(template.New("StructField").Parse(
		"{{if not .Embedded}}{{.Name}} {{end}}{{if .Slice}}[]{{end}}{{if .Pointer}}*{{end}}" +
			"{{$length := len .Pkg}}{{if ne $length 0}}{{.Pkg}}.{{end}}" +
			"{{.Typ}} {{$length := len .JSON}}{{if ne $length 0}}`json:\"{{.JSON}}\"`{{end}}",
	))

	FuncParam = template.Must(template.New("FuncParam").Parse(
		"{{.LowerCaseName}} {{if .Slice}}[]{{end}}{{if .Pointer}}*{{end}}" +
			"{{$length := len .Pkg}}{{if ne $length 0}}{{.Pkg}}.{{end}}{{.Typ}}",
	))

	Recipient = template.Must(template.New("Recipient").Parse(
		"({{.LowerCaseName}} {{if .Pointer}}*{{end}}{{.Typ}})",
	))

	Params = template.Must(template.New("Params").Parse(
		"{{range $i,$v := .}}{{if ne $i 0}}, {{end}}{{.ToFuncParam}}{{end}}",
	))

	Struct = template.Must(template.New("Struct").Parse(
		"type {{.Name}} struct {\n" +
			"{{range .Fields}}{{.ToStructField}}\n{{end}}}",
	))

	Repo = template.Must(template.New("Repo").Parse(
		`{{.Struct.Generate}}`,
	))

	Func = template.Must(template.New("Func").Parse(
		"func {{if ne .Recipient nil}}{{.Recipient.ToRecipient}} {{end}}" +
			"{{.Name}}({{.Params.Generate}}) " +
			"{{$length := len .Returns}}{{if ge $length 2}}({{end}}" +
			"{{.Params.Generate}}{{if ge $length 2}}){{end}} {\n" +
			"{{.Body.Generate}}\n}",
	))
)
