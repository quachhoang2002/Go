package httpserver

import (
	todoHTTP "gitlab.com/gma-vietnam/tanca-event/internal/todo/delivery/http"
	todoRepository "gitlab.com/gma-vietnam/tanca-event/internal/todo/repository"
	todoMongo "gitlab.com/gma-vietnam/tanca-event/internal/todo/repository/mongo"
	todoPG "gitlab.com/gma-vietnam/tanca-event/internal/todo/repository/pg"
	todoUsecase "gitlab.com/gma-vietnam/tanca-event/internal/todo/usecase"
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
