package bootstrap

import (
	"context"
	"sync"

	"github.com/leojin/go-service-bootstrap/library"
	"github.com/leojin/go-service-bootstrap/utils"
)

type Server interface {
	Start(ctx context.Context) error
	ServerName() string
	EnableDebug()
}

type Servers struct {
	cancel context.CancelFunc
	wg     sync.WaitGroup

	list []Server
}

func (s *Servers) Add(newServer Server) {
	s.list = append(s.list, newServer)
}

func (s *Servers) Start(ctx context.Context) {
	cancelCtx, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	s.wg = sync.WaitGroup{}

	for _, item := range s.list {
		utils.Out.Printf("[servers] [server: %s] start", item.ServerName())
		s.wg.Add(1)
		go func(newServer Server) {
			if library.ConfigApp.Debug {
				newServer.EnableDebug()
			}
			err := newServer.Start(cancelCtx)
			utils.Out.Printf("[servers] [server: %s] [err: %+v] stop", item.ServerName(), err)
			s.wg.Done()
			cancel()
		}(item)
	}

	for {
		select {
		case <-cancelCtx.Done():
			utils.Out.Printf("[servers] context done: %v", cancelCtx.Err())
			return
		}
	}
}

func (s *Servers) Stop() {
	if s.cancel == nil {
		return
	}
	s.cancel()
	s.cancel = nil
	utils.Out.Printf("[servers] stop wait")
	s.wg.Wait()
	utils.Out.Printf("[servers] stop done")
}
