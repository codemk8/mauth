package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/codemk8/mauth/pkg/auth"
)

var (
	port   = flag.Int("port", 55555, "The webhook server port")
	apiKey = flag.String("apikey", "", "The API key for authentication")
)

func main() {
	flag.Parse()
	sh := auth.NewServiceHandler(*apiKey)
	httpHandler := sh.Register("")
	http.Handle("/", httpHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
