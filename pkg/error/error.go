package error

import "fmt"

//Errors - just getting these into the codebase, they dont belong here.
var Errors = map[string]*OAuthError{
	"invalid_request":           {Type: "invalid_request"},
	"unauthorized_client":       {Type: "unauthorized_client"},
	"access_denied":             {Type: "access_denied"},
	"unsupported_response_type": {Type: "unsupported_response_type"},
	"invalid_scope":             {Type: "invalid_scope"},
	"server_error":              {Type: "server_error"},
	"temporarily_unavailable":   {Type: "temporarily_unavailable"},
}

type OAuthError struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	URI         string `json:"error_uri"`
}

func (e OAuthError) Error() string {
	return fmt.Sprintf("type: %v, description: %v", e.Type, e.Description)
}
