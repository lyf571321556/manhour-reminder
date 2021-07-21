package main

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/cmd"
)

var (
	Time string
	User string
)

func main() {
	fmt.Println("build.Time:\t", Time)
	fmt.Println("build.User:\t", User)
	cmd.Execute()
}
