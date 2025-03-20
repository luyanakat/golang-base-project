package server

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/luyanakat/golang-base-project/app/routes"
	"github.com/luyanakat/golang-base-project/internal/middleware"
	"go.uber.org/zap"
)

type HttpServer struct {
	*gin.Engine
}

func NewHttpServer(logger *zap.SugaredLogger) *HttpServer {
	r := gin.New()
	gin.SetMode(os.Getenv("GIN_MODE"))

	middleware := middleware.NewMiddleware(logger)
	r.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "x-api-key"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Authorization"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}),
		middleware.RequestTraceMiddleware,
	)

	rootRoute := r.Group("/api/v1")
	routes.PingRoute(rootRoute)

	logger.Infof("HTTP server is running on port %s", os.Getenv("PORT"))
	return &HttpServer{r}
}
