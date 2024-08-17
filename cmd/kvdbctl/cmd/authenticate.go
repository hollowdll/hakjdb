package cmd

import "github.com/spf13/cobra"

var (
	cmdAuthenticate = &cobra.Command{
		Use:   "authenticate",
		Short: "Authenticate to the server",
		Long:  "Authenticate to the server using password. If no options provided, prompts the user to enter password.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	passedPassword        string
	isReadPasswordFromEnv bool
)

func init() {
	cmdAuthenticate.Flags().StringVarP(&passedPassword, "password", "p", "", "The password to use")
	cmdAuthenticate.Flags().BoolVar(&isReadPasswordFromEnv, "password-from-env", false, "Read password from environment variable KVDBCTL_PASSWORD")
}

func promptPassword() {

}

func authenticate() {

}

func readPasswordFromEnv() {

}
