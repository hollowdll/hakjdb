package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Hostname is the server's address.
	Hostname = "localhost"
	// Port is the server's port number.
	Port       uint16 = 12345
	cmdConnect        = &cobra.Command{
		Use:   "connect [flags]",
		Short: "Change connection settings to a server",
		Long:  "Change connection settings to a server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("connect")
		},
	}
)

func init() {
	cmdConnect.Flags().StringVarP(&Hostname, "host", "a", "localhost", "server address")
	cmdConnect.Flags().Uint16VarP(&Port, "port", "p", 12345, "port number")
}
