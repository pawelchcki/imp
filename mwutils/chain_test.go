package mwutils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChainAllRun(t *testing.T) {
	c := NewChain()
	aRun := false
	bRun := false
	cRun := false
	c.Then(func(con *Connection) bool {
		aRun = true
		return true
	}).Then(func(con *Connection) bool {
		bRun = true
		return true
	})
	c.HandleFunc(func(con *Connection) {
		cRun = true
	})

	c.ServeHTTP(httptest.NewRecorder(), new(http.Request))
	if !(aRun && bRun && cRun) {
		t.Fail()
	}
}

func TestChainSomeCausedChainBreak(t *testing.T) {
	c := NewChain()
	aRun := false
	bRun := false
	cRun := false
	c.Then(func(con *Connection) bool {
		aRun = true
		return true
	}).Then(func(con *Connection) bool {
		bRun = true
		return false
	})
	c.HandleFunc(func(con *Connection) {
		cRun = true
	})

	c.ServeHTTP(httptest.NewRecorder(), new(http.Request))
	if !(aRun && bRun && !cRun) {
		t.Fail()
	}
}
