package connect

import (
	"github.com/hollowdll/kvdb/cmd/kvdb-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultHost   string = ""
	defaultPort   uint16 = 0
	defaultDbName string = ""
)

var (
	host          string
	port          uint16
	dbName        string
	cmdConnectSet = &cobra.Command{
		Use:   "set",
		Short: "Change connection settings",
		Long:  "Change connection settings to a kvdb server",
		Run: func(cmd *cobra.Command, args []string) {
			setConnectionSettings()
		},
	}
)

func init() {
	cmdConnectSet.Flags().StringVarP(&host, "host", "a", defaultHost, "host name or IP address")
	cmdConnectSet.Flags().Uint16VarP(&port, "port", "p", defaultPort, "port number")
	cmdConnectSet.Flags().StringVarP(&dbName, "db", "d", defaultDbName, "database to use")
}

func setConnectionSettings() {
	if host != defaultHost {
		viper.Set(config.ConfigKeyHost, host)
	}
	if port != defaultPort {
		viper.Set(config.ConfigKeyPort, port)
	}
	if dbName != defaultDbName {
		viper.Set(config.ConfigKeyDatabase, dbName)
	}
	err := viper.WriteConfig()
	cobra.CheckErr(err)
}
