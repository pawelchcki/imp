package main

import (
	"./registry"

	_ "./handlers"
	log "./nlog"
	"fmt"
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
