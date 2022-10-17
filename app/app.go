package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func RunSession(session *Session) {
	var quitChan = make(chan os.Signal, 1)
	session.Out.Info("services are starting...\n")
	svcName := session.Service.Name()
	go func() {
		session.Out.Info("service `%v` has started\n", svcName)
		if err := session.Service.Start(); err != nil {
			session.Out.Fatal("service `%v` run failed\n", svcName)
		}
	}()
	signal.Notify(quitChan, os.Interrupt, syscall.SIGTERM)
	<-quitChan
	session.Service.BeforeStop()
	session.Out.Info("servies are shutting down...\n")
	ctx, cancel := context.WithTimeout(context.Background(), session.Options.GracefulTimeout)
	defer cancel()
	if err := session.End(ctx); err != nil {
		session.Out.Error("service `%v` shutdown with failure | err=%s\n", svcName, err.Error())
	}
}
