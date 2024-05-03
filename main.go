package main

import (
	"fmt"
	"os"
	"zbchat/app/db"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("123")

	fmt.Println(os.Getenv("HOST"))

	db.Connect()
}