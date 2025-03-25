package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luyanakat/golang-base-project/app"
	"github.com/luyanakat/golang-base-project/internal/helpers"
	"github.com/luyanakat/golang-base-project/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()
	sugarLog := logger.Sugar()

	err := helpers.LoadEnv()
	if err != nil {
		sugarLog.Fatalf("Failed to load env: %v", err)
	}

	app := app.InitApp(sugarLog)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: app.R,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugarLog.Fatalf("Server failed: %v", err)
		}
	}()
	sugarLog.Info("Server started on :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	sugarLog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		sugarLog.Fatalf("Server shutdown failed: %v", err)
	}

	sugarLog.Info("Server gracefully stopped")
}
