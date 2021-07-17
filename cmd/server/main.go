package main

import "github.com/nelbermora/meli-bootcamp-storage/db"

func main() {
	db.StorageDB.Ping()
}
