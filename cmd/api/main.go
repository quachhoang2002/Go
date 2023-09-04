package main

import (
	"fmt"

	"github.com/quachhoang2002/Go/config"
	"github.com/quachhoang2002/Go/internal/appconfig/mongo"
	"github.com/quachhoang2002/Go/internal/appconfig/pg"
	"github.com/quachhoang2002/Go/internal/httpserver"
	pkgLog "github.com/quachhoang2002/Go/pkg/log"
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
