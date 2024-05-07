package server

import (
	"fmt"
	"log"
	"net/http"
	"zbchat/app/entities"
	"zbchat/app/modules/messageprocessor"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["userId"]) // connected user id
	conn, err := upgrader.Upgrade(w, r, nil)

	user := entities.New(conn, vars["userId"])

	if err != nil {
		log.Println(err)
		return
	}

	for {
		user.SetToConnectionMap()

		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		messageprocessor.ProcessPayload(&p, user)
	}
}
