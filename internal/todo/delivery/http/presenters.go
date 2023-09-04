package http

import (
	"strings"

	"github.com/quachhoang2002/Go/internal/model"
	todoUseCase "github.com/quachhoang2002/Go/internal/todo/usecase"
	pkgErrors "github.com/quachhoang2002/Go/pkg/errors"
	"github.com/quachhoang2002/Go/pkg/response"
)

type addRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r addRequest) toInput() (todoUseCase.CreateInput, error) {
	vErrCollect := pkgErrors.NewValidationErrorCollector()

	if strings.TrimSpace(r.Name) == "" {
		vErrCollect.Add(pkgErrors.NewValidationError("name", errMsgNameIsRequired))
	}

	if vErrCollect.HasError() {
		return todoUseCase.CreateInput{}, vErrCollect
	}

	return todoUseCase.CreateInput{
		Name:        r.Name,
		Description: r.Description,
	}, nil
}

type todoResponse struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	CreatedAt   response.DateTime `json:"created_at"`
}

func newListResponse(todos []model.Todo) []todoResponse {
	data := []todoResponse{}
	for _, todo := range todos {
		data = append(data, todoResponse{
			ID:          todo.ID,
			Name:        todo.Name,
			Description: todo.Description,
			CreatedAt:   response.DateTime(todo.CreatedAt),
		})
	}

	return data
}
