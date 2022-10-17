package svcprovider

import "context"

type ServiceProvider interface {
	Name() string
	BeforeStop()
	Start() error
	Stop(ctx context.Context) error
}
