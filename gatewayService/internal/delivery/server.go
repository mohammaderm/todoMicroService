package delivery

import (
	"context"
	"net/http"
	"time"

	"github.com/mohammaderm/todoMicroService/gatewayService/config"
	"github.com/mohammaderm/todoMicroService/gatewayService/pkg/logger"
)

func ServerProvider(logger logger.Logger, config config.Server, handler http.Handler) (*http.Server, func()) {
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: handler,
	}
	return srv, func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.GracefulShutdownPeriod)*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Warning("error while shutting server down", map[string]interface{}{
				"err": err.Error(),
			})
		}
	}
}
