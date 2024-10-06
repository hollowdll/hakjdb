package db

import (
	"context"
	"fmt"
	"time"

	"github.com/hollowdll/hakjdb/api/v1/dbpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var cmdDbInfo = &cobra.Command{
	Use:   "info",
	Short: "Show information about a database",
	Long: `Show information about a database. If the database name is not specified, shows information about the default database.

Meaning of the returned fields:
- name: Name of the database
- description: Description of the database
- created_at: UTC timestamp specifying when the database was created
- updated_at: UTC timestamp specifying when the database was last updated
- key_count: Number of keys stored in the database
- data_size: Size of the stored data in bytes
`,
	Example: `# Use the default database
hakjctl db info

# Specify the database to use
hakjctl db info -n "mydb"`,
	Run: func(cmd *cobra.Command, args []string) {
		showDbInfo()
	},
}

func init() {
	cmdDbInfo.Flags().StringVarP(&dbName, "name", "n", "", "The name of the database")
}

func showDbInfo() {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()
	if len(dbName) < 1 {
		dbName = config.GetDefaultDB()
	}
	resp, err := client.GrpcDBClient.GetDBInfo(ctx, &dbpb.GetDBInfoRequest{DbName: dbName})
	client.CheckGrpcError(err)

	var info string
	info += fmt.Sprintf("name: %s\n", resp.Data.Name)
	info += fmt.Sprintf("description: \"%s\"\n", resp.Data.Description)
	info += fmt.Sprintf("created_at: %s00:00\n", resp.Data.CreatedAt.AsTime().Format(time.RFC3339))
	info += fmt.Sprintf("updated_at: %s00:00\n", resp.Data.UpdatedAt.AsTime().Format(time.RFC3339))
	info += fmt.Sprintf("key_count: %d\n", resp.Data.GetKeyCount())
	info += fmt.Sprintf("data_size: %dB", resp.Data.GetDataSize())

	fmt.Println(info)
}
