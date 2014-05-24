package handlers

import (
	"../mwchain"
	"../registry"

	// "net/http/httputil"
	// "net/url"
)

func init() {
	reg := registry.Group("sample_h")
	// u, _ := url.Parse("https://google.pl")

	c := mwchain.NewChain()
	c.Then(registry.VerifyToken)
	c.HandleFunc(registry.DefaultWikiProxy().ServeHTTP)
	reg.HandleFunc("/api/", c.ServeHTTP)

	// httputil.NewSingleHostReverseProxy(u).ServeHTTP(/w, r)

}

// func sampleApi(w http.ResponseWriter, r *http.Request) {

// 	httputil.NewSingleHostReverseProxy(u).ServeHTTP(/w, r)
// }
