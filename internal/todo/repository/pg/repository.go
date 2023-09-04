package pg

import (
	"context"
	"database/sql"

	"github.com/quachhoang2002/Go/internal/dbmodel"
	"github.com/quachhoang2002/Go/internal/model"
	"github.com/quachhoang2002/Go/internal/todo/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Create creates a new todo
func (repo implRepository) Create(ctx context.Context, todo model.Todo) error {
	dbm := dbmodel.Todo{
		Name:        todo.Name,
		Description: todo.Description,
	}

	wl := boil.Whitelist(
		dbmodel.TodoColumns.Name,
		dbmodel.TodoColumns.Description,
	)
	if err := dbm.Insert(ctx, repo.db, wl); err != nil {
		repo.l.Errorf(ctx, "todo.repository.Create.dbm.Insert %+v: %s", dbm, err)
		return err
	}

	return nil
}

// All returns all todos
func (repo implRepository) All(ctx context.Context) ([]model.Todo, error) {
	dbtodos, err := dbmodel.Todos().All(ctx, repo.db)
	if err != nil {
		repo.l.Errorf(ctx, "todo.repository.All.dbm.All: %s", err)
		return nil, err
	}

	var todos []model.Todo
	for _, dbtodo := range dbtodos {
		todos = append(todos, model.Todo{
			ID:          dbtodo.ID,
			Name:        dbtodo.Name,
			Description: dbtodo.Description,
			CreatedAt:   dbtodo.CreatedAt,
			UpdatedAt:   dbtodo.UpdatedAt,
		})
	}

	return todos, nil
}

// Delete deletes a todo
func (repo implRepository) Delete(ctx context.Context, id int) error {
	dbm, err := dbmodel.Todos(dbmodel.TodoWhere.ID.EQ(id)).One(ctx, repo.db)
	if err != nil {
		if err == sql.ErrNoRows {
			repo.l.Warnf(ctx, "todo.repository.Delete.dbm.One id=%d: %s", id, err)
			return repository.ErrNotFound
		}

		repo.l.Errorf(ctx, "todo.repository.Delete.dbm.One id=%d: %s", id, err)
		return err
	}

	_, err = dbm.Delete(ctx, repo.db, false)
	if err != nil {
		repo.l.Warnf(ctx, "todo.repository.Delete.dbm.Delete id=%d: %s", id, err)
		return err
	}

	return nil
}
