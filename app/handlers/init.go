package handlers

import "github.com/luyanakat/golang-base-project/app/service"

type BaseHandlers struct {
	BaseSvc     *service.BaseService
	PingHandler *PingHandler
}

func NewBaseHandlers(baseSvc *service.BaseService) *BaseHandlers {
	pingHandler := NewPingHandler()
	return &BaseHandlers{
		PingHandler: pingHandler,
		BaseSvc:     baseSvc,
	}
}
