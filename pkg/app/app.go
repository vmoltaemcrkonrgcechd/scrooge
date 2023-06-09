package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"scrooge/pkg/converter"
	"scrooge/pkg/pgerudite"
	"scrooge/pkg/templates"
	"scrooge/pkg/utils"
	"strconv"
	"time"
)

type App struct {
	Mods    map[string]*Mod
	pg      *pgerudite.PgErudite
	conv    *converter.Converter
	t       string
	ModName string
	PgURL   string
}

func New(pg *pgerudite.PgErudite, conv *converter.Converter, target, pgURL string) *App {
	return &App{
		Mods:  make(map[string]*Mod),
		pg:    pg,
		conv:  conv,
		t:     target,
		PgURL: pgURL,
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
	app.createProjectStructure()

	var path = fmt.Sprintf("%s\\%s", app.t, app.ModName)

	for _, mod := range app.Mods {
		mod.GenerateEntities()

		app.mustWriteFile(path+"\\"+repoName+"\\"+mod.Name+".go", mod.GenerateRepo())
		app.mustWriteFile(path+"\\"+controllerName+"\\"+mod.Name+".go", mod.GenerateController())
		app.mustWriteFile(path+"\\"+entitiesName+"\\"+mod.Name+".go", mod.GenerateEntities())
	}

	app.mustWriteFile(path+"\\"+"main.go", app.getMain())
	app.mustWriteFile(path+"\\"+"pg\\"+"pg.go", app.getPG())

	app.mustExecCmd(path, "go", "mod", "init", app.ModName)
	app.mustExecCmd(path, "go", "get", "github.com/gofiber/swagger", app.ModName)
	app.mustExecCmd(path, "go", "get", "github.com/gofiber/fiber/v2", app.ModName)
	app.mustExecCmd(path, "go", "get", "github.com/Masterminds/squirrel", app.ModName)
	app.mustExecCmd(path, "go", "get", "github.com/lib/pq", app.ModName)
	app.mustExecCmd(path, "go", "get", "golang.org/x/tools/cmd/goimports", app.ModName)
	app.mustExecCmd(path, "swag", "init", ".", app.ModName)
	app.mustExecCmd(path, "goimports", "-w", ".", app.ModName)

	app.raiseVersion(path)
}

const (
	controllerName = "controller"
	repoName       = "repo"
	entitiesName   = "entities"
	pgName         = "pg"
)

func (app *App) createProjectStructure() {
	app.ModName = strconv.Itoa(int(time.Now().UnixNano()))

	app.mustMkdir("")
	app.mustMkdir(controllerName)
	app.mustMkdir(repoName)
	app.mustMkdir(entitiesName)
	app.mustMkdir(pgName)
}

func (app *App) mustMkdir(name string) {
	path := fmt.Sprintf("%s\\%s", app.t, app.ModName)

	if name != "" {
		path += "\\" + name
	}

	if err := os.Mkdir(path, 0666); err != nil {
		panic(err)
	}
}

func (app *App) mustWriteFile(name, data string) {
	if err := os.WriteFile(name, []byte(data), 0666); err != nil {
		panic(err)
	}
}

func (app *App) getMain() string {
	return utils.MustExecTemplate(templates.Main, app)
}

func (app *App) getPG() string {
	return utils.MustExecTemplate(templates.PG, app)
}

func (app *App) mustExecCmd(cd, name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = cd
	if err := cmd.Run(); err != nil {
		return
	}
}

func (app *App) mustReadFile(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return data
}

func (app *App) raiseVersion(path string) {
	data := app.mustReadFile(path + "\\" + "main.go")

	newData := bytes.ReplaceAll(data, []byte("\"github.com/gofiber/fiber\""), []byte("\"github.com/gofiber/fiber/v2\""))

	newData = bytes.ReplaceAll(newData, []byte("import ("),
		[]byte(fmt.Sprintf("import (\n\t_ \"%s/docs\"\n", app.ModName)))

	app.mustWriteFile(path+"\\"+"main.go", string(newData))

	files, err := os.ReadDir(path + "\\controller")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		data := app.mustReadFile(path + "\\controller\\" + f.Name())
		newData := bytes.ReplaceAll(data, []byte("\"github.com/gofiber/fiber\""), []byte("\"github.com/gofiber/fiber/v2\""))
		app.mustWriteFile(path+"\\controller\\"+f.Name(), string(newData))
	}
}
