package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yosa12978/lizardpoint/internal/config"
	"github.com/yosa12978/lizardpoint/internal/logging"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	conf := config.Get()
	logger := logging.NewJsonLogger(os.Stdout)

	server := newServer(ctx, conf.Server.Addr, logger)

	doneCh := make(chan struct{})
	go func() {
		logger.Info("server listening", "addr", conf.Server.Addr)
		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			logger.Error(err.Error())
		}
		close(doneCh)
	}()

	select {
	case <-doneCh:
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(timeout)
	}
	return nil
}
