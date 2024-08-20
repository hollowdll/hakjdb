package db

import (
	"context"
	"fmt"

	"github.com/hollowdll/kvdb/api/v0/dbpb"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/client"
	"github.com/hollowdll/kvdb/cmd/kvdbctl/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
)

var (
	cmdDbChange = &cobra.Command{
		Use:   "change NAME",
		Short: "Change a database",
		Long:  "Change the metadata of the specified database. Returns the name of the changed database.",
		Example: `# Change the name of database 'mydb'
kvdbctl db change mydb --name "my-new-db"

# Change the description of database 'mydb'
kvdbctl db change mydb --description "New database description."

# Change the name and description of database 'mydb'
kvdbctl db change mydb -n "my-new-db" -d "New database description."`,
		Args: cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			changeName := cmd.Flags().Changed("name")
			changeDesc := cmd.Flags().Changed("description")
			changeDatabaseMetadata(args[0], changeName, changeDesc)
		},
	}
	newDesc string
	newName string
)

func init() {
	cmdDbChange.Flags().StringVarP(&newName, "name", "n", "", "New name of the database")
	cmdDbChange.Flags().StringVarP(&newDesc, "description", "d", "", "New description of the database")
}

func changeDatabaseMetadata(dbName string, changeName bool, changeDesc bool) {
	ctx := metadata.NewOutgoingContext(context.Background(), client.GetBaseGrpcMetadata())
	ctx, cancel := context.WithTimeout(ctx, config.GetCmdTimeout())
	defer cancel()

	req := &dbpb.ChangeDBRequest{
		DbName:            dbName,
		NewName:           newName,
		ChangeName:        changeName,
		NewDescription:    newDesc,
		ChangeDescription: changeDesc,
	}
	res, err := client.GrpcDBClient.ChangeDB(ctx, req)
	client.CheckGrpcError(err)
	fmt.Println(res.DbName)
}
