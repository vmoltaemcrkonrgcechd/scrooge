package app

import "encoding/json"

type Cfg struct {
	URL      string        `json:"url"`
	Commands []CommandJSON `json:"commands"`
}

type CommandJSON struct {
	Typ  string `json:"typ"`
	Mod  string `json:"mod"`
	Info json.RawMessage
}

type CommandC struct {
	Columns []string `json:"columns"`
}

type CommandU struct {
	Columns []string `json:"columns"`
}
