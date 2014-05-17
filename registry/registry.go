package registry

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	fmt.Println("")
}

type ExtendedInfo struct {
	OriginalRequest *http.Request
}

type Info struct {
	mux *http.ServeMux
}

type mainHandler struct{}

var serveMux *http.ServeMux = http.NewServeMux()

func (mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h, pattern := serveMux.Handler(r)
	// modify r url and cut out the matched pattern
	// fmt.Printf("aaaa %+v\n", pattern)

	trimPrefixFromURL(r.URL, pattern)
	// fmt.Printf("aaaa %+v\n", r.URL)

	h.ServeHTTP(w, r)
}

func trimPrefixFromURL(u *url.URL, pattern string) {
	u.Path = strings.TrimPrefix(u.Path, pattern)
	if len(u.Path) == 0 || u.Path[0] != '/' {
		u.Path = "/" + u.Path
	}
}

func GetMainMux() http.Handler {
	return mainHandler{}
}

func SignUp(name string) *Info {
	ret := new(Info)
	ret.mux = serveMux
	return ret
}

func (i *Info) Handle(path string, f func(w http.ResponseWriter, r *http.Request, e *ExtendedInfo)) {

}

func (i *Info) HandleFunc(path string, f func(w http.ResponseWriter, r *http.Request)) {
	i.mux.HandleFunc(path, f)
}
