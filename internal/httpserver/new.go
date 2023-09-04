package httpserver

import (
	"database/sql"

	pkgLog "github.com/quachhoang2002/Go/pkg/log"
	"github.com/quachhoang2002/Go/pkg/mongo"

	"github.com/gin-gonic/gin"
)

const (
	DBDriverPG    = "postgres"
	DBDriverMongo = "mongodb"
)

type HTTPServer struct {
	gin      *gin.Engine
	l        pkgLog.Logger
	port     int
	dbConfig DBConfig
}

type DBConfig struct {
	driver  string
	pgDB    *sql.DB
	mongoDB mongo.Database
}

func NewMongoDBConfig(db mongo.Database) DBConfig {
	return DBConfig{
		driver:  DBDriverMongo,
		mongoDB: db,
	}
}

func NewPGDBConfig(db *sql.DB) DBConfig {
	return DBConfig{
		driver: DBDriverPG,
		pgDB:   db,
	}
}

type Config struct {
	Port     int
	DBConfig DBConfig
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	return &HTTPServer{
		l:        l,
		gin:      gin.Default(),
		port:     cfg.Port,
		dbConfig: cfg.DBConfig,
	}
}
