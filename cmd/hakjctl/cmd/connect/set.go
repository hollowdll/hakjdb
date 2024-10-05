package connect

import (
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
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
		Long:  "Change the connection settings used to connect to a HakjDB server. Only sets those that are specified.",
		Example: `# Change the host and port
hakjctl connect set --host 127.0.0.1 --port 9000

# Change only the default database
hakjctl connect set --database default`,
		Run: func(cmd *cobra.Command, args []string) {
			setConnectionSettings()
		},
	}
)

func init() {
	cmdConnectSet.Flags().StringVarP(&host, "host", "a", defaultHost, "Host or IP address")
	cmdConnectSet.Flags().Uint16VarP(&port, "port", "p", defaultPort, "Port number")
	cmdConnectSet.Flags().StringVarP(&dbName, "database", "d", defaultDbName, "Default database to use")
}

func setConnectionSettings() {
	if host != defaultHost {
		config.SetHost(host)
	}
	if port != defaultPort {
		config.SetPort(port)
	}
	if dbName != defaultDbName {
		config.SetDefaultDB(dbName)
	}
	cobra.CheckErr(config.WriteFile())
}
