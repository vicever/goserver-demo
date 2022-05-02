package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add person",
	Short: "add one item",
	Long:  "add one item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add commmand running ")
	},
}

func init() {
	rootcmd.AddCommand(add)
}
