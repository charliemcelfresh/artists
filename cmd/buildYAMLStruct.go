package cmd

import (
	"artists/internal/translationbuilder"
	"artists/internal/translations"

	"github.com/spf13/cobra"
)

var (
	buildYAMLStructCmd *cobra.Command = &cobra.Command{
		Use:   translations.StringValues.Commands.BuildYamlStruct.Use,
		Short: translations.StringValues.Commands.BuildYamlStruct.Short,
		RunE:  BuildYAMLStructCommandRunner,
	}
)

func init() {
	rootCmd.AddCommand(buildYAMLStructCmd)
}

func BuildYAMLStructCommandRunner(cmd *cobra.Command, args []string) error {
	return translationbuilder.BuildYAMLStructDefinition(args[0])
}
