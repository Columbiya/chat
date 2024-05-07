package main

import (
	"zbchat/app/db"
	websocket_http "zbchat/app/modules/http"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	db.Connect()

	websocket_http.Serve()
}
