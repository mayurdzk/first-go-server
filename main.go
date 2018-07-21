package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	namePtr, passwordPtr := getDBCredentials()
	db := prepareDBConnection(namePtr, passwordPtr)

	router := FirstGoServerRouter(db)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getDBCredentials() (*string, *string) {
	namePtr := flag.String("username", "", "Username for connecting to the database")
	passwordPtr := flag.String("password", "", "Password corresponding to the username")

	flag.Parse()
	if *namePtr == "" || *passwordPtr == "" {
		panic("Valid credentials must be passed to be connected to the database")
	}

	return namePtr, passwordPtr
}

func prepareDBConnection(name, password *string) *DB {
	dataSourceName := fmt.Sprintf("%s:%s@/godb", *name, *password)
	db, err := sql.Open("mysql", dataSourceName)
	assertNotError(err)

	err = db.Ping()
	assertNotError(err)

	return db
}

func assertNotError(e error) {
	if e != nil {
		panic(e)
	}
}

type DB = sql.DB