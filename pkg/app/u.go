package app

import (
	"fmt"
	"scrooge/pkg/param"
	"scrooge/pkg/templates"
)

func (mod *Mod) U(fields ...string) {
	//создать структуру
	strName := fmt.Sprintf("Edit%sDTO", mod.conv.PgToPascalCase(mod.Name))
	str := mod.newStruct(strName, fields)

	//создать метод репозитория
	pk := mod.pg.GetTable(mod.Name).GetPk()

	parID := param.New().SetName(mod.conv.PgToPascalCase(pk.Name)).
		SetTyp(mod.conv.PgTypeToGo(pk.Type)).SetIn(param.Path).SetSQLName(pk.Name).
		SetJSON(mod.conv.PgToCamelCase(pk.Name))

	parErr := param.New().SetName("err").SetTyp("error")
	parStr := param.New().SetName(strName).SetTyp(strName).SetPkg("entities").
		SetIn(param.Body)

	rm := mod.newRepoMethod(
		templates.BodyRepoEdit,
		"Edit",
		param.Params{parStr, parID},
		param.Params{parID, parErr},
		str,
	)

	//создать метод контроллера
	mod.newControllerMethod(
		templates.BodyControllerEdit,
		"Edit",
		param.Params{param.New().SetName("ctx").SetPkg("fiber").
			SetPointer(true).SetTyp("Ctx")},
		param.Params{parErr},
		str,
		rm,
	)
}
