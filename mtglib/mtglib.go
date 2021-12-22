package mtglib

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/9seconds/mtg/v2/essentials"
	"github.com/9seconds/mtg/v2/mtglib"
	m "github.com/9seconds/mtg/v2/mtglib"
	"github.com/panjf2000/ants/v2"
)

type Proxy struct {
	m.Proxy
	streamWaitGroup sync.WaitGroup
	logger          m.Logger
	workerPool      *ants.PoolWithFunc
	ctx             context.Context
	eventStream     m.EventStream
}

type ProxyOpts m.ProxyOpts

func (p ProxyOpts) getConcurrency() int {
	if p.Concurrency == 0 {
		return m.DefaultConcurrency
	}

	return int(p.Concurrency)
}

func (p ProxyOpts) getLogger(name string) m.Logger {
	return p.Logger.Named(name)
}

// Serve starts a proxy on a given listener.
func (p *Proxy) Serve(listener net.Listener) error { // nolint: cyclop
	p.streamWaitGroup.Add(1)
	defer p.streamWaitGroup.Done()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-p.ctx.Done():
				return nil
			default:
				return fmt.Errorf("cannot accept a new connection: %w", err)
			}
		}

		ipAddr := conn.RemoteAddr()
		logger := p.logger.BindStr("ip", ipAddr.String())

		/*if p.whitelist != nil && !p.whitelist.Contains(ipAddr) {
			conn.Close()
			logger.Info("ip was rejected by whitelist")
			p.eventStream.Send(p.ctx, m.NewEventIPBlocklisted(ipAddr))

			continue
		}*/

		/*if p.blocklist.Contains(ipAddr) {
			conn.Close()
			logger.Info("ip was blacklisted")
			p.eventStream.Send(p.ctx, m.NewEventIPBlocklisted(ipAddr))

			continue
		}*/

		err = p.workerPool.Invoke(conn)

		switch {
		case err == nil:
		case errors.Is(err, ants.ErrPoolClosed):
			return nil
		case errors.Is(err, ants.ErrPoolOverload):
			logger.Info("connection was concurrency limited")
			p.eventStream.Send(p.ctx, m.NewEventConcurrencyLimited())
		}
	}
}

// NewProxy makes a new proxy instance.
func NewProxy(opts ProxyOpts) (*Proxy, error) {
	mp, err := m.NewProxy(mtglib.ProxyOpts(opts))
	if err != nil {
		return nil, err
	}
	p := &Proxy{
		Proxy:  *mp,
		logger: opts.getLogger("proxy"),
		//logger: opts.Logger,
		/*workerPool:      opts.WorkerPool,
		ctx:             opts.Ctx,*/
		eventStream:     opts.EventStream,
		streamWaitGroup: sync.WaitGroup{},
	}
	pool, err := ants.NewPoolWithFunc(opts.getConcurrency(),
		func(arg interface{}) {
			p.ServeConn(arg.(essentials.Conn))
		},
		ants.WithLogger(opts.getLogger("ants")),
		ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}

	p.workerPool = pool
	return p, nil
}
