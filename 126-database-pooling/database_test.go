package _26_database_pooling

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:Colonelgila123@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
