package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cfgfile string
var cfg = Config{}

type Config struct {
	User struct {
		Age int
		Sex string
	}
	Version      string
	Debug        bool
	redis_option RedisOptions
}

var configCmd = &cobra.Command{
	Use:   "config file",
	Short: "be a config file",
	Long:  "be a config file",
	Args:  cobra.ExactArgs(1),
	Run:   GetConfig,
}

// RedisOptions defines options for redis cluster.
type RedisOptions struct {
	Host                  string   `json:"host"                     mapstructure:"host"                     description:"Redis service host address"`
	Port                  int      `json:"port"`
	Addrs                 []string `json:"addrs"                    mapstructure:"addrs"`
	Username              string   `json:"username"                 mapstructure:"username"`
	Password              string   `json:"password"                 mapstructure:"password"`
	Database              int      `json:"database"                 mapstructure:"database"`
	MasterName            string   `json:"master-name"              mapstructure:"master-name"`
	MaxIdle               int      `json:"optimisation-max-idle"    mapstructure:"optimisation-max-idle"`
	MaxActive             int      `json:"optimisation-max-active"  mapstructure:"optimisation-max-active"`
	Timeout               int      `json:"timeout"                  mapstructure:"timeout"`
	EnableCluster         bool     `json:"enable-cluster"           mapstructure:"enable-cluster"`
	UseSSL                bool     `json:"use-ssl"                  mapstructure:"use-ssl"`
	SSLInsecureSkipVerify bool     `json:"ssl-insecure-skip-verify" mapstructure:"ssl-insecure-skip-verify"`
}

// NewRedisOptions create a `zero` value instance.
func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:                  "127.0.0.1",
		Port:                  6379,
		Addrs:                 []string{},
		Username:              "",
		Password:              "",
		Database:              0,
		MasterName:            "",
		MaxIdle:               2000,
		MaxActive:             4000,
		Timeout:               0,
		EnableCluster:         false,
		UseSSL:                false,
		SSLInsecureSkipVerify: false,
	}
}

func GetConfig(cmd *cobra.Command, args []string) {
	fmt.Println("configcmd running and get args:", args)
}

func init() {

	rootcmd.AddCommand(configCmd)
}
