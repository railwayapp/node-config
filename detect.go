package nodeconfig

import (
	"github.com/paketo-buildpacks/packit"
)

type BuildPlanMetadata struct {
	Build bool `toml:"build"`
}

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Requires: []packit.BuildPlanRequirement{
					{
						Name:     "node",
						Metadata: BuildPlanMetadata{Build: true},
					},
				},
			},
		}, nil
	}
}
