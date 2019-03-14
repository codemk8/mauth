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

func createModelHandler() *httptest.Server {
	sh := ServiceHandler{}
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

func genReq() *resty.Request {
	return resty.R().SetHeader("Content-Type", restful.MIME_JSON) //.SetAuthToken(testUser)
}

func TestAPI(t *testing.T) {
	// A simple create model test case
	ts := createModelHandler()
	defer ts.Close()
	res, err := genReq().Post(ts.URL + APIVersion + RegisterPath)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, res.StatusCode())
}
