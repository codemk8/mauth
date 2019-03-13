package auth

import (
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

// ServiceHandler implements the auth restful APIs
type ServiceHandler struct {
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

// Register sets up the
func (p ServiceHandler) Register(urlRoot string) {
	ws := new(restful.WebService)
	ws.Route(ws.POST(urlRoot + "/register").To(p.register))
	ws.Route(ws.GET(urlRoot + "/login").To(p.login))
	ws.Route(ws.POST(urlRoot + "/auth").To(p.auth))
	ws.Route(ws.POST(urlRoot + "/logout").To(p.logout))
	restful.Add(ws)
}
