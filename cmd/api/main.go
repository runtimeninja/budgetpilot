package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/runtimeninja/budgetpilot/internal/config"
	"github.com/runtimeninja/budgetpilot/internal/observability"
	"github.com/runtimeninja/budgetpilot/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	logger := observability.NewLogger(cfg.Env)

	h := router.New(router.Deps{
		Env:    cfg.Env,
		Logger: logger,
	})

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("server_start", "addr", cfg.HTTPAddr, "env", cfg.Env)
		errCh <- srv.ListenAndServe()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		logger.Info("shutdown_signal_received")
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			logger.Error("server_failed", "error", err)
			os.Exit(1)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown_failed", "error", err)
	} else {
		logger.Info("shutdown_complete")
	}
}
