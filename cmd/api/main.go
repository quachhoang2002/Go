package main

import (
	"fmt"

	"gitlab.com/gma-vietnam/tanca-event/config"
	"gitlab.com/gma-vietnam/tanca-event/internal/appconfig/mongo"
	"gitlab.com/gma-vietnam/tanca-event/internal/appconfig/pg"
	"gitlab.com/gma-vietnam/tanca-event/internal/httpserver"
	pkgLog "gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	var dbConfig httpserver.DBConfig
	switch cfg.Database.Driver {
	case httpserver.DBDriverMongo:
		client, err := mongo.Connect(cfg.Mongo.URL)
		if err != nil {
			panic(err)
		}
		defer mongo.Disconnect(client)

		db := client.Database(cfg.Mongo.DBName)
		dbConfig = httpserver.NewMongoDBConfig(db)
	case httpserver.DBDriverPG:
		db, err := pg.Connect(cfg.PG.URL)
		if err != nil {
			panic(err)
		}

		dbConfig = httpserver.NewPGDBConfig(db)
	default:
		panic(fmt.Sprintf("Database driver %s is not supported", cfg.Database.Driver))
	}

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	srv := httpserver.New(l, httpserver.Config{
		Port:     cfg.HTTPServer.Port,
		DBConfig: dbConfig,
	})
	srv.Run()
}
