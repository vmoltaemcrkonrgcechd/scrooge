package app

import (
	"fmt"
	"scrooge/pkg/converter"
	"scrooge/pkg/pgerudite"
	"scrooge/pkg/repo"
	"scrooge/pkg/struct_builder"
)

type Mod struct {
	Name     string
	Entities []*struct_builder.Struct
	Repo     *repo.Repo
	pg       *pgerudite.PgErudite
	conv     *converter.Converter
}

func NewMod(name string) *Mod {
	return &Mod{
		Name: name,
	}
}

func (mod *Mod) C(fields ...string) *Mod {
	//создать структуру
	str := struct_builder.New()
	strName := fmt.Sprintf("Add%sDTO", mod.Name)
	str.SetName(strName)

	return mod
}
