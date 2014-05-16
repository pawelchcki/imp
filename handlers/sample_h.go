package handlers

import (
	"../registry"
	"net/http"
)

func init() {
	reg := registry.SignUp("sample_h")
	reg.Handle("/api", sampleApi)
}

func sampleApi(w http.ResponseWriter, r *http.Request, e *registry.ExtendedInfo) {

}
