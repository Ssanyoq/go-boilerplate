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
		println("Coudn't connect to the DB")
		panic(err)
	}

	table := db.AddTableWithName(models.User{}, "users")
	table.SetKeys(true, "id")
	table.ColMap("email").Unique = true

	createTableErr := db.CreateTablesIfNotExists()
	if createTableErr != nil {
		println("Couldn't create table")
		panic(createTableErr.Error())
	}
	//dropTableErr := db.DropTables() // Uncomment if weird database data is encountered
	//if dropTableErr != nil {
	//	println("Drop table Error")
	//	panic(dropTableErr.Error())
	//}
	if os.Getenv("ENV") == "DEVELOPMENT" {
		db.TraceOn("[gorp]", log.New(os.Stdout, "boilerplate:", log.Lmicroseconds))
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
	return dbmap, err
}

func GetDB() *gorp.DbMap {
	return db
}
