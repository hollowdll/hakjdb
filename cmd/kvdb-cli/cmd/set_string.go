package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdSetString = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a string value",
	Long:  "Set a string value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(args[0])
		}
	},
}

func init() {
	cmdSetString.Flags().StringVarP(&dbName, "db", "d", "", "name of the database")
}
