package mwutils

import (
	"net/http"
)

type Connection struct {
	ResponseWriter     http.ResponseWriter
	Request            *http.Request
	ConnectionMetadata *interface{}
}
