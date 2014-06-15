package mwutils

import (
	"github.com/pchojnacki/intelligent_maybe_proxy/nlog"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// taken and simplified from gorilla/context
// to be used in 3rd party handler to access Connection Object

var (
	mutex   sync.RWMutex
	data    = make(map[*http.Request]*Connection)
	dataUrl = make(map[*url.URL]*Connection)
	datat   = make(map[*http.Request]int64)
)

func MapperSet(val *Connection) {
	nlog.Debugf("Setting %p", val.Request)
	r := val.Request
	mutex.Lock()
	data[r] = val
	dataUrl[r.URL] = val
	datat[r] = time.Now().Unix()
	mutex.Unlock()
}

// GetOk returns stored value and presence state like multi-value return of map access.
func MapperGetOk(r *http.Request) (*Connection, bool) {
	for k, d := range data {
		nlog.Debugf("Getting fromXXXX KEY: %p, DATA:%+v", k, d)

	}
	nlog.Debugf("Getting for %p", r)

	mutex.RLock()
	value, ok := data[r]
	// if request lookup fail, lookup using url pointer. This is workaround to use it in ReverseProxy director, because it copied Request object
	if !ok && r.URL != nil {
		value, ok = dataUrl[r.URL]
	}
	mutex.RUnlock()
	return value, ok
}

// Clear removes all values stored for a given request.
//
// This is usually called by a handler wrapper to clean up request
// variables at the end of a request lifetime. See ClearHandler().
func MapperClear(r *http.Request) {
	mutex.Lock()
	clear(r)
	mutex.Unlock()
}

// clear is Clear without the lock.
func clear(r *http.Request) {
	delete(data, r)
	delete(datat, r)
	delete(dataUrl, r.URL)
}

// Purge removes request data stored for longer than maxAge, in seconds.
// It returns the amount of requests removed.
//
// If maxAge <= 0, all request data is removed.
//
// This is only used for sanity check: in case context cleaning was not
// properly set some request data can be kept forever, consuming an increasing
// amount of memory. In case this is detected, Purge() must be called
// periodically until the problem is fixed.
func MapperPurge(maxAge int) int {
	mutex.Lock()
	count := 0
	if maxAge <= 0 {
		count = len(data)
		data = make(map[*http.Request]*Connection)
		dataUrl = make(map[*url.URL]*Connection)
		datat = make(map[*http.Request]int64)
	} else {
		min := time.Now().Unix() - int64(maxAge)
		for r := range data {
			if datat[r] < min {
				clear(r)
				count++
			}
		}
	}
	mutex.Unlock()
	return count
}

// ClearHandler wraps an http.Handler and clears request values at the end
// of a request lifetime.
func ConnectionMapperFuncWrapper(f func(w http.ResponseWriter, r *http.Request)) func(con *Connection) {
	return func(con *Connection) {
		MapperSet(con)
		f(con.ResponseWriter, con.Request)
		defer MapperClear(con.Request)
	}
}

func ConnectionMapperHandlerWrapper(h http.Handler) func(con *Connection) {
	return func(con *Connection) {
		MapperSet(con)
		defer MapperClear(con.Request)
		h.ServeHTTP(con.ResponseWriter, con.Request)
	}
}
