package cmd

import (
	"fmt"
	"github.com/lyf571321556/manhour-reminder/log"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "manhour-robot",
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./conf/config.yaml)")
	rootCmd.PersistentFlags().BoolP("daemon", "d", false, "service run on background")
	rootCmd.PersistentFlags().StringP("account", "a", "", "ONES的用户名")
	rootCmd.PersistentFlags().StringP("password", "p", "", "ONES的密码")

	//通过配置文件绑定变量
	viper.BindPFlag("account", rootCmd.PersistentFlags().Lookup("account"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("daemon", rootCmd.PersistentFlags().Lookup("daemon"))
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
		viper.AddConfigPath(currentPath + "/conf")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	daemon := viper.GetBool("daemon")
	if err == nil && !daemon {
		fmt.Println("current config file is:", viper.ConfigFileUsed())
	}
}
