package app

import (
	"encoding/json"
	"scrooge/pkg/converter"
	"scrooge/pkg/pgerudite"
)

type App struct {
	Mods map[string]*Mod
	pg   *pgerudite.PgErudite
	conv *converter.Converter
}

func New(pg *pgerudite.PgErudite, conv *converter.Converter) *App {
	return &App{
		Mods: make(map[string]*Mod),
		pg:   pg,
		conv: conv,
	}
}

func (app *App) Command(command CommandJSON) {
	if _, ok := app.Mods[command.Mod]; !ok {
		app.Mods[command.Mod] = NewMod(command.Mod, app.pg, app.conv)
	}

	switch command.Typ {

	case "c":
		var cmdC CommandC
		json.Unmarshal(command.Info, &cmdC)
		app.Mods[command.Mod].C(cmdC.Columns...)

	case "r":
		app.Mods[command.Mod].R()

	case "u":
		var cmdU CommandU
		json.Unmarshal(command.Info, &cmdU)
		app.Mods[command.Mod].U(cmdU.Columns...)

	case "d":
		app.Mods[command.Mod].D()

	}
}

func (app *App) Exec() {
	for _, mod := range app.Mods {
		mod.Generate()
	}
}
