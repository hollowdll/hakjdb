package db

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbName string
var cmdCreateDb = &cobra.Command{
	Use:   "create [flags]",
	Short: "Create a new database",
	Long:  "Create a new database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database name:", dbName)
	},
}

func init() {
	cmdCreateDb.Flags().StringVarP(&dbName, "name", "n", "", "name of the database (required)")
	cmdCreateDb.MarkFlagRequired("name")
}
