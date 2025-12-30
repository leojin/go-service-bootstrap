package bootstrap

import (
	"context"
	"github.com/leojin/go-service-bootstrap/library"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/leojin/go-service-bootstrap/utils"
)

type ServiceBootstrap struct {
	ctx    context.Context
	cancel context.CancelFunc

	basePath string

	servers Servers
	jobs    Jobs

	callbackStop []func()
}

func (b *ServiceBootstrap) Init(basePath string) (err error) {
	b.basePath = basePath
	b.ctx, b.cancel = utils.GetSignalContext()

	if _, err = toml.Decode(path.Join(basePath, "config/app.toml"), &library.ConfigApp); err != nil {
		return err
	}

	return nil
}

func (b *ServiceBootstrap) Start() {

	defer func() {
		b.cancel()
		for _, fn := range b.callbackStop {
			fn()
		}
	}()

	// Start Jobs
	//b.RegisterCallbackStop(b.jobs.Stop)
	//b.jobs.Start(b.ctx)

	// Start Servers
	b.RegisterCallbackStop(b.servers.Stop)
	b.servers.Start(b.ctx)

}

//func (b *ServiceBootstrap) RegisterJob(newJob job) {
//	b.jobs.Add(newJob)
//}

func (b *ServiceBootstrap) RegisterServer(newServer Server) {
	b.servers.Add(newServer)
}

func (b *ServiceBootstrap) RegisterLog(configName string, logger Lo) {

}

func (b *ServiceBootstrap) RegisterCallbackStop(newCallback func()) {
	b.callbackStop = append(b.callbackStop, newCallback)
}
