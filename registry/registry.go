package registry

import (
	log "../nlog"

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
	mux       *http.ServeMux
	groupName string
}

type mainHandler struct{}

var serveMux *http.ServeMux = http.NewServeMux()
var groupMap map[string]*Info = make(map[string]*Info)

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
	trimPrefixFromURL(r.URL, pattern)

	log.Debug("info")
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

func Group(name string) *Info {
	ret := new(Info)
	ret.mux = serveMux
	ret.groupName = name
	return ret
}

func (i *Info) HandleFunc(path string, f func(w http.ResponseWriter, r *http.Request)) {
	if _, ok := groupMap[path]; ok == true {
		panic("path duplicated: " + path)
	}
	groupMap[path] = i
	i.mux.HandleFunc(path, f)
}
