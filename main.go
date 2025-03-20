package main

import (
	"github.com/luyanakat/golang-base-project/app/server"
	"github.com/luyanakat/golang-base-project/pkg/logger"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()

	sugarLog := logger.Sugar()

	engine := server.NewHttpServer(sugarLog)

	engine.Run(":" + "8080")
}
