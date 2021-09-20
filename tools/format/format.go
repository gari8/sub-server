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

func Print(number int, str string) {
	fmt.Printf("\x1b[%dm%s\x1b[0m\n", number, str)
}

func Logo() {
	Print(PCyan, logo)
}

func SlashAssign (text string) string {
	cut := strings.Trim(text, "/")
	return "/"+cut
}
