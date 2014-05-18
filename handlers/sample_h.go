package handlers

import (
	"../registry"
	// "net/http"
	"net/http/httputil"
	"net/url"
)

func init() {
	reg := registry.Group("sample_h")
	u, _ := url.Parse("https://google.pl")

	reg.HandleFunc("/api/", httputil.NewSingleHostReverseProxy(u).ServeHTTP)
}

// func sampleApi(w http.ResponseWriter, r *http.Request) {

// 	httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
// }
