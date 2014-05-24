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
	c.Then(func(w http.ResponseWriter, r *http.Request) bool {
		aRun = true
		return true
	}).Then(func(w http.ResponseWriter, r *http.Request) bool {
		bRun = true
		return true
	})
	c.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
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
	c.Then(func(w http.ResponseWriter, r *http.Request) bool {
		aRun = true
		return true
	}).Then(func(w http.ResponseWriter, r *http.Request) bool {
		bRun = true
		return false
	})
	c.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		cRun = true
	})

	c.ServeHTTP(httptest.NewRecorder(), new(http.Request))
	if !(aRun && bRun && !cRun) {
		t.Fail()
	}
}
