package app

import (
	"scrooge/pkg/param"
	"scrooge/pkg/templates"
)

func (mod *Mod) D() {
	parErr := param.New().SetName("err").SetTyp("error")
	pk := mod.pg.GetTable(mod.Name).GetPk()

	parID := param.New().SetName(mod.conv.PgToPascalCase(pk.Name)).
		SetTyp(mod.conv.PgTypeToGo(pk.Type)).SetIn(param.Path).SetSQLName(pk.Name).
		SetJSON(mod.conv.PgToCamelCase(pk.Name))

	repoMethod := mod.newRepoMethod(
		templates.BodyRepoDelete,
		"Delete",
		param.Params{parID},
		param.Params{parErr},
		nil,
	)

	mod.newControllerMethod(
		templates.BodyControllerDelete,
		"Delete",
		param.Params{param.New().SetName("ctx").SetPkg("fiber").
			SetPointer(true).SetTyp("Ctx")},
		param.Params{parErr},
		nil,
		repoMethod,
	)
}
