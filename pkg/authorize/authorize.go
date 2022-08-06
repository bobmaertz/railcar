package authorize

import (
	"net/url"
	"time"

	oauthErr "github.com/bobmaertz/railcar/pkg/error"
	"github.com/bobmaertz/railcar/pkg/storage"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type Request struct {
	ResponseType string
	ClientId     string
	State        string
	RedirectUri  string
}

const (
	AuthorizationCode = "code"
	ClientCredentials = "client_credentials"
)

type Authorizer struct {
	backend storage.Backend
	log     logrus.Logger //TODO:
}

func NewAuthorizer(s storage.Backend) (Authorizer, error) {
	return Authorizer{backend: s}, nil
}

func (a Authorizer) Authorize(req Request) (string, *oauthErr.OAuthError) {

	//Validate that the client exists
	client, err := a.backend.GetClient(req.ClientId)
	if err != nil {
		a.log.Warnf("Authorize: client is not authorized: %v", err)
		return "", oauthErr.Errors["unauthorized_client"]
	}

	switch req.ResponseType {
	case AuthorizationCode:
		return a.processAuthCodeRequest(client, req)
	default:
		return "", oauthErr.Errors["invalid_request"]
	}

}

func (a Authorizer) processAuthCodeRequest(client storage.Client, req Request) (string, *oauthErr.OAuthError) {
	//TODO: The spec stays this is optional but leave mandatory for now.
	if !slices.Contains(client.RedirectUris, req.RedirectUri) {
		return "", oauthErr.Errors["invalid_request"]
	}

	code := generateAuthCode()
	err := a.backend.CreateAuthorizationCode(code, client, time.Now().Add(10*time.Minute))
	if err != nil {
		a.log.Errorf("processAuthCodeRequest: unable to save authorization code: %v", err)
		return "", oauthErr.Errors["server_error"]
	}

	queries := map[string]string{
		"code":  code,
		"state": req.State,
	}

	redirect, err := createRedirectUrl(req.RedirectUri, queries)
	if err != nil {
		a.log.Errorf("processAuthCodeRequest: unable to create redirect url: %v", err)
		return "", oauthErr.Errors["server_error"]
	}

	return redirect, nil
}

func generateAuthCode() string {
	//Generate random string ..
	return "abcd"
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
