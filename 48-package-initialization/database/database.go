package database

import "fmt"

var connection string

func init() {
	fmt.Println("init database di jalankan")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
