package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	hostname   string
	port       uint16
	cmdConnect = &cobra.Command{
		Use:   "connect [flags]",
		Short: "Change connection settings to a server",
		Long:  "Change connection settings to a server",
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set("host", hostname)
			viper.Set("port", port)
			err := viper.WriteConfig()
			cobra.CheckErr(err)
		},
	}
)

func init() {
	cmdConnect.Flags().StringVarP(&hostname, "host", "a", viper.GetString("host"), "server address")
	cmdConnect.Flags().Uint16VarP(&port, "port", "p", viper.GetUint16("port"), "port number")
}
