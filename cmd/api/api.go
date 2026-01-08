package api

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/xOphiuchus/OnlyDrive/routes"
	"gorm.io/gorm"
)

type APIServer struct {
	Host string
	DB   *gorm.DB
	App  *fiber.App
}

func NewAPIServer(host string, db *gorm.DB) *APIServer {
	return &APIServer{
		Host: host,
		DB:   db,
		App: fiber.New(fiber.Config{
			StrictRouting:         true,
			CaseSensitive:         true,
			Prefork:               false,
			DisableStartupMessage: true,
			AppName:               "OnlyDrive API",
			TrustedProxies:        []string{"0.0.0.0/0"},
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				log.Error().Err(err).Msgf("API error: %v", err.Error())
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Something went wrong",
				})
			},
			BodyLimit:    30 * 1024 * 1024, // 30 MB
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		}),
	}
}

func (s *APIServer) Start() error {

	v1_router := s.App.Group("/api/v1")

	routes.SetupRoutes(&v1_router, s.DB)

	go s.listenForShutdown()

	log.Info().Msgf("Starting API server on %s", s.Host)
	return s.App.Listen(s.Host)
}

func (s *APIServer) listenForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Info().Msg("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Info().Msg("Shutdown signal received")
	log.Info().Msg("Gracefully stopping server...")
	log.Info().Msg("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	countdownDone := make(chan struct{})
	go func() {
		for i := 10; i > 0; i-- {
			log.Info().Msgf("Shutting down in %d seconds...", i)
			time.Sleep(1 * time.Second)
		}
		countdownDone <- struct{}{}
	}()

	<-countdownDone

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.App.ShutdownWithContext(ctx); err != nil {
		log.Error().Err(err).Msg("Error during graceful shutdown")
		os.Exit(1)
	}

	log.Info().Msg("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Info().Msg("Server shutdown complete")
	log.Info().Msg("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	os.Exit(0)
}
