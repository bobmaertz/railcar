package authorize

import (
	"crypto/rand"
	"fmt"
	"net/url"
	"time"

    "log/slog"
	oauthErr "github.com/bobmaertz/railcar/pkg/error"
	"github.com/bobmaertz/railcar/pkg/storage"
)

const (
	AuthorizationCode = "code"
	ClientCredentials = "client_credentials"
)

type Request struct {
	ResponseType string
	ClientId     string
	State        string
	RedirectUri  string
}

type Authorizer struct {
	backend              storage.Backend
	generateAuthCodeFunc func() (string, error)
}

func NewAuthorizer(s storage.Backend) (Authorizer, error) {
	return Authorizer{backend: s, generateAuthCodeFunc: defaultAuthCodeGenerator}, nil
}

func (a *Authorizer) Authorize(req Request) (string, *oauthErr.OAuthError) {

	//Validate that the client exists
	client, err := a.backend.GetClient(req.ClientId)
	if err != nil {
		slog.Warn("Authorize: client is not authorized: %v", err)
		return "", oauthErr.Errors["unauthorized_client"]
	}

	switch req.ResponseType {
	case AuthorizationCode:
		return a.processAuthCodeRequest(client, req)
	default:
		return "", oauthErr.Errors["invalid_request"]
	}
}

func (a *Authorizer) processAuthCodeRequest(client storage.Client, req Request) (string, *oauthErr.OAuthError) {
	//TODO: The spec stays this is optional but leave mandatory for now.
	if !contains(client.RedirectUris, req.RedirectUri) {
		return "", oauthErr.Errors["invalid_request"]
	}

	//TODO: The only reason we have to reference Authorizer is for this function and the logging,
	// can this be removed so this function can operate independantly?
	code, err := a.generateAuthCodeFunc()
	if err != nil {
		slog.Error("processAuthCodeRequest: unable to create authorization code: %v", err)
		return "", oauthErr.Errors["server_error"]
	}

	err = a.backend.CreateAuthorizationCode(code, client, time.Now().Add(10*time.Minute))
	if err != nil {
		slog.Error("processAuthCodeRequest: unable to save authorization code: %v", err)
		return "", oauthErr.Errors["server_error"]
	}

	queries := map[string]string{
		"code":  code,
		"state": req.State,
	}

	redirect, err := createRedirectUrl(req.RedirectUri, queries)
	if err != nil {
		slog.Error("processAuthCodeRequest: unable to create redirect url: %v", err)
		return "", oauthErr.Errors["server_error"]
	}

	return redirect, nil
}

func defaultAuthCodeGenerator() (string, error) {
	n := 16
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	s := fmt.Sprintf("%x", b)

	return s, nil
}

func createRedirectUrl(baseUrl string, queries map[string]string) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range queries {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}


func contains(arr []string, target string) bool {
    for _, e := range arr {
        if e == target {
            return true
        }
    }
    return false
}
