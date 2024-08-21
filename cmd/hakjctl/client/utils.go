package client

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/hollowdll/hakjdb/internal/common"
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

func WriteTokenCache(filePath string, token string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(token))
	if err != nil {
		return err
	}

	return nil
}

func ReadTokenFromCache(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	token, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(token), nil
}

func GetTokenCacheFilePath() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	subCacheDir := filepath.Join(cacheDir, config.CacheDirSubDirName)
	err = common.CreateDirectoriesInDirPath(subCacheDir)
	if err != nil {
		return "", err
	}
	return filepath.Join(subCacheDir, config.TokenCacheFileName), nil
}

func createEmptyTokenCache() {
	path, err := GetTokenCacheFilePath()
	cobra.CheckErr(err)
	cobra.CheckErr(common.CreateFileIfNotExist(path))
}
