package main

import (
	"github.com/lyf571321556/manhour-reminder/bot"
	"github.com/lyf571321556/manhour-reminder/cmd"
	"github.com/lyf571321556/manhour-reminder/config"
	"github.com/lyf571321556/manhour-reminder/service"
	"log"
)

func init() {
	log.SetPrefix("Rebot: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	config.Init()
	bot.InitBot()
}

var AppAuth service.AuthInfo

func main() {
	cmd.Execute()
	//flag.Parse()
	//flag.Lookup("logtostderr").Value.Set("true")
	//glog.Info("hi_b")
	//flag.Lookup("log_dir").Value.Set("./log")
}
