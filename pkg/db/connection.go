package db

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite", "devel.sqlite3")
	if err != nil {
		panic(err)
	}
	boil.SetDB(DB)
}
