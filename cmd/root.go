package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

//type Config struct {
//	UserInfo struct {
//		age int    `mapstructure:"age"`
//		sex string `mapstructure:"sex"`
//	} `mapstructure:"user"`
//	Version string `mapstructure:"version"`
//	Debug   bool   `mapstructure:"debug"`
//}

var rootcmd = &cobra.Command{
	Use:   "apiserver",
	Short: "apiserver",
	Long:  `a fast api server with cobra and viper`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root running")
	},
}

func Execute() {
	if error := rootcmd.Execute(); error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}
	fmt.Println("config file: ", cfgfile)

}

func init() {
	fmt.Println("root command init started")

	//PersistentFlag全局参数，可以被子命令一起被使用
	rootcmd.PersistentFlags().StringVar(&cfgfile, "config", "c", "config file(default is $LOCAL/config/config.yaml")

	//不能从这里读进config文件，文件名读取错误。init()函数在main函数之前运行，init函数是没有参数，也没有输入的。
	//initconfig(cfgfile)

	//cobra设定函数流程之一，设定在execute函数之前执行的init操作，这个时候是可以解析到flag的
	cobra.OnInitialize(initconfig)
}

func initconfig() {
	if cfgfile != "" {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(cfgfile)
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("using config file:", viper.ConfigFileUsed(), "err:", err)
	}
	fmt.Println("age:", viper.GetInt("user.age"))
	fmt.Println(viper.IsSet("user.age"))
	fmt.Println("sex:", viper.GetString("user.sex"))
	fmt.Println(viper.IsSet("user.sex"))
	fmt.Println("version:", viper.GetString("version"))
	fmt.Println(viper.IsSet("version"))

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("unmarshal config err:", err)
	}

	fmt.Println(cfg)
}
