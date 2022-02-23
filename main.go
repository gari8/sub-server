package main

import (
	"flag"
	"fmt"
	"github.com/gari8/sub-server/execution"
	fileManager "github.com/gari8/sub-server/file-manager"
	"github.com/gari8/sub-server/server"
	"github.com/gari8/sub-server/tools/format"
	"os"
)

type RunMode string

type Config struct {
	Mode RunMode
}

const (
	initMode              RunMode = "init"
	serveMode             RunMode = "serve"
	configFileName                = "config.toml"
	responseIndexFileName         = "index.json"
	responsePostFileName          = "post.json"
	hint                          = `hint: Enter the following subcommand in order for you to use this command

  // creating server config file
  sub-server init

  // Starting the server
  sub-server serve
`
)

func main() {
	var config Config
	flag.Parse()
	switch flag.Arg(0) {
	case string(initMode):
		config.Mode = initMode
	case string(serveMode):
		config.Mode = serveMode
	}

	switch config.Mode {
	case initMode:
		if err := execution.New(fileManager.NewFileManager(configFileName),
			nil).RunInit([]byte(format.FileContent)); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %+v", err)
		}

		if err := execution.New(fileManager.NewFileManager(responseIndexFileName),
			nil).RunInit([]byte(format.ResponseIndexFileContent)); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %+v", err)
		}

		err := execution.New(fileManager.NewFileManager(responsePostFileName),
			nil).RunInit([]byte(format.ResponsePostFileContent))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %+v", err)
		}
	case serveMode:
		fm := fileManager.NewFileManager(configFileName)
		s := server.NewHttpServer()
		e := execution.New(fm, s)
		err := e.RunServe()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: %+v", err)
		}
	default:
		format.Print(format.PYellow, hint)
	}
}
