package entities

import (
	"context"
	"zbchat/app/db"

	"github.com/gorilla/websocket"
)

type User struct {
	UserGuid string
	conn     *websocket.Conn
}

var Userconmap = make(map[string]*User)

func New(conn *websocket.Conn, userGuid string) *User {
	return &User{conn: conn, UserGuid: userGuid}
}

func (u *User) ReceiveMessageFrom(author *User, text *string) {
	if err := u.conn.WriteJSON(NewMessage(text, author.UserGuid)); err != nil {
		u.conn.Close()
		delete(Userconmap, u.UserGuid)
		panic(err)
	}
}

func (u *User) SendMessage(text *string, toUserGuid *string) {
	user, ok := Userconmap[*toUserGuid]

	if ok {
		user.ReceiveMessageFrom(u, text)
	}

	_, err := db.MessagesColl.InsertOne(context.Background(), NewDatabaseMessage(text, &u.UserGuid, toUserGuid))

	if err != nil {
		panic(err)
	}
}

func (u *User) SetToConnectionMap() {
	Userconmap[u.UserGuid] = u
}
