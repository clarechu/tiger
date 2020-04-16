package pkg

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "tiger",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

/*

*/
func NewRoot() *cobra.Command {
	return rootCmd
}
/*
type Root struct {
	SubCommand []*cobra.Command `json:"subCommand"`
}

func (root *Root) Registry(subCommand *cobra.Command) {
	root.SubCommand = append(root.SubCommand, subCommand)
}

func (root *Root) Notify(rootCommand *cobra.Command) {
	for _, command := range root.SubCommand {
		rootCommand.AddCommand(command)
	}
}
*/