package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/config"
	"github.com/AdityaTaggar05/annora-auth/internal/handler/https"
	authhandler "github.com/AdityaTaggar05/annora-auth/internal/handler/https/auth"
	tokenhandler "github.com/AdityaTaggar05/annora-auth/internal/handler/https/token"
	"github.com/AdityaTaggar05/annora-auth/internal/infrastructure/db"
	redisinfra "github.com/AdityaTaggar05/annora-auth/internal/infrastructure/redis"
	tokeninfra "github.com/AdityaTaggar05/annora-auth/internal/infrastructure/token"
	"github.com/AdityaTaggar05/annora-auth/internal/mailer"
	authrepo "github.com/AdityaTaggar05/annora-auth/internal/repository/auth"
	tokenrepo "github.com/AdityaTaggar05/annora-auth/internal/repository/token"
	authservice "github.com/AdityaTaggar05/annora-auth/internal/service/auth"
	tokenservice "github.com/AdityaTaggar05/annora-auth/internal/service/token"
)

type App struct {
	Server *http.Server
	Config *config.Config
}

func New(cfg *config.Config) (*App, error) {
	// 1) Infrastructure Setup
	ctx := context.Background()
	db := db.NewPostgresDB(ctx, cfg.Postgres)
	rdb := redisinfra.NewClient(cfg.Redis)

	// 2) Repository Setup
	authRepo := authrepo.NewRepository(db)
	tokenRepo := tokenrepo.NewRepository(db, rdb)

	// 3) Service Setup
	signingKey, err := tokeninfra.LoadSigningKey(cfg.JWT)
	if err != nil {
		log.Fatal(err)
	}
	mailer := mailer.NewMailer(cfg.Email)

	authService := authservice.NewService(authRepo, tokenRepo, mailer, cfg.JWT, cfg.Email.TokenTTL, signingKey)
	tokenService := tokenservice.NewService(authRepo, tokenRepo, cfg.JWT, signingKey)

	// 4) Handler Setup
	authHandler := authhandler.NewHandler(authService)
	tokenHandler := tokenhandler.NewHandler(tokenService)

	// 5) Router Setup
	router := https.NewRouter(authHandler, tokenHandler)

	// 6) Server Setup
	return &App {
		Config: cfg,
		Server: &http.Server{
			Addr:         ":" + cfg.Server.Port,
			Handler: router,
			ReadTimeout: cfg.Server.ReadTimeout,
			WriteTimeout: cfg.Server.WriteTimeout,
		},
	}, nil
}

func (a *App) Start() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Auth service listening on %s\n", a.Server.Addr)

		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v\n", err)
		}
	}()

	<-stop
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
		return err
	}

	log.Println("Auth service stopped gracefully")
	return nil
}
