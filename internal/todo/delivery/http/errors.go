package http

import (
	"gitlab.com/gma-vietnam/tanca-event/internal/todo/usecase"
	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
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
	errMsgNameIsRequired = "name is required"
)
