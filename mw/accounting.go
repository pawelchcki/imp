package mw

import (
	// "github.com/garyburd/redigo/redis"
	"github.com/pchojnacki/intelligent_maybe_proxy/nlog"
	"net/http"
	// "net/url"
)

var commChannel = make(chan *initialCheckRequest, 1000)

func init() {
	nlog.Info("starting accounting")
	go waruj()
}

func waruj() {
	for r := range commChannel {
		nlog.Debugf("%+v", r.request.URL)
		r.responseChan <- &initialCheckResponse{true}
	}
}

type initialCheckResponse struct {
	authorized bool
}

type initialCheckRequest struct {
	request      *http.Request
	responseChan chan *initialCheckResponse
}

func VerifyToken(w http.ResponseWriter, r *http.Request) bool {
	checkRequest := &initialCheckRequest{r, make(chan *initialCheckResponse)}
	commChannel <- checkRequest
	x := <-checkRequest.responseChan
	nlog.Debugf("%+v\n", x)
	return x.authorized
}
