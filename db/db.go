package db

import (
	"log"

	"course-teacher/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Init() {
	var err error
	if db, err = sqlx.Open("postgres", config.GetDatabaseURL()); err != nil {
		log.Panicln(err)
	}
}
func GetDB() *sqlx.DB {
	return db
}
