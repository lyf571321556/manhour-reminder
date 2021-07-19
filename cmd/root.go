package cmd

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/bot"
	"github.com/lyf571321556/manhour-reminder/config"
	"github.com/lyf571321556/manhour-reminder/log"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "manhour-reminder",
	Short: "",
	Long:  ``,
}

func Execute() {
	if log.Logger != nil {
		defer log.Logger.Sync()
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/config.yaml)")
	rootCmd.PersistentFlags().BoolP("daemon", "d", false, "service run on background?")
	rootCmd.PersistentFlags().StringP("account", "a", "", "ONES的用户名")
	rootCmd.PersistentFlags().StringP("password", "p", "", "ONES的密码")

	//通过配置文件绑定变量
	viper.BindPFlag("account", rootCmd.PersistentFlags().Lookup("account"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := homedir.Dir()
		currentPath, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in current directory with name "cofnig" (without extension).
		viper.AddConfigPath(currentPath + "/config")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	if err := config.Init(viper.ConfigFileUsed()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := log.InitLog(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bot.InitBot()
}
