package app

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/VadimGossip/platform_common/pkg/closer"
)

type App struct {
	serviceProvider *serviceProvider
	name            string
	appStartedAt    time.Time
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context, name string, appStartedAt time.Time) (*App, error) {
	a := &App{
		name:         name,
		appStartedAt: appStartedAt,
	}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}
