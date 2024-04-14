package actions

import (
	"net/http"

    "fmt"

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

        fmt.Println("Params:" , params)
		uri, oauthErr := a.Authorize(req)
		if oauthErr != nil {
			//TODO: Handle this
            fmt.Println("Error")
		}
		a, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			//TODO: Handle this
		}

		http.Redirect(w, a, uri, http.StatusFound)
	}
}
