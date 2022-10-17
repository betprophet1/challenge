package app

import (
	"context"

	"project/app/svcprovider"
	"project/common/log"
)

type Session struct {
	Options Options
	Service svcprovider.ServiceProvider
	Out     *log.Logger
}

func NewSession(svc svcprovider.ServiceProvider) *Session {
	return &Session{
		Service: svc,
		Options: parseOpt(),
		Out:     &log.Logger{},
	}
}

func (s Session) End(ctx context.Context) error {
	return s.Service.Stop(ctx)
}
