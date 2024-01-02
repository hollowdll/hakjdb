package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdGetString = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a string value",
	Long:  "Get a string value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(args[0])
		}
	},
}

func init() {
	cmdGetString.Flags().StringVarP(&dbName, "db", "d", "", "name of the database")
}
