package response

import (
	pkgErrors "github.com/quachhoang2002/Go/pkg/errors"

	"github.com/gin-gonic/gin"
)

const (
	// DefaultErrorMessage is the default error message.
	DefaultErrorMessage = "Something went wrong"
	// ValidationErrorCode is the validation error code.
	ValidationErrorCode = 401
	// ValidationErrorMsg is the validation error message.
	ValidationErrorMsg = "Validation failed"
)

// Resp is the response format.
type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

// NewOKResp returns a new OK response with the given data.
func NewOKResp(data any) Resp {
	return Resp{
		Code:    0,
		Message: "Success",
		Data:    data,
	}
}

// NewUnauthorizedResp returns a new Unauthorized response with the given data.
func NewUnauthorizedResp() Resp {
	return Resp{
		Code:    401,
		Message: "Unauthorized",
	}
}

// Ok returns a new OK response with the given data.
func OK(c *gin.Context, data any) {
	c.JSON(200, NewOKResp(data))
}

// Unauthorized returns a new Unauthorized response with the given data.
func Unauthorized(c *gin.Context, data any) {
	c.JSON(401, NewUnauthorizedResp())
}

func parseError(err error) (int, Resp) {
	switch parsedErr := err.(type) {
	case *pkgErrors.ValidationErrorCollector:
		return 400, Resp{
			Code:    ValidationErrorCode,
			Message: ValidationErrorMsg,
			Errors:  parsedErr.Errors(),
		}
	case *pkgErrors.HTTPError:
		return 400, Resp{
			Code:    parsedErr.Code,
			Message: parsedErr.Message,
		}
	default:
		return 500, Resp{
			Code:    500,
			Message: DefaultErrorMessage,
		}
	}
}

// Error returns a new Error response with the given error.
func Error(c *gin.Context, err error) {
	c.JSON(parseError(err))
}

// ErrorMapping is a map of error to HTTPError.
type ErrorMapping map[error]*pkgErrors.HTTPError

// ErrorWithMap returns a new Error response with the given error.
func ErrorWithMap(c *gin.Context, err error, eMap ErrorMapping) {
	if httpErr, ok := eMap[err]; ok {
		Error(c, httpErr)
		return
	}

	Error(c, err)
}
