package blockplugin

import (
	"context"
	"fmt"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"

	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("blockplugin")

type Block struct {
	Next plugin.Handler
}

func (e Block) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", w)

	return plugin.NextOrFailure(e.Name(), e.Next, ctx, w, r)
}

func (e Block) Name() string { return "blockplugin" }

type ResponsePrinter struct {
	dns.ResponseWriter
}
