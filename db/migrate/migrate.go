package main

import (
	"github.com/michael-halim/go-transactions/db/models"
	"github.com/michael-halim/go-transactions/pkg"
)

func MigrateDatabase() {
	conn := pkg.NewPostgresConnection()
	conn.AutoMigrate(&models.User{})
}

func main() {
	MigrateDatabase()
}
