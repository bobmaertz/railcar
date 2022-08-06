package actions

import (
	"encoding/json"
	"net/http"

	"github.com/bobmaertz/railcar/pkg/build"
	"github.com/bobmaertz/railcar/pkg/storage"
)

func tokenHandler(s storage.Backend) http.HandlerFunc {
	// a, _ := authorize.NewAuthorizer(s)

	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(build.BuildInfo())

	}
}
