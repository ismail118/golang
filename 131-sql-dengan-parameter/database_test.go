package _26_database_pooling

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
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

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlQuery := "SELECT id, name FROM customer;"

	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %s \nName: %s\n", id, name)
	}

	fmt.Println("Success Select Data to Database")
}

func TestTipeDataColumn(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM buyer;"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=========================")
		fmt.Println("id: ", id)
		fmt.Println("name: ", name)
		if email.Valid {
			fmt.Println("email: ", email.String)
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		if birthDate.Valid {
			fmt.Println("birth date: ", birthDate.Time)
		}
		fmt.Println("married: ", married)
		fmt.Println("created at: ", createdAt)
	}

	fmt.Println("Success Select Data to Database")
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin'; #" // contoh sql injection
	password := "salah"

	ctx := context.Background()
	sqlQuery := "SELECT username FROM user WHERE username='" + username + "' AND password ='" + password + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Berhasil Login: %s \n", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin'; #" // contoh sql injection
	password := "salah"

	ctx := context.Background()
	sqlQuery := "SELECT username FROM user WHERE username= ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Berhasil Login: %s \n", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "eko'; DROP TABLE user; #" // test sql injection
	password := "eko"

	ctx := context.Background()

	sqlQuery := "INSERT INTO user(username, password) VALUES (?, ?);"

	_, err := db.ExecContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data to Database")
}
