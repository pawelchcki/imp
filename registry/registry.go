package registry

import (
	"net/http"
)

type ExtendedInfo struct {
}

type Info struct {
}

func GetMainHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func SignUp(name string) *Info {
	ret := new(Info)
	return ret
}

func (*Info) Handle(path string, f func(w http.ResponseWriter, r *http.Request, e *ExtendedInfo)) {

}
