package mw

import (
	"github.com/pchojnacki/intelligent_maybe_proxy/mwutils"
	"net/http"
	// "net/http/httputil"

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

func TestWikiaDesignationQueryParser(t *testing.T) {
	con := getFakeCon()
	WikiaDesignationQueryParser(con)

	if con.Metadata.Wikia.Name != "muppet" {
		t.Fail()
	}
	if con.Metadata.Wikia.Lang != "pl" {
		t.Fail()
	}
}

func TestDefaultTargetWikiaURL(t *testing.T) {
	con := getFakeCon()
	WikiaDesignationQueryParser(con)
	con.Metadata.Wikia.Lang = ""
	DefaultTargetWikiaURL(con)
	if con.Metadata.TargetWikiaUrl.String() != "http://muppet.wikia.com" {
		t.Fatalf("%s != %s", con.Metadata.TargetWikiaUrl.String(), "http://muppet.wikia.com")
	}
}

func getFakeCon() *mwutils.Connection {
	con := &mwutils.Connection{}
	sampleUrl, _ := url.Parse("http://this.api.call.endpoint/api/v1/based/query?wikianame=muppet&wikialang=pl")
	con.Request = &http.Request{Method: "GET", URL: sampleUrl}
	return con
}

func TestWikiProxyDirector(t *testing.T) {
	con := getFakeCon()

	WikiaDesignationQueryParser(con)
	DefaultTargetWikiaURL(con)

	mwutils.MapperSet(con)
	proxyDirector := defaultWikiProxyDirector()

	proxyDirector(con.Request)
	if con.Request.URL.Host != "pl.muppet.wikia.com" {
		t.Fatal(con.Request.URL.Host)
	}

	if con.Request.URL.Path != "/api/v1/based/query" {
		t.Fatal(con.Request.URL.Path)
	}
	if con.Request.URL.RawQuery != "" {
		t.Fatal(con.Request.URL.RawQuery)
	}
}
