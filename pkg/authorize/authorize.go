package authorize

import (
	"fmt"

	"github.com/bobmaertz/railcar/pkg/storage"
)

// response_type=code&client_id=s6BhdRkqt3&state=xyz
// 	        &redirect_uri=https%3A%2F%2Fclient%2Eexample%2Ecom%2Fcb
type Request struct {
	ResponseType string
	ClientId     string
	State        string
	RedirectUri  string
}

type Authorizer struct {
	backend storage.Backend
	//logger
}

func NewAuthorizer(s storage.Backend) (Authorizer, error) {
	return Authorizer{backend: s}, nil
}

func (a Authorizer) Authorize(req Request) string {

	client, err := a.backend.GetClient(req.ClientId)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	//TODO: return state, with a code..
	fmt.Println(client.Id)
	fmt.Printf("\n\n%v\n", req.ResponseType)
	fmt.Printf("%v\n", req.ClientId)
	fmt.Printf("%v\n", req.State)
	fmt.Printf("%v\n", req.RedirectUri)

	return ""
}
