package main

import "github.com/willglynn/hatt/commands"

// goxc sets these variables via build-time magic:
//   https://github.com/laher/goxc/wiki/versioning
var (
	VERSION     = "0.1.?"
	BUILD_DATE  = "by something other than goxc"
	SOURCE_DATE = ""
)

// copy them somewhere out of package main
func init() {
	commands.Version.Version = VERSION
	commands.Version.BuildDate = BUILD_DATE
	commands.Version.SourceDate = SOURCE_DATE
}
