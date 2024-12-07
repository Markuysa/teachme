package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"session/internal/config"
	clientRepos "session/internal/repository/client"
	clientService "session/internal/service/client"
	"session/internal/transport/client"
	closerPkg "session/pkg/closer"

	"github.com/Markuysa/pkg/logger"
	"github.com/Markuysa/pkg/redis"
	"github.com/Markuysa/pkg/srv/grpc"
)

func Run(_ context.Context, cfg *config.Config) error {
	closer := closerPkg.New()

	err := logger.InitLogger(cfg.Logger)
	if err != nil {
		return err
	}

	rdConn, err := redis.New(cfg.Redis)
	if err != nil {
		return err
	}
	closer.AddErrCloser(rdConn.Close)

	sessionsRepository := clientRepos.New(rdConn)
	sessionService := clientService.New(sessionsRepository)

	grpcTransport := client.New(sessionService)

	grpc, err := grpc.NewServer(
		grpc.WithConfig(&cfg.GRPC),
		grpc.WithRegistes(
			grpcTransport,
		),
	)
	if err != nil {
		return err
	}
	closer.AddCloser(grpc.GracefulStop)

	// TODO add serve of grpc.

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-quitCh

	return closer.Close()
}
