package nodeconfig

import (
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

type Environment struct {
	logger scribe.Logger
}

func NewEnvironment(logger scribe.Logger) Environment {
	return Environment{
		logger: logger,
	}
}

func (e Environment) Configure(buildEnv packit.Environment) error {
	e.logger.Process("Configuring NodeJS Environment for Railway")

	buildEnv.Default("NPM_CONFIG_PRODUCTION", "false")

	e.logger.Detail(scribe.NewFormattedMapFromEnvironment(buildEnv).String())

	return nil
}
