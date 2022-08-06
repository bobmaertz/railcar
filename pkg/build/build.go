package build

import "fmt"

var (
	sha1      string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

//TODO: Fix me, need to get ldflags working correctly
func BuildInfo() string {
	return fmt.Sprintf("sha1: %s, buildTime: %v", sha1, buildTime)
}
