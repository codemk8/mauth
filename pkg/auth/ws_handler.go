package auth

import (
	"net/http"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

const (
	// APIVersion is the prefix root of the serving uri (version)
	APIVersion   = "/v1"
	RegisterPath = "/register"
	LoginPath    = "/login"
	LogoutPath   = "/logout"
	AuthPath     = "/auth"
)

// ServiceHandler implements the auth restful APIs
type ServiceHandler struct {
	APIKey string
}

func NewServiceHandler(apiKey string) *ServiceHandler {
	return &ServiceHandler{
		APIKey: apiKey,
	}
}

func (p ServiceHandler) register(req *restful.Request, resp *restful.Response) {
	glog.Infof("here in register")
}

func (p ServiceHandler) login(req *restful.Request, resp *restful.Response) {
	glog.Infof("here in login")
}

func (p ServiceHandler) auth(req *restful.Request, resp *restful.Response) {
	glog.Infof("here in auth")
}

func (p ServiceHandler) logout(req *restful.Request, resp *restful.Response) {
	glog.Infof("here in logout")
}

func (p ServiceHandler) apiKeyFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	keyName := "Authorization"
	glog.Infof("auth %s", req.Request.Header.Get(keyName))
	bearerAndToken := strings.Split(req.Request.Header.Get(keyName), " ")
	if len(bearerAndToken) == 2 {
		if bearerAndToken[1] == p.APIKey {
			chain.ProcessFilter(req, resp)
			return
		}
	}
	glog.Warning("Received a request with an invalid API key.")
	resp.WriteHeader(http.StatusUnauthorized)
}

// Register sets up the
func (p ServiceHandler) Register(urlRoot string) http.Handler {
	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)
	wsContainer.Filter(p.apiKeyFilter)
	ws := new(restful.WebService)
	ws.Consumes("*/*"). // Set media acceptance to wildcast
				Produces(restful.MIME_JSON)
	ws.Route(ws.POST(urlRoot + APIVersion + RegisterPath).To(p.register))
	ws.Route(ws.GET(urlRoot + APIVersion + LoginPath).To(p.login))
	ws.Route(ws.POST(urlRoot + APIVersion + AuthPath).To(p.auth))
	ws.Route(ws.POST(urlRoot + APIVersion + LogoutPath).To(p.logout))
	//wsContainer.Add(apiv1)
	wsContainer.Add(ws)
	return wsContainer
}
