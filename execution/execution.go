//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./mock.go
package execution

import (
	"github.com/gari8/sub-server/setting"
	"github.com/gari8/sub-server/tools/format"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

type FileManager interface {
	Create(content []byte) error
	Read() ([]byte, error)
}

type Server interface {
	Serve(setting setting.Setting) error
}

type Execution struct {
	FileManager
	Server
}

// New Creating Execution struct
func New(fm FileManager, serv Server) Execution {
	return Execution{
		FileManager: fm,
		Server:      serv,
	}
}

// RunInit Creating config file
func (e Execution) RunInit(content []byte) error {
	return e.FileManager.Create(content)
}

// RunServe Getting started serve
func (e Execution) RunServe() error {
	// reading toml
	bytes, err := e.FileManager.Read()
	if err != nil {
		return err
	}

	// writing toml
	stg := setting.NewSetting()
	if err := stg.ReadToml(string(bytes)); err != nil {
		return err
	}

	// reading json
	var newOrigins []setting.Origin
	if err := filepath.Walk(stg.TomlSetting.RootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".json" {
			return nil
		}
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		for _, origin := range stg.Origins {
			if format.RemoveDotSlash(stg.RootPath+format.SlashAssign(origin.FilePath)) == path {
				origin.Content = bytes
				newOrigins = append(newOrigins, origin)
			}
		}
		return nil
	}); err != nil {
		format.Print(format.PRed, err.Error())
	}
	stg.Origins = newOrigins

	// server
	return e.Server.Serve(stg)
}
