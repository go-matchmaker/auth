//go:build wireinject
// +build wireinject

package app

import (
	"auth/internal/adapter/auth/paseto"
	"auth/internal/adapter/config"
	psql "auth/internal/adapter/storage/postgres"
	adapter_http "auth/internal/adapter/transport/http"
	"auth/internal/core/port/auth"
	"auth/internal/core/port/db"
	"auth/internal/core/port/http"
	"auth/internal/core/port/user"
	port_service "auth/internal/core/service"
	"context"
	"github.com/google/wire"
	"go.uber.org/zap"
	"sync"
)

func InitApp(
	ctx context.Context,
	wg *sync.WaitGroup,
	rw *sync.RWMutex,
	Cfg *config.Container,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		httpServerFunc,
		psql.UserRepositorySet,
		port_service.UserServiceSet,
		paseto.PasetoSet,
	))
}

func dbEngineFunc(
	ctx context.Context,
	Cfg *config.Container) (db.PostgresEngineMaker, func(), error) {
	psqlDb := psql.NewDB(Cfg)
	err := psqlDb.Start(ctx)
	if err != nil {
		zap.S().Fatal("failed to start db:", err)
	}

	return psqlDb, func() { psqlDb.Close(ctx) }, nil
}

func httpServerFunc(
	ctx context.Context,
	Cfg *config.Container,
	UserService user.UserServicePort,
	tokenMaker auth.TokenMaker,
) (http.ServerMaker, func(), error) {
	httpServer := adapter_http.NewHTTPServer(ctx, Cfg, UserService, tokenMaker)
	err := httpServer.Start(ctx)
	if err != nil {
		return nil, nil, err
	}
	return httpServer, func() { httpServer.Close(ctx) }, nil
}
