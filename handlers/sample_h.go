package handlers

import (
	"../registry"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func init() {
	reg := registry.Group("sample_h")
	reg.HandleFunc("/api/", sampleApi)
}

func sampleApi(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse("https://google.pl")
	httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
}
