package main

import "github.com/nelbermora/meli-bootcamp-storage/pkg/db"

func main() {
	db.EmployeesDB.Ping()
}
