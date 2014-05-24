package mwutils

import (
	"net/http"
)

type Chain struct {
	mws        []func(rw http.ResponseWriter, req *http.Request) bool
	handleFunc func(rw http.ResponseWriter, req *http.Request)
}

func NewChain() *Chain {
	return new(Chain)
}

func (c *Chain) Then(mw func(w http.ResponseWriter, r *http.Request) bool) *Chain {
	c.mws = append(c.mws, mw)
	return c
}

func (c *Chain) HandleFunc(f func(w http.ResponseWriter, r *http.Request)) *Chain {
	c.handleFunc = f
	return c
}

func (c *Chain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	brokenChain := false
	for _, mw := range c.mws {
		if mw(w, r) == false {
			brokenChain = true
			break
		}
	}
	if !brokenChain && c.handleFunc != nil {
		c.handleFunc(w, r)
	} else if c.handleFunc == nil {
		http.NotFound(w, r)
	}
}
