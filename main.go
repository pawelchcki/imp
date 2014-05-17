package main

import (
	"./registry"

	_ "./handlers"
	//"fmt"
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

	http.ListenAndServe("0.0.0.0:8080", registry.GetMainMux())
}

/*
request comes in
is
*/
//resp, err := g
