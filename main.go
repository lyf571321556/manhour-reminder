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
	fmt.Println("Build Time:\t", Time)
	fmt.Println("Build User:\t", User)
	cmd.Execute()
}
