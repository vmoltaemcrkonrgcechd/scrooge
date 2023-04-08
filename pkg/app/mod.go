package app

import (
	"fmt"
	"scrooge/pkg/controller"
	"scrooge/pkg/converter"
	"scrooge/pkg/func_builder"
	"scrooge/pkg/param"
	"scrooge/pkg/pgerudite"
	"scrooge/pkg/repo"
	"scrooge/pkg/struct_builder"
	"scrooge/pkg/templates"
	"text/template"
)

type Mod struct {
	Name       string
	Entities   []*struct_builder.Struct
	Repo       *repo.Repo
	Controller *controller.Controller
	pg         *pgerudite.PgErudite
	conv       *converter.Converter
}

func NewMod(name string,
	pg *pgerudite.PgErudite,
	conv *converter.Converter) *Mod {

	mod := &Mod{
		Name: name,
		pg:   pg,
		conv: conv,
	}

	mod.newRepo()
	mod.newController()

	return mod
}

func (mod *Mod) C(fields ...string) *Mod {
	//создать структуру
	strName := fmt.Sprintf("Add%sDTO", mod.conv.PgToPascalCase(mod.Name))
	str := mod.newStruct(strName, fields)

	//создать метод репозитория
	pk := mod.pg.GetTable(mod.Name).GetPk()

	parID := param.New().SetName(mod.conv.PgToCamelCase(pk.Name)).
		SetTyp(mod.conv.PgTypeToGo(pk.Type)).SetIn(param.Path).SetSQLName(pk.Name)

	parErr := param.New().SetName("err").SetTyp("error")
	parStr := param.New().SetName(strName).SetTyp(strName).SetPkg("entities").
		SetIn(param.Body)

	mod.newRepoMethod(
		templates.BodyRepoAdd,
		fmt.Sprintf("Add"),
		param.Params{parStr},
		param.Params{parID, parErr},
		str,
	)

	//создать метод контроллера

	return mod
}

func (mod *Mod) newRepo() {
	repoPar := param.New().SetEmbedded(true).SetPointer(true).
		SetName("PG").SetTyp("PG").SetPkg("pg")

	repoName := fmt.Sprintf("%sRepo", mod.conv.PgToPascalCase(mod.Name))

	repoStr := struct_builder.New().SetName(repoName).AddField(repoPar)

	repoConstructor := func_builder.New().SetName(fmt.Sprintf("New%s", repoName)).
		AddParam(repoPar).AddReturn(param.New().SetName(repoName).SetTyp(repoName)).
		SetBody(func_builder.NewBody().SetTpl(templates.Constructor).
			SetParams(param.Params{repoPar}).
			SetReturns(param.Params{param.New().SetName(repoName).SetTyp(repoName)}))

	mod.Repo = repo.New().SetStruct(repoStr).SetConstructor(repoConstructor)
}

func (mod *Mod) newController() {
	ctrParam := param.New().SetName(mod.Repo.Struct.Name).
		SetTyp(mod.Repo.Struct.Name).SetPkg("repo")
	ctrName := fmt.Sprintf("%sController", mod.conv.PgToPascalCase(mod.Name))
	str := struct_builder.New().SetName(ctrName).
		AddField(ctrParam)

	ctr := controller.New().SetStruct(str).
		SetConstructor(
			func_builder.New().SetName(fmt.Sprintf("New%s", ctrName)).
				AddParam(ctrParam).AddReturn(param.New().SetName(ctrName).
				SetTyp(ctrName)).SetBody(
				func_builder.NewBody().SetTpl(templates.Constructor).
					SetParams(param.Params{ctrParam}).
					SetReturns(param.Params{param.New().SetName(ctrName).SetTyp(ctrName)}),
			),
		)

	mod.Controller = ctr
}

func (mod *Mod) newRepoMethod(
	tpl *template.Template,
	fnName string,
	params, returns param.Params,
	str *struct_builder.Struct,
) *func_builder.Func {
	recipient := param.New().SetName(mod.Repo.Struct.Name).SetTyp(mod.Repo.Struct.Name)

	body := func_builder.NewBody().SetTpl(tpl).SetReturns(returns).SetParams(params).
		SetTable(mod.Name).SetStruct(str).SetRecipient(recipient)

	fn := func_builder.New().SetName(fnName).SetRecipient(recipient).SetBody(body)

	for _, p := range params {
		fn.AddParam(p)
	}

	for _, r := range returns {
		fn.AddReturn(r)
	}

	mod.Repo.Methods = append(mod.Repo.Methods, fn)

	return fn
}

func (mod *Mod) newStruct(strName string, fields []string) *struct_builder.Struct {
	str := struct_builder.New().SetName(strName)

	for _, f := range fields {
		col := mod.pg.GetColumn(mod.Name, f)

		str.AddField(
			param.New().
				SetName(mod.conv.PgToPascalCase(col.Name)).
				SetTyp(mod.conv.PgTypeToGo(col.Type)).
				SetPointer(col.IsNullable).
				SetJSON(mod.conv.PgToCamelCase(col.Name)).
				SetSQLName(col.Name),
		)
	}

	mod.Entities = append(mod.Entities, str)

	return str
}

func (mod *Mod) Generate() {
	fmt.Println("package entities")

	for _, e := range mod.Entities {
		fmt.Println(e.Generate())
	}

	fmt.Println("package repo")
	fmt.Println(mod.Repo.Generate())

	fmt.Println("package controller")
	fmt.Println(mod.Controller.Generate())
}
