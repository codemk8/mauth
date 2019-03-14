package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/codemk8/mauth/pkg/auth"
	"github.com/golang/glog"
)

var (
	port   = flag.Int("port", 55555, "The webhook server port")
	apiKey = flag.String("apikey", "", "The API key for authentication")
)

func main() {
	flag.Parse()
	apiKey := os.Getenv("MAUTH_API_KEY")
	if apiKey == "" {
		glog.Fatal("MAUTH_API_KEY not found.")
	}
	sh := auth.NewServiceHandler(apiKey)
	httpHandler := sh.Register("")
	http.Handle("/", httpHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
