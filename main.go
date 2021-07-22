package main

import (
	"github.com/lyf571321556/manhour-reminder/cmd"
)

var (
	time string
	user string
)

func main() {
	cmd.Time = time
	cmd.User = user
	cmd.Execute()
}
