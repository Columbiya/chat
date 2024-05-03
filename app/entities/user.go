package entities

import "golang.org/x/net/websocket"

type User struct {
	conn *websocket.Conn
}

func New(conn *websocket.Conn) *User {
	return &User{conn}
}

func (u *User) SendMessage(mes *Message, to *websocket.Conn) {
	u.conn.Write([]byte(mes.Text))
}