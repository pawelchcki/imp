package registry

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestUrlTrimming(t *testing.T) {

	td := [][]string{
		[]string{"http://www.onet.pl/some/url", "/some", "/url"},
		[]string{"http://onet.pl/some/url", "url", "/some/url"},
		[]string{"http://onet.pl/some/url", "some/url", "/some/url"},
		[]string{"http://onet.pl/some/url", "/some/url", "/"},
	}

	for _, params := range td {
		u, _ := url.Parse(params[0])
		if trimPrefixFromURL(u, params[1]); u.Path != params[2] {
			t.Errorf("%s != %s", u.Path, params[2])
		}
	}
}

func BenchmarkUrlTrimming(b *testing.B) {
	u, _ := url.Parse("http://www.onet.pl/some/url")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var t url.URL
		t = *u
		trimPrefixFromURL(&t, "/some/")
	}
}

func TestHandlerRun(t *testing.T) {
	// x := GetMainMux()
	i := Group("test")

	didRun := false

	i.HandleFunc("/some/", func(w http.ResponseWriter, r *http.Request) {
		didRun = true
		if r.URL.Path != "/rest/" {
			t.Fatalf("%+v != /rest/", r.URL.Path)
		}
	})

	resp := httptest.NewRecorder()
	req := new(http.Request)
	req.URL, _ = url.Parse("/some/rest/")

	GetMainMux().ServeHTTP(resp, req)

	if !didRun {
		t.Fatal("handler wasn't run")
	}
}
func TestFailToAddSameRouteTwice(t *testing.T) {
	i := Group("test")
	i.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {})

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	i.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {})
}
