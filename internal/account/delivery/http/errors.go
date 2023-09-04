package http

import (
	"github.com/quachhoang2002/Go/internal/todo/usecase"
	pkgErrors "github.com/quachhoang2002/Go/pkg/errors"
	"github.com/quachhoang2002/Go/pkg/response"
)

var errMap = response.ErrorMapping{
	usecase.ErrNotFound: pkgErrors.NewHTTPError(404, "todo not found"),
}

var (
	errIDIsRequired       = pkgErrors.NewHTTPError(402, "id is required")
	errInvalidID          = pkgErrors.NewHTTPError(403, "id is invalid")
	errInvalidRequestBody = pkgErrors.NewHTTPError(501, "invalid request body")
)

const (
	errMsgNameIsRequired = "owner is required"
)
