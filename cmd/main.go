package main

import (
	"context"
	"errors"
	"go-sneed/internal/config"
	"go-sneed/internal/handlers"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

/*

 */

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    cfg := config.MustLoadConfig()

	r := chi.NewRouter()
    fileServer := http.FileServer(http.Dir("./static"))
    r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	r.Get("/", handlers.NewGetHomeHandler().ServeHTTP)
    r.NotFound(handlers.NewGetNotFoundHandler().ServeHTTP)

	killSig := make(chan os.Signal, 1)
    signal.Notify(killSig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	srv := &http.Server{
        Addr:    ":"+cfg.Port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else if err != nil {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", cfg.Port))
	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")
}
