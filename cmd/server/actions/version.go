package actions

import (
	"encoding/json"
	"net/http"

	"github.com/bobmaertz/railcar/pkg/build"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(build.BuildInfo())
}
