//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./mock.go
package execution

import (
	"github.com/gari8/sub-server/setting"
	"github.com/gari8/sub-server/tools/format"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

const (
	fileContent = `[app]
server_name = "server name"
port = "8080"

[routing]
root_path = "_template/app/v1/test"
origin_root = "app/v1/test"

[[routing.origins]]
id = 0
uri = "/start"
file_path = "/start/index.json"
method = "GET"

[[routing.origins]]
id = 1
uri = "/start/1"
file_path = "/start/1/index.json"
method = "GET"

[[routing.origins]]
id = 2
uri = "/start"
file_path = "/start/index.json"
method = "POST"`
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

// New Executionを作成
func New(fm FileManager, serv Server) Execution {
	return Execution{
		FileManager: fm,
		Server: serv,
	}
}

func (e Execution) RunInit () error {
	return e.FileManager.Create([]byte(fileContent))
}

func (e Execution)RunServe () error {
	// toml 読み込み
	bytes, err := e.FileManager.Read()
	hasToml := true
	if err != nil {
		format.Print(format.PCyan, "No config file")
		hasToml = false
	}

	// toml 書き出し
	stg := setting.NewSetting()
	if hasToml {
		err := stg.ReadToml(string(bytes))
		if err != nil {
			return err
		}
	}

	// json 読み込み
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
			if stg.RootPath+origin.FilePath == path {
				origin.Content = bytes
				newOrigins = append(newOrigins, origin)
			}
		}
		return nil
	}); err != nil {
		format.Print(format.PRed, err.Error())
	}
	stg.Origins = newOrigins

	// server 起動
	return e.Server.Serve(stg)
}