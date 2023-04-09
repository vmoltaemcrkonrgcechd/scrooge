package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"scrooge/pkg/app"
	"scrooge/pkg/converter"
	"scrooge/pkg/pgerudite"
)

func main() {
	t := flag.String("t", "", "")
	flag.Parse()

	if t == nil {
		log.Fatal("t")
	}

	data, err := os.ReadFile(*t)
	if err != nil {
		log.Fatal(err)
	}

	var cfg app.Cfg
	json.Unmarshal(data, &cfg)

	pg, err := pgerudite.New(cfg.URL)
	if err != nil {
		log.Fatal(err)
	}

	conv := converter.New()

	a := app.New(pg, conv)

	for _, cmd := range cfg.Commands {
		a.Command(cmd)
	}

	a.Exec()
}
