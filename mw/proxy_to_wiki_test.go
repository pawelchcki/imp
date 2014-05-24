package mw

import (
	// "io/ioutil"
	// // "fmt"

	"net/http"
	// "net/http/httptest"
	"net/url"
	"testing"
)

func TestSingleJoiningSlash(t *testing.T) {
	if singleJoiningSlash("a", "b") != "a/b" {
		t.Fail()
	}
	if singleJoiningSlash("a/", "b") != "a/b" {
		t.Fail()
	}
}

func TestSampleWikiUrl(t *testing.T) {
	if defaultWikiBaseUrl("muppet", "en").String() != "http://muppet.wikia.com" {
		t.Fail()
	}
	if defaultWikiBaseUrl("muppet", "pl").String() != "http://pl.muppet.wikia.com" {
		t.Fail()
	}
}

func TestWikiProxyDirector(t *testing.T) {
	proxyDirector := defaultWikiProxyDirector(defaultWikiBaseUrl)
	sampleUrl, _ := url.Parse("http://this.api.call.endpoint/api/v1/based/query?wikianame=muppet&wikialang=pl")
	req := http.Request{Method: "GET", URL: sampleUrl}
	proxyDirector(&req)
	if req.URL.Host != "pl.muppet.wikia.com" {
		t.Fatal(req.URL.Host)
	}

	if req.URL.Path != "/api/v1/based/query" {
		t.Fatal(req.URL.Path)
	}
	if req.URL.RawQuery != "" {
		t.Fatal(req.URL.RawQuery)
	}
}
