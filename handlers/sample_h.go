package handlers

import (
	"../mw"
	"../registry"
)

func init() {
	reg := registry.Group("sample_h")
	// u, _ := url.Parse("https://google.pl")
	c := reg.HandleNewChain("/api")
	c.Then(mw.VerifyToken)
	c.HandleFunc(mw.DefaultWikiProxy())
}

// func sampleApi(w http.ResponseWriter, r *http.Request) {

// 	httputil.NewSingleHostReverseProxy(u).ServeHTTP(/w, r)
// }
