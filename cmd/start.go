package cmd

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/bot"
	"github.com/lyf571321556/manhour-reminder/config"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start man-hour rebot.",
	Long:  `start man-hour rebot.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon, _ := cmd.Flags().GetBool("daemon")
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		if daemon {
			command := exec.Command("./manhour-reminder", "start", fmt.Sprintf("-u=%s", user), fmt.Sprintf("-p=%s", password)) //go run main.go start
			command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Foreground: false}
			command.Stdout = os.Stdout
			command.Stdin = os.Stdin
			err := command.Start()
			if err != nil {
				fmt.Println("service start occur err ", err)
				os.Exit(0)
				return
			}
			log.Printf("service start, [PID] %d running...\n", command.Process.Pid)
			ioutil.WriteFile(".pid.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
			os.Exit(0)
		}
		startServer(user, password)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func startServer(user string, password string) {
	//获取token
	loginUrl := fmt.Sprintf("%s%s", config.AppConfig.OnesProjectUrl, service.AUTH_LOGIN)
	AppAuth, err := service.Login(loginUrl, user, password)
	if err != nil {
		log.Fatal(err)
	}

	c := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))
	c.AddFunc(config.AppConfig.TaskCrontab, func() {
		go bot.SendMsgToUser(AppAuth)
	})
	c.Start()
	select {}
}
