package app

import (
	"github.com/gin-gonic/gin"
	"github.com/luyanakat/golang-base-project/app/server"
	"go.uber.org/zap"
)

type App struct {
	R *gin.Engine
}

func InitApp(logger *zap.SugaredLogger) *App {
	return &App{
		R: server.NewHttpServer(logger),
	}
}
