package server

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/luyanakat/golang-base-project/app/handlers"
	"github.com/luyanakat/golang-base-project/app/repository"
	"github.com/luyanakat/golang-base-project/app/service"
	"github.com/luyanakat/golang-base-project/internal/middleware"
	"github.com/luyanakat/golang-base-project/pkg/db"
	"go.uber.org/zap"
)

func NewHttpServer(logger *zap.SugaredLogger) *gin.Engine {
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

	baseRepo := InitRepository(logger)
	baseService := service.NewBaseService(baseRepo, logger)
	baseHandlers := handlers.NewBaseHandlers(baseService)
	rootRoute := r.Group("/api/v1")

	Routes(rootRoute, baseHandlers)

	logger.Infof("HTTP server is running on port %s", os.Getenv("PORT"))
	return r
}

func Routes(r *gin.RouterGroup, baseHandlers *handlers.BaseHandlers) {
	r.GET("/ping", baseHandlers.PingHandler.Ping)
}

func InitRepository(logger *zap.SugaredLogger) *repository.BaseRepository {
	sqlConn, err := db.NewConnection()
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Info("Database connection established")

	baseRepo := repository.NewBaseRepository(sqlConn)
	return baseRepo
}
