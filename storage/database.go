package storage

import (
	"os"
	"strings"

	"github.com/go-pg/pg/v10"
)

func GetDatabase() *pg.DB {
	var POSTGRES_HOST string = os.Getenv("POSTGRES_HOST")
	var POSTGRES_USER string = os.Getenv("POSTGRES_USER")
	var POSTGRES_PASS string = os.Getenv("POSTGRESS_PASSWORD")
	var POSTGRES_DB string = os.Getenv("POSTGRES_DATABASE")
	var POSTGRES_PORT string = os.Getenv("POSTGRES_PORT")

	var db *pg.DB = pg.Connect(&pg.Options{
		User:     POSTGRES_USER,
		Password: POSTGRES_PASS,
		Database: POSTGRES_DB,
		Addr:     strings.Join([]string{POSTGRES_HOST, POSTGRES_PORT}, ":"),
	})

	return db
}
