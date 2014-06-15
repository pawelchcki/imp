package handlers

import (
	"github.com/pchojnacki/intelligent_maybe_proxy/mw"
	"github.com/pchojnacki/intelligent_maybe_proxy/mwutils"
	"github.com/pchojnacki/intelligent_maybe_proxy/registry"
)

func init() {
	reg := registry.Group("sample_h")
	// u, _ := url.Parse("https://google.pl")
	c := reg.HandleNewChain("/api/")
	c.ThenSimple(mw.VerifyToken)
	c.HandleFunc(mwutils.ConnectionMapperFuncWrapper(mw.DefaultWikiProxy()))

}
