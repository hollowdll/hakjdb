package connect

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultHost string = ""
	defaultPort uint16 = 0
)

var (
	host       string
	port       uint16
	CmdConnect = &cobra.Command{
		Use:   "connect [flags]",
		Short: "Change connection settings to a server",
		Long:  "Change connection settings to a server",
		Run: func(cmd *cobra.Command, args []string) {
			if host != defaultHost {
				viper.Set("host", host)
			}
			if port != defaultPort {
				viper.Set("port", port)
			}
			err := viper.WriteConfig()
			cobra.CheckErr(err)
		},
	}
)

func init() {
	CmdConnect.AddCommand(CmdConnectLs)
	CmdConnect.Flags().StringVarP(&host, "host", "a", defaultHost, "server address")
	CmdConnect.Flags().Uint16VarP(&port, "port", "p", defaultPort, "port number")
}
