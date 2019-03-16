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
	StandAlone bool
	APIKey     string
}

// NewServiceHandler creates the http handler for auth service
// There are two modes: standalone or integrated modes.
// Standalone: Run mauth as a standalone HTTP microservice
// Integrated: Integrate mauth into existing services
func NewServiceHandler(apiKey string, standAlone bool) *ServiceHandler {
	return &ServiceHandler{
		StandAlone: standAlone,
		APIKey:     apiKey,
	}
}

func (p ServiceHandler) register(req *restful.Request, resp *restful.Response) {
	registerRequest := &RegisterRequest{}
	// err := json.NewDecoder(req.Request.Body).Decode(&registerRequest)
	err := req.ReadEntity(registerRequest)
	if err != nil {
		glog.Warningf("Get register request error: %v", err)
		resp.WriteErrorString(http.StatusBadRequest, "Wrong JSON input for reigster.")
		return
	}
	if registerRequest.Username == "" || registerRequest.Password == "" {
		glog.Warningf("Empty username or password")
		resp.WriteErrorString(http.StatusBadRequest, "Invalid username or password")
		return
	}
	glog.Infof("here in register")
	resp.WriteHeader(http.StatusOK)
}

func (p ServiceHandler) loginPost(req *restful.Request, resp *restful.Response) {
	loginRequest := &LoginRequest{}
	err := req.ReadEntity(loginRequest)
	if err != nil {
		glog.Warningf("Login post request error: %v", err)
		resp.WriteErrorString(http.StatusBadRequest, "Wrong JSON input for login.")
		return
	}
	glog.Infof("here in login post")
	resp.WriteHeader(http.StatusOK)
}

func (p ServiceHandler) loginGet(req *restful.Request, resp *restful.Response) {
	glog.Infof("here in login get")
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
	// When running in standalone mode, we use api key to ensure
	// requester is one of us
	if p.StandAlone {
		wsContainer.Filter(p.apiKeyFilter)
	}

	ws := new(restful.WebService)
	ws.Consumes("*/*"). // Set media acceptance to wildcast
				Produces(restful.MIME_JSON)
	ws.Route(ws.POST(urlRoot + APIVersion + RegisterPath).To(p.register))
	// Different method in different mode:
	// standalone: the username and password usually comes as POST JSON
	// integrated: the user name and password comes in the header
	if p.StandAlone {
		ws.Route(ws.POST(urlRoot + APIVersion + LoginPath).To(p.loginPost))
	} else {
		ws.Route(ws.GET(urlRoot + APIVersion + LoginPath).To(p.loginGet))
	}
	ws.Route(ws.POST(urlRoot + APIVersion + AuthPath).To(p.auth))
	ws.Route(ws.DELETE(urlRoot + APIVersion + LogoutPath).To(p.logout))
	//wsContainer.Add(apiv1)
	wsContainer.Add(ws)
	return wsContainer
}
