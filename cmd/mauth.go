package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/codemk8/mauth/pkg/auth"
)

var (
	port = flag.Int("port", 55555, "The webhook server port")
)

func main() {
	flag.Parse()
	sh := auth.ServiceHandler{}
	httpHandler := sh.Register("")
	http.Handle("/", httpHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
