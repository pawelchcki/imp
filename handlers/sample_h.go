package handlers

import (
	"../mw"
	"../mwutils"
	"../registry"

	// "net/http/httputil"
	// "net/url"
)

func init() {
	reg := registry.Group("sample_h")
	// u, _ := url.Parse("https://google.pl")

	c := mwutils.NewChain()
	c.Then(mw.VerifyToken)
	c.HandleFunc(mw.DefaultWikiProxy())
	reg.HandleFunc("/api/", c.ServeHTTP)

	// httputil.NewSingleHostReverseProxy(u).ServeHTTP(/w, r)

}

// func sampleApi(w http.ResponseWriter, r *http.Request) {

// 	httputil.NewSingleHostReverseProxy(u).ServeHTTP(/w, r)
// }
