package main

import (
	"log"

	"github.com/mecha-ci/hyperbuild/internal/app"
	"github.com/mecha-ci/hyperbuild/internal/cmd"
)

var (
	version   = "unknown"
	gitCommit = "unknown"
	buildTime = "unknown"
	goVersion = "unknown"
	osArch    = "unknown"
)

func main() {
	ctr := app.NewContainer(app.Versions{
		BuildTime: buildTime,
		GitCommit: gitCommit,
		GoVersion: goVersion,
		OsArch:    osArch,
		Version:   version,
	})

	if err := cmd.NewRootCommand(ctr).Execute(); err != nil {
		log.Fatal(err)
	}
}
