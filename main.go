package main

import (
	"github.com/joho/godotenv"
	"github.com/michael-halim/go-transactions/cmd"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
