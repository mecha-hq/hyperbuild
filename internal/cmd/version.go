package cmd

import (
	"fmt"

	"github.com/omissis/hyperbuild/internal/app"
	"github.com/spf13/cobra"
)

func NewVersionCommand(ctr *app.Container) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display version information about benth",
		Args:  cobra.ExactArgs(0),
		RunE: func(_ *cobra.Command, _ []string) error {
			fmt.Printf("BuildTime: %s\n", ctr.Versions.BuildTime)
			fmt.Printf("GitCommit: %s\n", ctr.Versions.GitCommit)
			fmt.Printf("GoVersion: %s\n", ctr.Versions.GoVersion)
			fmt.Printf("OsArch: %s\n", ctr.Versions.OsArch)
			fmt.Printf("Version: %s\n", ctr.Versions.Version)

			return nil
		},
	}
}
