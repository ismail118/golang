package _26_database_pooling

import (
	"context"
	"database/sql"
	"fmt"
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

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "INSERT INTO customer(id, name) VALUES ('eko', 'EKO');"

	_, err := db.ExecContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data to Database")
}
