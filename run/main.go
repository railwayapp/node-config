package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
	nodeConfig "github.com/railwayapp/node-config"
)

func main() {
	logger := scribe.NewLogger(os.Stdout)

	environment := nodeConfig.NewEnvironment(logger)

	packit.Run(
		nodeConfig.Detect(),
		nodeConfig.Build(environment, logger),
	)
}
