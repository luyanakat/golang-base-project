package logsvc

import (
	"github.com/luyanakat/golang-base-project/app/repository"
	"go.uber.org/zap"
)

type LogService struct {
	*repository.LogRepository
	*zap.SugaredLogger
}

func NewLogService(repo *repository.LogRepository, logger *zap.SugaredLogger) *LogService {
	return &LogService{
		LogRepository: repo,
		SugaredLogger: logger,
	}
}

func (s *LogService) LogInfo(message string) {
	s.SugaredLogger.Info(message)
}
