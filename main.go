package main

import (
	"runtime"

	"github.com/willglynn/hatt/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.RootCmd.Execute()
}
