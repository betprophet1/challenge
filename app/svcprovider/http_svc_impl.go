package svcprovider

import (
	"context"
	"fmt"
	"net/http"
)

type httpsvc struct {
	http *http.Server
}

func (g httpsvc) Name() string {
	return "http"
}

func (g httpsvc) Start() error {
	if err := g.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http: start server: %w", err)
	}
	return nil
}

func (g httpsvc) Stop(ctx context.Context) error {
	return g.http.Shutdown(ctx)
}

func (g httpsvc) BeforeStop() {}

func NewHttpService(h *http.Server) ServiceProvider {
	return &httpsvc{http: h}
}
