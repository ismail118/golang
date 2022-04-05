package _35_repository_pattern

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "eko@gmail.com"
	comment := "hello world"

	ctx := context.Background()

	sqlQuery := "INSERT INTO comments(email, comment) VALUES (?, ?);"

	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data to Database", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email, comment) VALUES (?, ?);"
	stmt, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("eko%s@gmail.com", strconv.Itoa(i))
		comment := fmt.Sprintf("ini koemen ke %s", strconv.Itoa(i))
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Comment Id:", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	sqlQuery := "INSERT INTO comments(email, comment) VALUES (?, ?);"

	for i := 0; i < 10; i++ {
		email := fmt.Sprintf("eko%s@gmail.com", strconv.Itoa(i))
		comment := fmt.Sprintf("ini koemen ke %s", strconv.Itoa(i))
		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Comment Id:", lastInsertId)
	}

	//err = tx.Commit()
	err = tx.Rollback() // data yang di insert akan di batalakan jika memanggil ini
	if err != nil {
		panic(err)
	}
}
