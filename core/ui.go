package core

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"
)

// Enable ANSI colors on Windows
func init() {
	if runtime.GOOS == "windows" {
		h := windows.Handle(windows.Stdout)
		var mode uint32
		if windows.GetConsoleMode(h, &mode) == nil {
			_ = windows.SetConsoleMode(h, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
		}
	}
}

// ANSI helpers
func color(code string) string {
	return code
}

func reset() string {
	return "\033[0m"
}

func Critical(s string) string {
	return color("\033[1;31m") + s + reset()
}

func Info(s string) string {
	return color("\033[1;36m") + s + reset()
}

func Dim(s string) string {
	return color("\033[2m") + s + reset()
}

// Symbols
const (
	SymCritical = "✖"
	SymTrust    = "➜"
	SymPath     = "📁"
	SymService  = "🛠"
	SymTask     = "⏱"
)

// Banner
func PrintBanner() {
	fmt.Println(Info(`
████████╗██████╗ ██╗   ██╗███████╗████████╗██████╗ ██████╗ ███████╗██╗  ██╗
╚══██╔══╝██╔══██╗██║   ██║██╔════╝╚══██╔══╝██╔══██╗██╔══██╗██╔════╝╚██╗██╔╝
   ██║   ██████╔╝██║   ██║███████╗   ██║   ██████╔╝██████╔╝█████╗   ╚███╔╝
   ██║   ██╔══██╗██║   ██║╚════██║   ██║   ██╔══██╗██╔══██╗██╔══╝   ██╔██╗
   ██║   ██║  ██║╚██████╔╝███████║   ██║   ██║  ██║██║  ██║███████╗██╔╝ ██╗
   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝
`))
	fmt.Println(Dim("SYSTEM Trust Analysis Engine — v1.0\n"))
}
