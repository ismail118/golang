package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"kursus/belajar-golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Colonelgila123@tcp(localhost:3306)/belajar_golang_restful_api?parseTime=true")
	helper.PanicIfError(err)

	// setup database connection pooling
	db.SetMaxOpenConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
