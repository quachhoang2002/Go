package httpserver

import (
	todoHTTP "github.com/quachhoang2002/Go/internal/todo/delivery/http"
	todoRepository "github.com/quachhoang2002/Go/internal/todo/repository"
	todoMongo "github.com/quachhoang2002/Go/internal/todo/repository/mongo"
	todoPG "github.com/quachhoang2002/Go/internal/todo/repository/pg"
	todoUsecase "github.com/quachhoang2002/Go/internal/todo/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// Repositories
	var todoRepo todoRepository.Repository
	if srv.dbConfig.driver == DBDriverMongo {
		todoRepo = todoMongo.NewRepository(srv.l, srv.dbConfig.mongoDB)
	} else {
		todoRepo = todoPG.NewRepository(srv.l, srv.dbConfig.pgDB)
	}

	// Usecases
	todoUC := todoUsecase.New(srv.l, todoRepo)

	// Handlers
	todo := todoHTTP.NewHandler(srv.l, todoUC)

	api := srv.gin.Group("/api")

	todo.MapRoutes(api.Group("/todos"))
}
