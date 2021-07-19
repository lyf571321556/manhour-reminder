package main

import (
	"github.com/lyf571321556/manhour-reminder/cmd"
	"github.com/lyf571321556/manhour-reminder/service"
	"log"
)

func init() {
	log.SetPrefix("Robot: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

var AppAuth service.AuthInfo

func main() {
	cmd.Execute()
	//flag.Parse()
	//flag.Lookup("logtostderr").Value.Set("true")
	//glog.Info("hi_b")
	//flag.Lookup("log_dir").Value.Set("./log")
}
