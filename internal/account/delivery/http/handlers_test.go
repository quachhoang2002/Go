package http

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/quachhoang2002/Go/internal/todo/usecase"
	"github.com/quachhoang2002/Go/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type mockDeps struct {
	uc *usecase.MockUsecase
}

func initHandler(t *testing.T) (handler, mockDeps) {
	uc := usecase.NewMockUsecase(t)
	l := log.InitializeZapLogger(log.NewTestZapConfig())

	return NewHandler(l, uc), mockDeps{
		uc: uc,
	}
}

func TestDeliveryHTTPTodo_add(t *testing.T) {
	type mockUsecase struct {
		expCall bool
		input   usecase.CreateInput
		err     error
	}

	tcs := map[string]struct {
		req         string
		mockUsecase mockUsecase
		wantBody    string
		wantCode    int
	}{
		"success": {
			req: `{
				"name": "test",
				"description": "test"
			}`,
			mockUsecase: mockUsecase{
				expCall: true,
				input: usecase.CreateInput{
					Name:        "test",
					Description: "test",
				},
			},
			wantBody: `{
				"code": 0,
				"message": "Success"
			}`,
			wantCode: http.StatusOK,
		},
		"Invalid body": {
			req: `{
				"name": 1,
				"description": "test"
			}`,
			wantBody: `{
				"code": 501,
				"message": "invalid request body"
			}`,
			wantCode: http.StatusBadRequest,
		},
		"Failed validation": {
			req: `{
				"name": "",
				"description": "test"
			}`,
			wantBody: `{
				"code": 401,
				"message": "Validation failed",
				"errors": [
					{
						"field": "name",
						"messages": [
							"name is required"
						]
					}
				]
			}`,
			wantCode: http.StatusBadRequest,
		},
		"Internal error": {
			req: `{
				"name": "test",
				"description": "test"
			}`,
			mockUsecase: mockUsecase{
				expCall: true,
				input: usecase.CreateInput{
					Name:        "test",
					Description: "test",
				},
				err: errors.New("Internal error"),
			},
			wantBody: `{
				"code": 500,
				"message": "Something went wrong"
			}`,
			wantCode: http.StatusInternalServerError,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request, _ = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer([]byte(tc.req)))
			c.Request.Header.Add("Content-Type", "application/json")

			ctx := c.Request.Context()

			h, deps := initHandler(t)

			if tc.mockUsecase.expCall {
				deps.uc.EXPECT().
					Create(ctx, tc.mockUsecase.input).
					Return(tc.mockUsecase.err).
					Once()
			}

			// WHEN
			h.add(c)

			// THEN
			require.Equal(t, tc.wantCode, w.Code)
			require.JSONEq(t, tc.wantBody, w.Body.String())
		})
	}
}
