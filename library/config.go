package library

type DefConfigApp struct {
	Debug  bool           `toml:"debug"`
	Params map[string]any `toml:"params"`
}

var ConfigApp DefConfigApp
