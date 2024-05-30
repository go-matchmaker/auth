package app

import (
	"auth/internal/adapter/config"
	"auth/internal/core/port/auth"
	"auth/internal/core/port/db"
	"auth/internal/core/port/http"
	"auth/internal/core/port/user"
	"context"
	"go.uber.org/zap"
	"sync"
)

type App struct {
	rw          *sync.RWMutex
	Cfg         *config.Container
	HTTP        http.ServerMaker
	Token       auth.TokenMaker
	PG          db.PostgresEngineMaker
	UserRepo    user.UserRepositoryPort
	UserService user.UserServicePort
}

func New(
	rw *sync.RWMutex,
	cfg *config.Container,
	http http.ServerMaker,
	token auth.TokenMaker,
	pg db.PostgresEngineMaker,
	userRepo user.UserRepositoryPort,
	userService user.UserServicePort,
) *App {
	return &App{
		rw:          rw,
		Cfg:         cfg,
		HTTP:        http,
		Token:       token,
		PG:          pg,
		UserRepo:    userRepo,
		UserService: userService,
	}
}

func (a *App) Run(ctx context.Context) {
	zap.S().Info("RUNNER!")
}
