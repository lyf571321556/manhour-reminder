package main

import (
	"github.com/lyf571321556/manhour-reminder/cmd"
	"github.com/lyf571321556/manhour-reminder/service"
)

var AppAuth service.AuthInfo

func main() {
	cmd.Execute()
}
