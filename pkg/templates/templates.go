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

	Struct = template.Must(template.New("Struct").Parse(
		"type {{.Name}} struct {\n" +
			"{{range .Fields}}{{.ToStructField}}\n{{end}}}",
	))

	Repo = template.Must(template.New("Repo").Parse(
		`{{.Struct.Generate}}`,
	))
)
