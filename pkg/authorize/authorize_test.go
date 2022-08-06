package authorize

import (
	"os"
	"testing"

	oauthError "github.com/bobmaertz/railcar/pkg/error"
	"github.com/bobmaertz/railcar/pkg/storage"
	"github.com/bobmaertz/railcar/pkg/storage/memory"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizer_Authorize(t *testing.T) {

	var l = logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	type fields struct {
		backend storage.Backend
		log     logrus.Logger
	}
	type args struct {
		req Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
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
				log: l,
			},
			args: args{
				req: Request{
					RedirectUri:  "http://localhost",
					State:        "abcdef1234",
					ClientId:     "abcd",
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
				log: l,
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
				log: l,
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
				log: l,
			},
			args: args{
				req: Request{
					RedirectUri:  "http://invalid_redirect_url/callback",
					State:        "abcdef1234",
					ClientId:     "abcd",
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
				log: l,
			},
			args: args{
				req: Request{
					RedirectUri:  "http://localhost",
					State:        "abcdef1234",
					ClientId:     "abcd",
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
				backend: tt.fields.backend,
				log:     tt.fields.log,
			}
			got, err := a.Authorize(tt.args.req)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}
			assert.Nil(t, err)
			if got != tt.want {
				t.Errorf("Authorizer.Authorize() = %v, want %v", got, tt.want)
			}
		})
	}
}
