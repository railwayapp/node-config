package nodeconfig

import (
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/scribe"
)

func Build(environment Environment, logger scribe.Logger) packit.BuildFunc {
	return func(context packit.BuildContext) (packit.BuildResult, error) {
		logger.Title("%s %s", context.BuildpackInfo.Name, context.BuildpackInfo.Version)

		layer, err := context.Layers.Get("nodeconfig")
		if err != nil {
			return packit.BuildResult{}, err

		}

		environment.Configure(layer.BuildEnv)

		logger.Break()

		return packit.BuildResult{}, nil
	}
}
