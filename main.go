package main

import (
	"github.com/lyf571321556/manhour-reminder/cmd"
)

var (
	version string
)

func main() {
	cmd.Version = version
	cmd.Execute()
}
