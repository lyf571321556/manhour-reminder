package cmd

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/conf"
	"github.com/lyf571321556/manhour-reminder/log"
	"github.com/lyf571321556/manhour-reminder/robot"
	"github.com/lyf571321556/manhour-reminder/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start man-hour robot.",
	Long:  `start man-hour robot.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemon, _ := cmd.Flags().GetBool("daemon")
		//user, _ := cmd.Flags().GetString("user")
		//password, _ := cmd.Flags().GetString("password")
		user := viper.GetString("account")
		password := viper.GetString("password")
		if daemon {
			command := exec.Command("./manhour-reminder", fmt.Sprintf("--config=%s", viper.ConfigFileUsed()), "start", fmt.Sprintf("-a=%s", user), fmt.Sprintf("-p=%s", password)) //go run main.go start
			command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true, Foreground: false}
			command.Stdout = os.Stdout
			command.Stdin = os.Stdin
			err := command.Start()
			if err != nil {
				fmt.Println("service start occur err ", err)
				os.Exit(0)
				return
			}
			log.Info(fmt.Sprintf("service start, [PID] %d running...\n", command.Process.Pid))
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
	if err := robot.InitBot(viper.ConfigFileUsed()); err != nil {
		fmt.Println(err)
		return
	}
	//获取token
	loginUrl := fmt.Sprintf("%s%s", conf.AppConfig.OnesProjectUrl, service.AUTH_LOGIN)
	AppAuth, err := service.Login(loginUrl, user, password)
	if err != nil {
		log.Fatal(err.Error())
	}

	//支持秒级(可选)的cron表达式
	//c := cron.New(cron.WithParser(cron.NewParser(
	//	cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	//)))
	//c.AddFunc(conf.AppConfig.TaskCrontab, func() {
	robot.StartCheckUsersManhourInEveryRobot(AppAuth)
	//})
	//c.Start()
	//select {}
}
