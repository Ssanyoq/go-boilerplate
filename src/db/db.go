package db

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	"github.com/ssanyoq/go-boilerplate/src/models"
	"log"
	"os"
)

var db *gorp.DbMap

func Init() {

	authData := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_LOGIN"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(authData)
	if err != nil {
		panic(err)
	}
	db.AddTableWithName(models.User{}, "users").SetKeys(true, "id")
	db.CreateTables()
	db.TraceOn("[gorp]", log.New(os.Stdout, "boilerplate:", log.Lmicroseconds))
}

func ConnectDB(authData string) (*gorp.DbMap, error) {
	db, err := sql.Open(os.Getenv("DB_DIALECT"), authData)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return dbmap, err
}

func GetDB() *gorp.DbMap {
	return db
}
