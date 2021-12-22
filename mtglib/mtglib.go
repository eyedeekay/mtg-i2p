package mtglib

import (
	"errors"
	"fmt"
	"net"

	mtglib2 "github.com/9seconds/mtg/v2/mtglib"
	"github.com/panjf2000/ants/v2"
)

type Proxy mtglib2.Proxy

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

		ipAddr := conn.RemoteAddr().(*net.TCPAddr).IP
		logger := p.logger.BindStr("ip", ipAddr.String())

		if p.whitelist != nil && !p.whitelist.Contains(ipAddr) {
			conn.Close()
			logger.Info("ip was rejected by whitelist")
			p.eventStream.Send(p.ctx, NewEventIPBlocklisted(ipAddr))

			continue
		}

		if p.blocklist.Contains(ipAddr) {
			conn.Close()
			logger.Info("ip was blacklisted")
			p.eventStream.Send(p.ctx, NewEventIPBlocklisted(ipAddr))

			continue
		}

		err = p.workerPool.Invoke(conn)

		switch {
		case err == nil:
		case errors.Is(err, ants.ErrPoolClosed):
			return nil
		case errors.Is(err, ants.ErrPoolOverload):
			logger.Info("connection was concurrency limited")
			p.eventStream.Send(p.ctx, NewEventConcurrencyLimited())
		}
	}
}
