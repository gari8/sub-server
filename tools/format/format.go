package format

import (
	"fmt"
	"strings"
)

const (
	PBlack = iota + 30
	PRed
	PGreen
	PYellow
	PBlue
	PMagenta
	PCyan
	PWhite
)

const logo = `
	░██████╗██╗░░░██╗██████╗░  ░██████╗███████╗██████╗░██╗░░░██╗███████╗██████╗░
	██╔════╝██║░░░██║██╔══██╗  ██╔════╝██╔════╝██╔══██╗██║░░░██║██╔════╝██╔══██╗
	╚█████╗░██║░░░██║██████╦╝  ╚█████╗░█████╗░░██████╔╝╚██╗░██╔╝█████╗░░██████╔╝
	░╚═══██╗██║░░░██║██╔══██╗  ░╚═══██╗██╔══╝░░██╔══██╗░╚████╔╝░██╔══╝░░██╔══██╗
	██████╔╝╚██████╔╝██████╦╝  ██████╔╝███████╗██║░░██║░░╚██╔╝░░███████╗██║░░██║
	╚═════╝░░╚═════╝░╚═════╝░  ╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝
`

const (
	FileContent = `[app]
server_name = "server name"
port = "8080"

[routing]
root_path = "."
origin_root = "/app/v1"

[[routing.origins]]
uri = "/start"
file_path = "index.json"
method = "GET"

[[routing.origins]]
uri = "/start"
file_path = "post.json"
method = "POST"`
	ResponseIndexFileContent = `[
  {
    "id": 1,
    "title": "Who I am"
  },
  {
    "id": 2,
    "title": "Who I am"
  },
  {
    "id": 3,
    "title": "Who I am"
  },
  {
    "id": 4,
    "title": "Who I am"
  }
]`
	ResponsePostFileContent = `{
  "id": 5,
  "title": "new one"
}`
)

func Print(number int, str string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", number, str)
}

func Logo() {
	Print(PCyan, logo)
}

func SlashAssign(text string) string {
	cut := strings.Trim(text, "/")
	return "/" + cut
}

func RemoveDotSlash(text string) string {
	return strings.Trim(text, "./")
}
