package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/postgres"
	"gitlab.com/coinhubs/balance/internal/config"
	closerPkg "gitlab.com/coinhubs/balance/pkg/closer"
)

func Run(_ context.Context, cfg *config.Config) error {
	closer := closerPkg.New()

	err := logger.InitLogger(cfg.Logger)
	if err != nil {
		return err
	}

	pgConn, err := postgres.New(cfg.Postgres)
	if err != nil {
		return err
	}
	closer.AddCloser(pgConn.Close)

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quitCh

	closer.Close()

	return nil
}
