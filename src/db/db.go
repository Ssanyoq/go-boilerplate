package db

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
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
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, err
}

func GetDB() *gorp.DbMap {
	return db
}
