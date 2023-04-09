package app

import (
	"fmt"
	"scrooge/pkg/param"
	"scrooge/pkg/templates"
)

func (mod *Mod) R() {
	//создать структуру
	strName := fmt.Sprintf("%s", mod.conv.PgToPascalCase(mod.Name))
	str := mod.newStruct(strName, mod.pg.GetTable(mod.Name).Names())

	//создать метод репозитория
	parErr := param.New().SetName("err").SetTyp("error")

	rm := mod.newRepoMethod(
		templates.BodyRepoAll,
		"All",
		param.Params{},
		param.Params{
			param.New().SetSlice(true).
				SetName(str.Name).
				SetTyp(str.Name).
				SetPkg("entities").
				SetIn(param.Body),
			parErr},
		str,
	)

	//создать метод контроллера
	mod.newControllerMethod(
		templates.BodyControllerAll,
		"All",
		param.Params{param.New().SetName("ctx").SetPkg("fiber").
			SetPointer(true).SetTyp("Ctx")},
		param.Params{parErr},
		str,
		rm,
		"Get",
	)
}
