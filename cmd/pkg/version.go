package pkg

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	SetCommand(versionCmd)
	rootCmd.AddCommand(versionCmd)
}

const (
	Version string = "v1.0"
)

var (
	version string
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tigger",
	Args:  Args,
	Long:  `All software has versions. This is Tigger's`,
	Run:   Run,
}

func Args(cmd *cobra.Command, args []string) error {
	return nil
}

func SetCommand(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&version, "verbose", "v", Version, "当前版本号")
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	fmt.Printf("Version: %v", version)
}
