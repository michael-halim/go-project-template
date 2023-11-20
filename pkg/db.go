package pkg

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgrework dbname=go_transactions port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// go get github.com/google/uuid
// go mod init github.com/michael-halim
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql
// go get -u github.com/gin-gonic/gin
// go get -u golang.org/x/crypto/bcrypt
// go get -u github.com/golang-jwt/jwt/v5
// go get github.com/joho/godotenv
// go get github.com/githubnemo/CompileDaemon
// go install github.com/githubnemo/CompileDaemon
// compiledaemon --command "./michael-halim"
