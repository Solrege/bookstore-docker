package main

import (
	"bookstore/internal/business"
	"bookstore/internal/platform"
)

func main() {
	db := platform.DbConnection()
	err := db.Migrator().CreateTable(&business.Product{})
	if err != nil {
		return
	}
}
