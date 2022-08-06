package actions

import (
	"encoding/json"
	"net/http"

	"github.com/bobmaertz/railcar/pkg/authorize"
	"github.com/bobmaertz/railcar/pkg/storage"
)

func authorizeHandler(s storage.Backend) http.HandlerFunc {
	a, _ := authorize.NewAuthorizer(s)

	return func(w http.ResponseWriter, r *http.Request) {

		params := r.URL.Query()
		req := authorize.Request{
			ClientId:     params.Get("client_id"),
			ResponseType: params.Get("response_type"),
			State:        params.Get("state"),
			RedirectUri:  params.Get("redirect_uri"),
		}

		a.Authorize(req)

		json.NewEncoder(w).Encode("authorize")

		//redirect to callback url
	}
}
