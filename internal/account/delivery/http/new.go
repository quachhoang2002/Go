package http

import (
	accountUsecase "github.com/quachhoang2002/Go/internal/account/usecase"
	"github.com/quachhoang2002/Go/pkg/log"
)

type handler struct {
	l         log.Logger
	accountUC accountUsecase.Usecase
}

// NewHandler returns a new instance of the HTTPHandler interface
func NewHandler(l log.Logger, accountUC accountUsecase.Usecase) handler {
	return handler{
		l:         l,
		accountUC: accountUC,
	}
}
