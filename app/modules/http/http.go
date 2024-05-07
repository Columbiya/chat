package websocket_http

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"zbchat/app/modules/server"

	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/ws/{userId}", server.HandleConnection)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("12")
		w.Write([]byte("ping"))
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:5001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
