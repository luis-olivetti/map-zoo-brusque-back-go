package main

import (
	"fmt"

	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/config"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/http"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()

	logger := log.NewLogger()
	defer logger.Sync()

	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
