package version

import (
	"fmt"
	"runtime"
)

var (
	buildVersion     = "unknown"
	buildGitRevision = "unknown"
	buildStatus      = "unknown"
	buildTag         = "unknown"

	// Info exports the build version information.
	Info BuildInfo
)

// BuildInfo describes version information about the binary build.
type BuildInfo struct {
	Version     string
	GitRevision string
	BuildStatus string
	GitTag      string
	GoVersion   string
	GoArch      string
	OperatorSDK string
}

func (b BuildInfo) String() string {
	return fmt.Sprintf("%#v", b)
}

func init() {
	Info = BuildInfo{
		Version:     buildVersion,
		GitRevision: buildGitRevision,
		BuildStatus: buildStatus,
		GitTag:      buildTag,
		GoVersion:   runtime.Version(),
		GoArch:      fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
