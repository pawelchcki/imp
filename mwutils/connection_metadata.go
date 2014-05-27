package mwutils

import (
	"net/url"
)

type Wikia struct {
	Name string
	Lang string
}

type ConnectionMetadata struct {
	TargetWikiaUrl *url.URL
	Wikia          *Wikia
}
