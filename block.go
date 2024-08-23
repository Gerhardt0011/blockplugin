package blockplugin

import (
	"context"
	"fmt"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("blockplugin")

type Block struct {
	Next  plugin.Handler
	Zones []string
}

func (e Block) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	domain := state.Name()
	fmt.Printf("Request received: %s", domain)

	zone := plugin.Zones(e.Zones).Matches(state.Name())
	if zone == "" {
		fmt.Println("Zone not found on Server")
		return plugin.NextOrFailure(state.Name(), e.Next, ctx, w, r)
	}

	return plugin.NextOrFailure(e.Name(), e.Next, ctx, w, r)
}

func (e Block) Name() string { return "blockplugin" }

type ResponsePrinter struct {
	dns.ResponseWriter
}
