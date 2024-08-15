package client

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// CheckGrpcError checks if error is a gRPC error.
// Prints error with the gRPC status message if it is.
// Otherwise prints the error if it is not nil.
func CheckGrpcError(err error) {
	if err != nil {
		if st, ok := status.FromError(err); ok {
			cobra.CheckErr(fmt.Sprintf("response from server: %s", st.Message()))
		} else {
			cobra.CheckErr(err)
		}
	}
}

// Prompts user a confirm message and reads input.
// The input should be of type Yes/No.
// Returns true if the user entered Yes.
func PromptConfirm(msg string) bool {
	var input string
	fmt.Printf(msg)
	_, err := fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}
	cobra.CheckErr(err)

	return strings.ToLower(input) == "yes"
}
