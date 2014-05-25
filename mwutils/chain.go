package mwutils

import (
	"net/http"
)

type Chain struct {
	mws        []func(con *Connection) bool
	handleFunc func(con *Connection)
}

func NewChain() *Chain {
	return new(Chain)
}

func (c *Chain) ThenSimple(mw func(w http.ResponseWriter, r *http.Request) bool) *Chain {
	return c.Then(func(con *Connection) bool {
		return mw(con.ResponseWriter, con.Request)
	})
}

func (c *Chain) Then(mw func(con *Connection) bool) *Chain {
	c.mws = append(c.mws, mw)
	return c
}

func (c *Chain) HandleFuncSimple(f func(w http.ResponseWriter, r *http.Request)) *Chain {
	return c.HandleFunc(func(con *Connection) {
		f(con.ResponseWriter, con.Request)
	})
}

func (c *Chain) HandleFunc(f func(con *Connection)) *Chain {
	c.handleFunc = f
	return c
}

func (c *Chain) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	brokenChain := false
	connection := &Connection{w, r, ConnectionMetadata{}}
	for _, mw := range c.mws {
		if mw(connection) == false {
			brokenChain = true
			break
		}
	}
	if !brokenChain && c.handleFunc != nil {
		c.handleFunc(connection)
	} else if c.handleFunc == nil {
		http.NotFound(w, r)
	}
}
