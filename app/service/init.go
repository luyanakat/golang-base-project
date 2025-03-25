package service

import (
	"github.com/luyanakat/golang-base-project/app/repository"
	"github.com/luyanakat/golang-base-project/app/service/logsvc"
	"go.uber.org/zap"
)

type BaseService struct {
	*logsvc.LogService
}

func NewBaseService(baseRepo *repository.BaseRepository, logger *zap.SugaredLogger) *BaseService {
	return &BaseService{
		LogService: logsvc.NewLogService(baseRepo.LogRepo, logger),
	}
}
