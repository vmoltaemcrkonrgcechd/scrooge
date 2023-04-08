package main

import (
	"log"
	"scrooge/pkg/app"
	"scrooge/pkg/converter"
	"scrooge/pkg/pgerudite"
)

func main() {
	pg, err := pgerudite.New("postgresql://postgres:4100@:4100/pocu?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	conv := converter.New()

	mod := app.NewMod("weapon", pg, conv)

	mod.C("name", "attack", "weight")
	mod.D()

	mod.Generate()
}
