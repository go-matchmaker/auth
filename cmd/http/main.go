package main

import (
	"auth/internal/adapter/app"
	"auth/internal/adapter/config"
	"auth/internal/adapter/logger"
	"auth/internal/core/util"
	"context"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
	"os/signal"
	"sync"
)

func main() {
	_, err := maxprocs.Set()
	if err != nil {
		panic("failed set max procs")
	}
	ctx, cancel := signal.NotifyContext(context.Background(), util.InterruptSignals...)
	defer cancel()
	wg := new(sync.WaitGroup)
	rw := new(sync.RWMutex)
	cfg, err := config.NewConfig()
	if err != nil {
		panic("failed get config: " + err.Error())
	}

	Logger := logger.InitLogger(cfg.Log.Level)
	defer Logger.Sync()

	cleanup := prepareApp(ctx, wg, rw, cfg)
	zap.S().Info("⚡ Service name:", cfg.Name)
	<-ctx.Done()
	zap.S().Info("Context signal received, shutting down")
	wg.Wait()
	zap.S().Info("Waiting for all goroutines to finish")
	cleanup()
	zap.S().Info("Shutting down successfully")

}

func prepareApp(ctx context.Context, wg *sync.WaitGroup, rw *sync.RWMutex, cfg *config.Container) func() {
	var errMsg error
	a, cleanUp, errMsg := app.InitApp(ctx, wg, rw, cfg)
	if errMsg != nil {
		zap.S().Error("failed init app", errMsg)
		<-ctx.Done()
	}
	a.Run(ctx)
	return cleanUp
}
