package auth

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/emicklei/go-restful"
	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
)

const (
	APIKey = "apikey"
)

func createModelHandler() *httptest.Server {
	sh := NewServiceHandler(APIKey, true)
	ts := sh.Register("")

	return httptest.NewServer(ts)
}

func TestServiceHandler_Register(t *testing.T) {
	type args struct {
		urlRoot string
	}
	tests := []struct {
		name string
		p    ServiceHandler
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ServiceHandler{}
			if got := p.Register(tt.args.urlRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServiceHandler.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func genReq(apiKey string) *resty.Request {
	return resty.R().SetHeader("Content-Type", restful.MIME_JSON).SetAuthToken(apiKey)
}

func TestAPI(t *testing.T) {
	// A simple create model test case
	ts := createModelHandler()
	defer ts.Close()
	res, err := genReq(APIKey).Post(ts.URL + APIVersion + RegisterPath)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode())
}

func TestHTTPServer_Register(t *testing.T) {
	type args struct {
		urlRoot string
	}
	tests := []struct {
		name   string
		apiKey string
		regReq interface{}
		want   int
	}{
		{
			name:   "Status OK",
			apiKey: APIKey,
			regReq: RegisterRequest{
				Username: "user1",
				Password: "password",
			},
			want: http.StatusOK,
		},
		{
			name:   "Correct API key but wrong input",
			apiKey: APIKey,
			regReq: AuthorizeRequest{
				Token: "token",
			},
			want: http.StatusBadRequest,
		},
		{
			name:   "wrongAPIKey",
			apiKey: "wasdfasfas",
			regReq: RegisterRequest{
				Username: "user1",
				Password: "password",
			},
			want: http.StatusUnauthorized,
		},
	}
	ts := createModelHandler()
	defer ts.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := genReq(tt.apiKey).SetBody(tt.regReq).Post(ts.URL + APIVersion + RegisterPath)
			assert.Equal(t, tt.want, res.StatusCode())
		})
	}
}
