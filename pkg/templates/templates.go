package templates

import "text/template"

var (
	StructField = template.Must(template.New("StructField").Parse(
		"{{if not .Embedded}}{{.Name}} {{end}}{{if .Slice}}[]{{end}}{{if .Pointer}}*{{end}}" +
			"{{$length := len .Pkg}}{{if ne $length 0}}{{.Pkg}}.{{end}}" +
			"{{.Typ}} {{$length := len .JSON}}{{if ne $length 0}}`json:\"{{.JSON}}\"`{{end}}",
	))

	StructFieldWithParams = template.Must(template.New("StructFieldWithParams").Parse(
		"{{if not .Embedded}}{{.Name}} {{end}}{{if .Slice}}[]{{end}}{{if .Pointer}}*{{end}}" +
			"{{$length := len .Pkg}}{{if ne $length 0}}{{.Pkg}}.{{end}}" +
			"{{.Typ}} {{$length := len .JSON}}{{if ne $length 0}}`query:\"{{.JSON}}\"`{{end}}",
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

	Method = template.Must(template.New("Func").Parse(
		"// @tags {{.ModName}}\n" +
			"{{range .RepoMethod.Params}}" +
			"// @param {{.LowerCaseName}} " +
			"{{if eq .In 0}}body{{else}}query{{end}} " +
			"{{.GetTyp}} true \"{{.LowerCaseName}}\"" +
			"\n" +
			"{{end}}" +
			"// @router /{{.ModName}} [{{.LowerTyp}}]\n" +
			"func {{if ne .Recipient nil}}{{.Recipient.ToRecipient}} {{end}}" +
			"{{.Name}}({{.Params.Generate}}) " +
			"({{.Returns.Generate}}) {\n" +
			"{{.Body.Generate}}\n}",
	))

	BodyRepoAdd = template.Must(template.New("BodyRepoAdd").Parse(
		"if err = {{.Recipient.LowerCaseName}}.Sq.Insert(\"{{.Table}}\").\n" +
			"Columns({{.Struct.Fields.ParamSQLNames}}).\n" +
			"Values(" +
			"{{$str := .Struct}}" +
			"{{range $i,$v := .Struct.Fields}}{{if ne $i 0}}, {{end}}{{$str.LowerCaseName}}.{{.Name}}{{end}}" +
			").\n" +
			"Suffix(\"RETURNING {{.Returns.Path.SQLName}}\").\n" +
			"QueryRow().Scan(&{{.Returns.Path.LowerCaseName}}); err != nil {\n" +
			"return {{.Returns.ParamNames}}\n}\n" +
			"return {{.Returns.ParamNames}}",
	))

	BodyRepoAll = template.Must(template.New("BodyRepoAll").Parse(
		"rows, err := {{.Recipient.LowerCaseName}}." +
			"Sq.Select({{.Struct.Fields.ParamSQLNames}})." +
			"From(\"{{.Table}}\").Query()\n" +
			"if err != nil {\n" +
			"return {{.Returns.ParamNames}}\n}\n" +
			"var el entities.{{.Struct.Name}}\n" +
			"for rows.Next() {\n" +
			"if err = rows.Scan(" +
			"{{range $i,$v := .Struct.Fields}}{{if ne $i 0}}, {{end}}&el.{{.Name}}{{end}}" +
			"); err != nil {\n" +
			"return {{.Returns.ParamNames}}\n}\n" +
			"{{.Returns.Body.LowerCaseName}} = append({{.Returns.Body.LowerCaseName}}, el)\n}\n" +
			"return {{.Returns.ParamNames}}",
	))

	BodyRepoEdit = template.Must(template.New("BodyRepoEdit").Parse(

		"if err = {{.Recipient.LowerCaseName}}.Sq.Update(\"{{.Table}}\").\n" +
			"{{$str := .Struct}}{{range .Struct.Fields}}Set(\"{{.SQLName}}\", {{$str.LowerCaseName}}." +
			"{{.Name}}).\n{{end}}" +
			"Where(\"{{.Params.Path.SQLName}} = ?\", {{.Params.Path.LowerCaseName}}).\n" +
			"Suffix(\"RETURNING {{.Returns.Path.SQLName}}\").QueryRow().\n" +
			"Scan(&{{.Returns.Path.LowerCaseName}}); err != nil {\n" +
			"return {{.Returns.ParamNames}}\n}\n" +
			"return {{.Returns.ParamNames}}",
	))

	BodyRepoDelete = template.Must(template.New("BodyRepoDelete").Parse(
		"if _, err = {{.Recipient.LowerCaseName}}.Sq.Delete(\"{{.Table}}\").\n" +
			"Where(\"{{.Params.Path.SQLName}} = ?\", {{.Params.Path.LowerCaseName}}).Exec(); err != nil {\n" +
			"return {{.Returns.ParamNames}}\n}\n" +
			"return {{.Returns.ParamNames}}",
	))

	BodyControllerAdd = template.Must(template.New("BodyControllerAdd").Parse(
		"var {{.RepoMethod.Params.Body.ToFuncParam}}\n" +
			"if err = ctx.BodyParser(&{{.RepoMethod.Params.Body.LowerCaseName}}); " +
			"err != nil {\n" +
			"return err\n}\n" +
			"var {{.RepoMethod.Returns.Path.ToFuncParam}}\n" +
			"if {{.RepoMethod.Returns.ParamNames}} = " +
			"{{.Recipient.LowerCaseName}}.{{.RepoMethod.Recipient.Name}}" +
			".{{.RepoMethod.Name}}({{.RepoMethod.Params.ParamNames}}); " +
			"err != nil {\n" +
			"return err\n}\n" +
			"return ctx.Status(201).JSON(struct{ {{.RepoMethod.Returns.Path.ToStructField}} }" +
			"{ {{.RepoMethod.Returns.Path.LowerCaseName}} })",
	))

	BodyControllerEdit = template.Must(template.New("BodyControllerEdit").Parse(
		"var params struct { {{.RepoMethod.Params.Path.ToStructFieldWithParams}} }\n" +
			"ctx.QueryParser(&params)\n" +
			"var {{.RepoMethod.Params.Body.ToFuncParam}}\n" +
			"if err = ctx.BodyParser(&{{.RepoMethod.Params.Body.LowerCaseName}}); " +
			"err != nil {\n" +
			"return err\n}\n" +
			"var {{.RepoMethod.Returns.Path.ToFuncParam}}\n" +
			"if {{.RepoMethod.Returns.ParamNames}} = " +
			"{{.Recipient.LowerCaseName}}.{{.RepoMethod.Recipient.Name}}" +
			".{{.RepoMethod.Name}}({{.RepoMethod.Params.Body.LowerCaseName}}, params.{{.RepoMethod.Params.Path.Name}}); " +
			"err != nil {\n" +
			"return err\n}\n" +
			"return ctx.JSON(struct{ {{.RepoMethod.Returns.Path.ToStructField}} }" +
			"{ {{.RepoMethod.Returns.Path.LowerCaseName}} })",
	))

	BodyControllerDelete = template.Must(template.New("BodyControllerDelete").Parse(
		"var params struct { {{.RepoMethod.Params.Path.ToStructFieldWithParams}} }\n" +
			"ctx.QueryParser(&params)\n" +
			"if err = {{.Recipient.LowerCaseName}}." +
			"{{.RepoMethod.Recipient.Name}}." +
			"{{.RepoMethod.Name}}(params.{{.RepoMethod.Params.Path.Name}}); err != nil {\n" +
			"return err\n}\n" +
			"return nil",
	))

	BodyControllerAll = template.Must(template.New("BodyControllerAll").Parse(
		"var {{.RepoMethod.Returns.Body.ToFuncParam}}\n" +
			"if {{.RepoMethod.Returns.Body.LowerCaseName}}, err = " +
			"{{.Recipient.LowerCaseName}}.{{.RepoMethod.Recipient.Name}}" +
			".{{.RepoMethod.Name}}(); err != nil {\n" +
			"return err\n}\n" +
			"return ctx.JSON({{.RepoMethod.Returns.Body.LowerCaseName}})",
	))

	Constructor = template.Must(template.New("Constructor").Parse(
		"{{.Returns.ParamNames}} = " +
			"{{range .Returns}}{{.Typ}}{{end}}{ {{.Params.ParamNames}} }\n" +
			"return {{.Returns.ParamNames}}",
	))

	Main = template.Must(template.New("Main").Parse(
		"package main\n" +
			"func main() {\n" +
			"app := fiber.New()\n" +
			"app.Use(cors.New(cors.Config{\n" +
			"AllowOrigins:     \"http://localhost:8080\",\n" +
			"AllowMethods:     \"GET,POST,HEAD,DELETE,PATCH\",\n" +
			"AllowCredentials: true,\n" +
			"}))\n" +
			"pq, err := pg.New(\"{{.PgURL}}\")\n" +
			"if err != nil {\n" +
			"panic(err)\n}\n" +
			"app.Get(\"/swagger-ui/*\", swagger.New(swagger.ConfigDefault))\n" +
			"{{range $k,$v := .Mods}}" +
			"{{$v.Repo.Struct.Name}} := repo.{{$v.Repo.Constructor.Name}}(pq)\n" +
			"{{$v.Controller.Struct.Name}} := controller.{{$v.Controller.Constructor.Name}}" +
			"({{$v.Repo.Struct.Name}})\n" +
			"{{range $v.Controller.Methods}}app.{{.Typ}}" +
			"(\"/{{.ModName}}\", {{$v.Controller.Struct.Name}}.{{.Name}})\n{{end}}" +
			"{{end}}" +
			"log.Fatal(app.Listen(\":80\"))\n}\n",
	))

	PG = template.Must(template.New("PG").Parse(
		"package pg\n" +
			"import (\n" +
			"\"database/sql\"\n" +
			"sq \"github.com/Masterminds/squirrel\"\n" +
			"_ \"github.com/lib/pq\"\n)\n" +
			"type PG struct {\n" +
			"Sq sq.StatementBuilderType\n" +
			"DB *sql.DB\n}\n" +
			"func New(url string) (*PG, error) {\n" +
			"db, err := sql.Open(\"postgres\", url)\n" +
			"if err != nil {\n" +
			"return nil, err\n}\n" +
			"return &PG{\n" +
			"Sq: sq.StatementBuilder.\n" +
			"PlaceholderFormat(sq.Dollar).\n" +
			"RunWith(db),\n" +
			"DB: db,\n" +
			"}, nil\n}\n",
	))
)
