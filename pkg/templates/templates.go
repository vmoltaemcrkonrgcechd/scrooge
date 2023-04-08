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

	ParamNames = template.Must(template.New("ParamNames").Parse(
		"{{range $i,$v := .}}{{if ne $i 0}}, {{end}}{{.LowerCaseName}}{{end}}",
	))

	ParamSQLNames = template.Must(template.New("ParamSQLNames").Parse(
		"{{range $i,$v := .}}{{if ne $i 0}}, {{end}}\"{{.SQLName}}\"{{end}}",
	))

	Struct = template.Must(template.New("Struct").Parse(
		"type {{.Name}} struct {\n" +
			"{{range .Fields}}{{.ToStructField}}\n{{end}}}",
	))

	Repo = template.Must(template.New("Repo").Parse(
		"{{.Struct.Generate}}\n" +
			"{{.Constructor.Generate}}\n" +
			"{{range .Methods}}{{.Generate}}\n{{end}}",
	))

	Func = template.Must(template.New("Func").Parse(
		"func {{if ne .Recipient nil}}{{.Recipient.ToRecipient}} {{end}}" +
			"{{.Name}}({{.Params.Generate}}) " +
			"({{.Returns.Generate}}) {\n" +
			"{{.Body.Generate}}\n}",
	))

	BodyRepoAdd = template.Must(template.New("BodyRepoAdd").Parse(
		"if _, err = {{.Recipient.LowerCaseName}}.Sq.Insert(\"{{.Table}}\").\n" +
			"Columns({{.Struct.Fields.ParamSQLNames}}).\n" +
			"Values(" +
			"{{$str := .Struct}}" +
			"{{range $i,$v := .Struct.Fields}}{{if ne $i 0}}, {{end}}{{$str.LowerCaseName}}.{{.Name}}{{end}}" +
			").\n" +
			"Suffix(\"RETURNING {{.Returns.Path.SQLName}}\")\n" +
			"QueryRow().Scan(&{{.Returns.Path.LowerCaseName}}); err != nil {\n" +
			"return {{.Returns.ParamNames}}\n}\n" +
			"return {{.Returns.ParamNames}}",
	))

	Constructor = template.Must(template.New("Constructor").Parse(
		"{{.Returns.ParamNames}} = " +
			"{{range .Returns}}{{.Typ}}{{end}}{ {{.Params.ParamNames}} }\n" +
			"return {{.Returns.ParamNames}}",
	))
)
