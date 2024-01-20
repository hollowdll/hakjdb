package connect

import (
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultHost string = ""
	defaultPort uint16 = 0
)

var (
	host          string
	port          uint16
	cmdConnectSet = &cobra.Command{
		Use:   "set",
		Short: "Change connection settings",
		Long:  "Change connection settings to a kvdb server",
		Run: func(cmd *cobra.Command, args []string) {
			if host != defaultHost {
				viper.Set(config.ConfigKeyHost, host)
			}
			if port != defaultPort {
				viper.Set(config.ConfigKeyPort, port)
			}
			err := viper.WriteConfig()
			cobra.CheckErr(err)
		},
	}
)

func init() {
	cmdConnectSet.Flags().StringVarP(&host, "host", "a", defaultHost, "server address")
	cmdConnectSet.Flags().Uint16VarP(&port, "port", "p", defaultPort, "port number")
}
