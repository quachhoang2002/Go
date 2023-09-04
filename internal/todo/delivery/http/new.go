package http

import (
	todoUsecase "github.com/quachhoang2002/Go/internal/todo/usecase"
	"github.com/quachhoang2002/Go/pkg/log"
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
