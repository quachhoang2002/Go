package mongo

import (
	"context"
	"time"

	"github.com/quachhoang2002/Go/internal/model"
	"github.com/quachhoang2002/Go/pkg/mongo"
)

const (
	collectionName = "todos"
)

func (repo implRepository) getTotoCollection() mongo.Collection {
	return repo.db.Collection(collectionName)
}

// Create creates a todo
func (repo implRepository) Create(ctx context.Context, todo model.Todo) error {
	dbtodo := newTodoModel(todo)

	dbtodo.CreatedAt = time.Now()
	dbtodo.UpdatedAt = time.Now()

	_, err := repo.getTotoCollection().InsertOne(ctx, dbtodo)

	return err
}

// All returns all todos
func (repo implRepository) All(ctx context.Context) ([]model.Todo, error) {
	panic("not implemented") // TODO: Implement
}

// Delete deletes a todo
func (repo implRepository) Delete(ctx context.Context, id int) error {
	panic("not implemented") // TODO: Implement:w
}
