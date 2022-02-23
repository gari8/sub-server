package setting

import (
	"github.com/BurntSushi/toml"
)

type Setting struct {
	TomlSetting
}

func NewSetting() Setting {
	return Setting{TomlSetting{
		App: App{
			ServerName: "server",
		},
		Routing: Routing{
			RootPath:   ".",
			OriginRoot: "/",
		},
	}}
}

type TomlSetting struct {
	App     `toml:"app"`
	Routing `toml:"routing"`
}

type App struct {
	ServerName string `toml:"server_name"`
	Port       string `toml:"port"`
}

type Routing struct {
	RootPath   string   `toml:"root_path"`
	OriginRoot string   `toml:"origin_root"`
	Origins    []Origin `toml:"origins"`
}

type Origin struct {
	URI      string `toml:"uri"`
	FilePath string `toml:"file_path"`
	Method   string `toml:"method"`
	Content  []byte
}

func (s *Setting) ReadToml(content string) error {
	if _, err := toml.Decode(content, &s); err != nil {
		return err
	}
	return nil
}
