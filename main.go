package main

import (
	"fmt"

	"github.com/Thalisonh/go-medium/server"
	"github.com/joho/godotenv"
)

func main() {
	errDotEnv := godotenv.Load()

	if errDotEnv != nil {
		fmt.Println("Error loading .env files")
	}

	s := server.Init()
	s.Run()
}
