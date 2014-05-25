package mw

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"../nlog"
)

const DefaultWikiaName = "www"
const WikiaNameQueryParam = "wikianame"
const WikiaLangQueryParam = "wikialang"

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func defaultWikiBaseUrl(wikiaName, wikiaLang string) *url.URL {
	u := new(url.URL)
	baseHost := "wikia.com"
	u.Scheme = "http"
	if wikiaLang != "en" && wikiaLang != "" {
		u.Host = wikiaLang + "." + wikiaName + "." + baseHost
	} else {
		u.Host = wikiaName + "." + baseHost
	}
	return u
}

func WikiProxy(wikiBaseUrl func(wikiaName, wikiaLang string) *url.URL) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{Director: defaultWikiProxyDirector(wikiBaseUrl)}
}

func defaultWikiProxyDirector(wikiBaseUrl func(wikiaName, wikiaLang string) *url.URL) func(req *http.Request) {
	return func(req *http.Request) {
		query := req.URL.Query()
		wikiaName := query.Get(WikiaNameQueryParam)
		if wikiaName != "" {
			query.Del(WikiaNameQueryParam)
		} else {
			wikiaName = DefaultWikiaName
		}
		wikiaLang := query.Get(WikiaLangQueryParam)
		if wikiaLang != "" {
			query.Del(WikiaLangQueryParam)
		}
		req.URL.RawQuery = query.Encode()

		target := wikiBaseUrl(wikiaName, wikiaLang)

		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		nlog.Debugf("default wiki request: %+v", req)

		if target.RawQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
		}
	}
}

func DefaultWikiProxy() func(rw http.ResponseWriter, req *http.Request) {
	return WikiProxy(defaultWikiBaseUrl).ServeHTTP
}
