package version

var (
	// BuildTime is a time label from when the binary was built.
	BuildTime = "unset"
	// Commit is the git hash from when the binary was built
	Commit = "unset"
	// Release is the semantic version of the current build
	Release = "unset"
)
