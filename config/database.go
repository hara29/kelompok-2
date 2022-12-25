package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn()(db *sql.DB, err error){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go_kelompok2"

	// melakukan koneksi ke database di golang
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	// db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/go_kelompok2")
	return 
}