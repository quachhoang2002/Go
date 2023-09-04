package httpserver

import (
	todoHTTP "github.com/quachhoang2002/Go/internal/todo/delivery/http"
	todoRepository "github.com/quachhoang2002/Go/internal/todo/repository"
	todoMongo "github.com/quachhoang2002/Go/internal/todo/repository/mongo"
	todoPG "github.com/quachhoang2002/Go/internal/todo/repository/pg"
	todoUsecase "github.com/quachhoang2002/Go/internal/todo/usecase"

	accountHTTP "github.com/quachhoang2002/Go/internal/account/delivery/http"
	accountRepository "github.com/quachhoang2002/Go/internal/account/repository"
	accountRepoPG "github.com/quachhoang2002/Go/internal/account/repository/pg"
	accountUsecase "github.com/quachhoang2002/Go/internal/account/usecase"
)

func (srv HTTPServer) mapHandlers() {
	api := srv.gin.Group("/api")

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
	todo.MapRoutes(api.Group("/todos"))

	var accountRepo accountRepository.Repository
	accountRepo = accountRepoPG.NewRepository(srv.l, srv.dbConfig.pgDB)
	accountUC := accountUsecase.New(srv.l, accountRepo)
	accountHTTP := accountHTTP.NewHandler(srv.l, accountUC)
	accountHTTP.MapRoutes(api.Group("/accounts"))

}
