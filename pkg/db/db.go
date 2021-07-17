package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	EmployeesDB *sql.DB
)

func init() {
	dataSource := "root@tcp(localhost:3306)/employees"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	EmployeesDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = EmployeesDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")
}
