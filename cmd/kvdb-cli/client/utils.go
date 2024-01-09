package client

import (
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// CheckGrpcError checks if error is a gRPC error.
// Prints the gRPC status message if it is. Otherwise prints the error.
func CheckGrpcError(err error) {
	if err != nil {
		if st, ok := status.FromError(err); ok {
			cobra.CheckErr(fmt.Sprintf("response from server: %s", st.Message()))
		} else {
			cobra.CheckErr(err)
		}
	}
}
