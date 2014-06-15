package main

import (
	"github.com/pchojnacki/intelligent_maybe_proxy/registry"

	"fmt"
	_ "github.com/pchojnacki/intelligent_maybe_proxy/handlers"
	log "github.com/pchojnacki/intelligent_maybe_proxy/nlog"
	"net/http"
	// "net/http/httputil"
	// "net/url"
)

// type Xx{
// 	string Prefix
// 	map[string]func()http.Handler Handler
// }

func main() {
	//	url, _ := url.Parse("https://www.google.com")
	ip := "0.0.0.0"
	port := 8080
	ip_port := fmt.Sprintf("%s:%d", ip, port)
	log.Infof("Listening on %s", ip_port)
	http.ListenAndServe(ip_port, registry.GetMainMux())
}

/*
request comes in
is
*/
//resp, err := g
