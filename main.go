package main

import (
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"technodom/internal/app/appserver"
	"technodom/internal/app/config"
	"technodom/internal/app/handler"
	"technodom/internal/repository/cache"
	repo "technodom/internal/repository/pgrepository"
	"technodom/internal/service"
	logrus_log "technodom/internal/util/logger/logrus-log"
)

func main() {
	fx.New(
		fx.Provide(
			service.NewService,
			config.NewConfig,
			logrus_log.New,
			mux.NewRouter,
			repo.New,
			cache.NewLocalCache,
		),

		fx.Invoke(
			handler.New,
			appserver.RegisterHooks,
		),
	).Run()

}
