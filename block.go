package blockplugin

import (
	"context"
	"fmt"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"
	clog "github.com/coredns/coredns/plugin/pkg/log"

	"github.com/miekg/dns"
)

var log = clog.NewWithPlugin("example")

type Block struct {
	Next plugin.Handler
}

func (e Block) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Debug("Received response")

	pw := NewResponsePrinter(w)

	requestCount.WithLabelValues(metrics.WithServer(ctx)).Inc()

	return plugin.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
}

func (e Block) Name() string { return "blockplugin" }

type ResponsePrinter struct {
	dns.ResponseWriter
}

func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}

func (r *ResponsePrinter) WriteMsg(res *dns.Msg) error {
	fmt.Println("======================================")
	fmt.Printf("%+v\n", res)
	fmt.Println("======================================")

	return r.ResponseWriter.WriteMsg(res)
}
