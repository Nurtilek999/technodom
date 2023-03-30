package appserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
	"technodom/internal/app/config"
	"technodom/internal/service"
	"technodom/internal/util/logger"
)

func RegisterHooks(
	lifecycle fx.Lifecycle,
	config *config.TomlConfig,
	mux *mux.Router,
	logger logger.Logger,
	service *service.Service,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("hey", nil)

				err := service.FillCache()
				if err != nil {
					errMsg := fmt.Errorf("ошибка при прогерве кэша").Error()
					logger.Error(errMsg, map[string]interface{}{})
				}

				go func() {
					// Запуск сервера
					err := http.ListenAndServe(config.BindAddr, mux)

					if err != nil {
						logger.Info(fmt.Sprintf(err.Error()), map[string]interface{}{"BindAddr": config.BindAddr})
					}
					logger.Info(fmt.Sprintf("Listening on %s", config.BindAddr), map[string]interface{}{})
				}()

				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
