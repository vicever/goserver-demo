package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config file",
	Short: "be a config file",
	Long:  "be a config file",
	Args:  cobra.ExactArgs(1),
	Run:   GetConfig,
}

func GetConfig(cmd *cobra.Command, args []string) {
	fmt.Println("configcmd running and get args:", args)
}

func init() {

	rootcmd.AddCommand(configCmd)
}
