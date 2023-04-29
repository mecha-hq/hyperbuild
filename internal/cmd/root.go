package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/mecha-ci/hyperbuild/internal/app"
	cobrax "github.com/mecha-ci/hyperbuild/internal/x/cobra"
)

type RootCommand struct {
	*cobra.Command
}

func NewRootCommand(ctr *app.Container) *RootCommand {
	const envPrefix = ""

	root := &RootCommand{
		Command: &cobra.Command{
			PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
				cobrax.BindFlags(cmd, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

				return nil
			},
			Use:           "hyperbuild",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
	}

	cobrax.BindFlags(root.Command, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

	root.AddCommand(NewVersionCommand(ctr))
	root.AddCommand(NewRunCommand(ctr))

	return root
}
