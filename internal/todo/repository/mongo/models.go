package mongo

import (
	"time"

	"gitlab.com/gma-vietnam/tanca-event/internal/model"
)

type todoModel struct {
	ID          string     `bson:"_id"`
	Name        string     `bson:"name"`
	Description string     `bson:"description"`
	CreatedAt   time.Time  `bson:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at"`
	DeleteAt    *time.Time `bson:"deleted_at"`
}

func newTodoModel(todo model.Todo) todoModel {
	return todoModel{
		Name:        todo.Name,
		Description: todo.Description,
	}
}
