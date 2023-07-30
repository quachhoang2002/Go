package http

import (
	todoUsecase "gitlab.com/gma-vietnam/tanca-event/internal/todo/usecase"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type handler struct {
	l      log.Logger
	todoUC todoUsecase.Usecase
}

// NewHandler returns a new instance of the HTTPHandler interface
func NewHandler(l log.Logger, todoUC todoUsecase.Usecase) handler {
	return handler{
		l:      l,
		todoUC: todoUC,
	}
}
