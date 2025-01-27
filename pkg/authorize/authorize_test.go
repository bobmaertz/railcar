package authorize

import (
	"testing"

	"github.com/bobmaertz/railcar/internal/assert"
	oauthError "github.com/bobmaertz/railcar/pkg/error"
	"github.com/bobmaertz/railcar/pkg/storage"
	"github.com/bobmaertz/railcar/pkg/storage/memory"
)

func TestAuthorizer_Authorize(t *testing.T) {
	type fields struct {
		backend              storage.Backend
		generateAuthCodeFunc func() (string, error)
	}
	type args struct {
		req Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr *oauthError.OAuthError
	}{
		{
			name: "happy path, authorization code is returned",
			fields: fields{
				backend: func() storage.Backend {
					mem, err := memory.NewMemory()
					if err != nil {
						t.Error(err)
					}

					return mem
				}(),
				generateAuthCodeFunc: func() (string, error) {
					return "abcd", nil
				},
			},
			args: args{
				req: Request{
					RedirectUri:  "http://localhost",
					State:        "abcdef1234",
					ClientId:     "29352735982374239857",
					ResponseType: "code",
				},
			},
			want:    "http://localhost?code=abcd&state=abcdef1234",
			wantErr: nil,
		},
		{
			name: "no client is supplied, error is returned",
			fields: fields{
				backend: func() storage.Backend {
					mem, err := memory.NewMemory()
					if err != nil {
						t.Error(err)
					}
					return mem
				}(),
			},
			args: args{
				req: Request{
					RedirectUri:  "http://localhost",
					State:        "abcdef1234",
					ResponseType: "code",
				},
			},
			want:    "",
			wantErr: oauthError.Errors["unauthorized_client"],
		},
		{
			name: "invalid client is supplied, unauthorized client error is returned",
			fields: fields{
				backend: func() storage.Backend {
					mem, err := memory.NewMemory()
					if err != nil {
						t.Error(err)
					}
					return mem
				}(),
			},
			args: args{
				req: Request{
					RedirectUri:  "http://localhost",
					ClientId:     "doesnt_exist",
					State:        "abcdef1234",
					ResponseType: "code",
				},
			},
			want:    "",
			wantErr: oauthError.Errors["unauthorized_client"],
		},
		{
			name: "incorrect redirect_uri is supplied, invalid_request is returned",
			fields: fields{
				backend: func() storage.Backend {
					mem, err := memory.NewMemory()
					if err != nil {
						t.Error(err)
					}
					return mem
				}(),
			},
			args: args{
				req: Request{
					RedirectUri:  "http://invalid_redirect_url/callback",
					State:        "abcdef1234",
					ClientId:     "29352735982374239857",
					ResponseType: "code",
				},
			},
			want:    "",
			wantErr: oauthError.Errors["invalid_request"],
		},
		{
			name: "unsupported grant type is supplied, invalid_request is returned",
			fields: fields{
				backend: func() storage.Backend {
					mem, err := memory.NewMemory()
					if err != nil {
						t.Error(err)
					}
					return mem
				}(),
			},
			args: args{
				req: Request{
					RedirectUri:  "http://localhost",
					State:        "abcdef1234",
					ClientId:     "29352735982374239857",
					ResponseType: "unsupported",
				},
			},
			want:    "",
			wantErr: oauthError.Errors["invalid_request"],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Authorizer{
				backend:              tt.fields.backend,
				generateAuthCodeFunc: tt.fields.generateAuthCodeFunc,
			}
			got, err := a.Authorize(tt.args.req)

			if tt.wantErr != nil {
				assert.AssertEqual(t, err, tt.wantErr)
				return
			}
			assert.AssertNil(t, err)
			if got != tt.want {
				t.Errorf("Authorizer.Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}


