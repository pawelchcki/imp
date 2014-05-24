package mw

import (
	"fmt"
	// "github.com/garyburd/redigo/redis"
	"../nlog"
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
		fmt.Printf("%+v", r.request.URL)
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
	fmt.Printf("%+v\n", x)
	return x.authorized
}
